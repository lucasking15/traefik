// Copyright (c) 2019 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package jsong

import (
	"github.com/golang/protobuf/proto"
	"github.com/jhump/protoreflect/desc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var _ InvocationEventHandler = &DefaultEventHandler{}

type DefaultEventHandler struct {
	headers  metadata.MD
	trailers metadata.MD
	response proto.Message
	err      error
}

func NewDefaultEventHandler() *DefaultEventHandler {
	return &DefaultEventHandler{}
}

func (i *DefaultEventHandler) OnResolveMethod(*desc.MethodDescriptor) {}

func (i *DefaultEventHandler) OnSendHeaders(metadata.MD) {}

func (i *DefaultEventHandler) OnReceiveHeaders(headers metadata.MD) {
	i.headers = headers
}
func (i *DefaultEventHandler) OnReceiveResponse(message proto.Message) {
	i.response = message
}

func (i *DefaultEventHandler) OnReceiveTrailers(s *status.Status, trailers metadata.MD) {
	if err := s.Err(); err != nil {
		i.err = err
	}
	i.trailers = trailers
}

func (i *DefaultEventHandler) GetHeaders() metadata.MD {
	return i.headers
}

func (i *DefaultEventHandler) GetTrailers() metadata.MD {
	return i.trailers
}

func (i *DefaultEventHandler) Err() error {
	return i.err
}

func (i *DefaultEventHandler) GetResponse() proto.Message {
	return i.response
}

func (i *DefaultEventHandler) Reset() {
	i.headers = nil
	i.trailers = nil
	i.response = nil
	i.err = nil
}

var _ InvocationEventHandler = &StreamEventHandler{}

type StreamEventHandler struct {
	headers  metadata.MD
	trailers metadata.MD
	response chan proto.Message
	err      error
}

func NewStreamEventHandler() *StreamEventHandler {
	return &StreamEventHandler{
		response: make(chan proto.Message, 100),
	}
}

func (i *StreamEventHandler) OnResolveMethod(*desc.MethodDescriptor) {}

func (i *StreamEventHandler) OnSendHeaders(metadata.MD) {}

func (i *StreamEventHandler) OnReceiveHeaders(headers metadata.MD) {
	i.headers = headers
}

func (i *StreamEventHandler) OnReceiveResponse(message proto.Message) {
	i.response <- message
}

func (i *StreamEventHandler) OnReceiveTrailers(s *status.Status, trailers metadata.MD) {
	if err := s.Err(); err != nil {
		i.err = err
	}

	i.trailers = trailers
}

func (i *StreamEventHandler) GetHeaders() metadata.MD {
	return i.headers
}

func (i *StreamEventHandler) GetTrailers() metadata.MD {
	return i.trailers
}

func (i *StreamEventHandler) Err() error {
	return i.err
}

func (i *StreamEventHandler) GetResponse() proto.Message {
	return <-i.response
}

func (i *StreamEventHandler) GetResponses() chan proto.Message {
	return i.response
}

// TODO
func (i *StreamEventHandler) Reset() {}
