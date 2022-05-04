package wsjsong

//go:generate protoc -I. -I$GOPATH/src --go_out=. --go_opt=paths=source_relative protocol/protocol.proto
//go:generate protoc -I. -I$GOPATH/src --go_out=. --go-grpc_out=. --openapi_out==paths=source_relative:. protocol/logic/logic.proto
