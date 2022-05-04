package jsong

import (
    "bytes"
    "context"
    "encoding/json"
    "github.com/containous/traefik/v2/pkg/log"
    "github.com/golang/protobuf/jsonpb"
    "github.com/golang/protobuf/proto"
    "github.com/jhump/protoreflect/desc"
    "github.com/jhump/protoreflect/dynamic"
    "github.com/pkg/errors"
    "google.golang.org/genproto/googleapis/rpc/errdetails"
    _ "google.golang.org/genproto/googleapis/rpc/errdetails"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
    "io"
    "io/ioutil"
    "net/http"
    "strconv"
    "strings"
    "sync"
)

var (
    unmarshaler              = &jsonpb.Unmarshaler{AllowUnknownFields: true}
    marshaler                = &jsonpb.Marshaler{EmitDefaults: true}
    urlNotFoundContent       = []byte(`{"code":404,"msg":"请求地址不存在","data":null}`)
    methodNotAllowedContent  = []byte(`{"code":405,"msg":"请求方法不允许","data":null}`)
    serverUnavailableContent = []byte(`{"code":503,"msg":"上游服务不可用","data":null}`)
    emptyHeader              = make(http.Header, 0)
)

type ErrorResp struct {
    Code    int                    `json:"code"`
    Msg     string                 `json:"msg"`
    Data    map[string]interface{} `json:"data"`
    Details map[string]interface{} `json:"details"`
}

type RoundTripper struct {
    clientConnHandler sync.Map
    logCtx            context.Context
}

func NewRoundTripper() *RoundTripper {
    return &RoundTripper{
        logCtx: log.With(context.Background(), log.Str("component", "jsong.RoundTripper")),
    }
}

func (r *RoundTripper) RoundTrip(req *http.Request) (resp *http.Response, err error) {
    method, ok := Reflections.Load(req.URL.Path)
    if !ok {
        return newResponse(http.StatusNotFound, req, emptyHeader, urlNotFoundContent), nil
    }

    var reqBody []byte
    switch req.Method {
    case "GET":
        query := make(map[string]interface{})
        for k, vv := range req.URL.Query() {
            clean := make([]string, 0, 1)
            for _, v := range vv {
                //过滤空字符串，如果后端需要int，前端传递了空串，转化的时候会err
                if v != "" {
                    clean = append(clean, v)
                }
            }

            l := len(clean)
            if l == 0 {
                continue
            }

            // 支持二级结构，pager.size=20
            if strings.ContainsRune(k, '.') {
                strs := strings.SplitN(k, ".", 2)
                var m map[string]interface{}
                var ok bool
                if query[strs[0]] == nil {
                    m = make(map[string]interface{})
                } else {
                    m, ok = query[strs[0]].(map[string]interface{})
                    if !ok {
                        continue
                    }
                }

                m[strs[1]] = clean
                query[strs[0]] = m
            } else {
                query[k] = clean
            }
        }
        reqBody, err = json.Marshal(query)
        if err != nil {
            log.FromContext(r.logCtx).Errorf("marshal request query err:%v", errors.WithStack(err))
            return nil, err
        }
    case "POST":
        if req.Body == nil {
            break
        }

        reqBody, err = ioutil.ReadAll(req.Body)
        defer func() {
            _ = req.Body.Close()
        }()

        if err != nil {
            log.FromContext(r.logCtx).Errorf("read all request body err:%v", errors.WithStack(err))
            return nil, err
        }
    default:
        return newResponse(http.StatusMethodNotAllowed, req, emptyHeader, methodNotAllowedContent), nil
    }

    eventHandler := GetDefaultEventHandler()
    defer PutDefaultEventHandler(eventHandler)
    cc, err := GetGrpcCC(req.Context(), req.URL.Host)
    if err != nil {
        return nil, err
    }

    methodDesc := method.(*desc.MethodDescriptor)

    if userAgent := req.Header.Get("user-agent"); userAgent != "" {
        req.Header.Add("x-user-agent", userAgent)
    }
    if err := InvokeRPC(req.Context(), methodDesc, cc, req.Header, eventHandler, decodeMessage(reqBody)); err != nil {
        if IsRequestError(err) {
            return newResponse(http.StatusBadRequest, req, nil, []byte(err.Error())), nil
        }
        return nil, err
    }

    //merge http header
    respHeader := make(http.Header, len(eventHandler.GetHeaders())+len(eventHandler.GetTrailers()))
    for k, v := range eventHandler.GetHeaders() {
        for _, vv := range v {
            respHeader.Add(k, vv)
        }
    }
    for k, v := range eventHandler.GetTrailers() {
        for _, vv := range v {
            respHeader.Add(k, vv)
        }
    }
    respHeader.Set("Content-Type", "application/json; charset=utf8")

    if err := eventHandler.Err(); err != nil {
        return r.handleErr(respHeader, req, err)
    }

    respHeader.Set("Grpc-Code", "0")
    respHeader.Set("Grpc-Message", "")
    //respHeader.Set("Trace-Id", xtrace.TraceIdFromContext(req.Context()))

    response, ok := eventHandler.GetResponse().(*dynamic.Message)
    if !ok {
        err = errors.New("响应不是 *dynamic.Message 类型")
        log.FromContext(r.logCtx).Error("grpc 请求失败 err:%v", err)
        return nil, err
    }

    codeDesc := response.FindFieldDescriptorByName("code")
    msgDesc := response.FindFieldDescriptorByName("msg")
    dataDesc := response.FindFieldDescriptorByName("data")
    body, err := response.MarshalJSONPB(marshaler)
    if err != nil {
        log.FromContext(r.logCtx).Error("序列化 grpc 响应失败 err:%v", err)
        return nil, err
    }
    return newResponse(http.StatusOK, req, respHeader, body), nil

    // 如果外层有code,msg,data字段直接返回（兼容旧的格式）
    // 没有话会包装成此结构
    if codeDesc != nil && msgDesc != nil && dataDesc != nil {
        return newResponse(http.StatusOK, req, respHeader, body), nil
    } else {
        var buf bytes.Buffer
        buf.WriteString(`{"code":0,"msg":"succ","data":`)
        buf.Write(body)
        buf.WriteString(`}`)
        return newResponse(http.StatusOK, req, respHeader, buf.Bytes()), nil
    }
}

