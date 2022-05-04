package jsong

import (
	"sync"
)

// DefaultEventHandlerPool bytes.Buffer 对象池
var DefaultEventHandlerPool = sync.Pool{New: func() interface{} {
	return NewDefaultEventHandler()
}}

func GetDefaultEventHandler() InvocationEventHandler {
	return DefaultEventHandlerPool.Get().(InvocationEventHandler)
}

func PutDefaultEventHandler(handler InvocationEventHandler) {
	handler.Reset()
	DefaultEventHandlerPool.Put(handler)
}

// StreamEventHandlerPool bytes.Buffer 对象池
//var StreamEventHandlerPool = sync.Pool{New: func() interface{} {
//	return NewStreamEventHandler()
//}}
//
//func GetStreamEventHandler() InvocationEventHandler {
//	return DefaultEventHandlerPool.Get().(InvocationEventHandler)
//}
//
//func PutStreamEventHandler(handler InvocationEventHandler) {
//	handler.Reset()
//	StreamEventHandlerPool.Put(handler)
//}