func (r *RoundTripper) handleErr(respHeader http.Header, req *http.Request, err error) (*http.Response, error) {
    st, ok := status.FromError(err)
    if !ok {
        return nil, err
    }
    httpCode := HTTPStatusFromCode(st.Code())
    respHeader.Set("Grpc-Code", strconv.FormatInt(int64(st.Code()), 10))
    respHeader.Set("Grpc-Message", st.Err().Error())

    //traceId := xtrace.TraceIdFromContext(req.Context())
    traceId := "88888888"

    errorResp := ErrorResp{
        Code: httpCode,
        Msg:  st.Message(),
        Data: map[string]interface{}{},
        Details: map[string]interface{}{
            "traceId": traceId,
        },
    }

    // 避免暴漏内部细节
    if httpCode >= 500 {
        log.FromContext(r.logCtx).WithError(err).WithField("traceId", traceId).Error("请求上游服务异常")
        errorResp.Msg = "系统繁忙，请稍后重试。"
    } else {
        log.FromContext(r.logCtx).WithError(err).WithField("traceId", traceId).Info("请求上游服务异常")
    }
    //errorInfo, _ := xerr.FirstErrorInfo(err)
    //errorInfo := err.Error()
    //if errorInfo != nil {
    //    errorResp.Data["errorInfo"] = errorInfo
    //}
    //help, _ := xerr.FirstHelp(err)
    //if help != nil {
    //    errorResp.Data["help"] = help
    //}
    //preconditionFailure, _ := xerr.FirstPreconditionFailure(err)
    //if help != nil {
    //    errorResp.Data["preconditionFailure"] = preconditionFailure
    //}
    //badRequest, _ := xerr.FirstBadRequest(err)
    //if help != nil {
    //    errorResp.Data["badRequest"] = badRequest
    //}

    for _, detail := range st.Details() {
        switch t := detail.(type) {
        case *errdetails.RetryInfo:
            errorResp.Details["retryInfo"] = t
        case *errdetails.RequestInfo:
            errorResp.Details["requestInfo"] = t
        case *errdetails.ResourceInfo:
            errorResp.Details["resourceInfo"] = t
        case *errdetails.LocalizedMessage:
            errorResp.Details["localizedMessage"] = t
        case *errdetails.QuotaFailure:
            errorResp.Details["quotaFailure"] = t
        }
    }

    if len(errorResp.Data) == 0 {
        errorResp.Data = nil
    }

    body, err := json.Marshal(errorResp)
    if err != nil {
        //return nil, xerr.WithStack(err)
        return nil, errors.New(err.Error())
    }

    return newResponse(httpCode, req, respHeader, body), nil
}

func decodeMessage(reqBody []byte) RequestSupplier {
    return func(msg proto.Message) error {
        err := msg.(*dynamic.Message).UnmarshalJSONPB(unmarshaler, reqBody)
        if err == nil {
            return io.EOF
        }
        return err
    }
}

func newResponse(status int, req *http.Request, header http.Header, body []byte) *http.Response {
    return &http.Response{
        Status:        http.StatusText(status),
        StatusCode:    status,
        Proto:         req.Proto,
        ProtoMajor:    req.ProtoMajor,
        ProtoMinor:    req.ProtoMinor,
        Body:          ioutil.NopCloser(bytes.NewBuffer(body)),
        ContentLength: int64(len(body)),
        Request:       req,
        Header:        header,
    }
}

func HTTPStatusFromCode(code codes.Code) int {
    switch code {
    case codes.OK:
        return http.StatusOK
    case codes.Canceled:
        return http.StatusRequestTimeout
    case codes.Unknown:
        return http.StatusInternalServerError
    case codes.InvalidArgument:
        return http.StatusBadRequest
    case codes.DeadlineExceeded:
        return http.StatusGatewayTimeout
    case codes.NotFound:
        return http.StatusNotFound
    case codes.AlreadyExists:
        return http.StatusConflict
    case codes.PermissionDenied:
        return http.StatusForbidden
    case codes.Unauthenticated:
        return http.StatusUnauthorized
    case codes.ResourceExhausted:
        return http.StatusTooManyRequests
    case codes.FailedPrecondition:
        // Note, this deliberately doesn't translate to the similarly named '412 Precondition Failed' HTTP response status.
        return http.StatusBadRequest
    case codes.Aborted:
        return http.StatusConflict
    case codes.OutOfRange:
        return http.StatusBadRequest
    case codes.Unimplemented:
        return http.StatusNotImplemented
    case codes.Internal:
        return http.StatusInternalServerError
    case codes.Unavailable:
        return http.StatusServiceUnavailable
    case codes.DataLoss:
        return http.StatusInternalServerError
    }
    return http.StatusInternalServerError
}
