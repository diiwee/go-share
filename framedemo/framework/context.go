package framework

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type Context struct {
	request        *http.Request
	responseWriter http.ResponseWriter
	ctx            context.Context

	isTimeOut bool

	*sync.Mutex
}

// NewContext 初始化Context
func NewContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		responseWriter: w,
		request:        req,
		ctx:            req.Context(),
		isTimeOut:      false,
		Mutex:          &sync.Mutex{},
	}
}

//Context基本功能

func (ctx *Context) GetRequest() *http.Request {
	return ctx.request
}

func (ctx *Context) GetResponseWriter() http.ResponseWriter {
	return ctx.responseWriter
}

func (ctx *Context) SetTimeout() {
	ctx.isTimeOut = true
}

func (ctx *Context) GetIsTimeout() bool {
	return ctx.isTimeOut
}

//实现标准库的context

func (ctx *Context) FrameContext() context.Context {
	return ctx.request.Context()
}

func (ctx *Context) Deadline() (deadline time.Time, ok bool) {
	return ctx.FrameContext().Deadline()
}

func (ctx *Context) Done() <-chan struct{} {
	return ctx.FrameContext().Done()
}

func (ctx *Context) Err() error {
	return ctx.FrameContext().Err()
}

func (ctx *Context) Value(key interface{}) interface{} {
	return ctx.FrameContext().Value(key)
}

//封装Request 路由参数和表单参数

func (ctx *Context) QueryAll() map[string][]string {
	if ctx.request != nil {
		return ctx.request.URL.Query()
	}

	return map[string][]string{}
}

func (ctx *Context) QueryStringDef(key string, def string) string {

	params := ctx.QueryAll()

	if v, ok := params[key]; ok {
		if len(v) > 0 {
			return v[len(v)-1]
		}
	}

	return def
}

func (ctx *Context) QueryIntDef(key string, def int) int {

	params := ctx.QueryAll()

	if v, ok := params[key]; ok {
		if len(v) > 0 {
			i, err := strconv.Atoi(v[len(v)-1])
			if err != nil {
				return def
			}
			return i
		}
	}

	return def
}

func (ctx *Context) FormAll() map[string][]string {
	if ctx.request != nil {
		return map[string][]string(ctx.request.PostForm)
	}
	return map[string][]string{}
}

func (ctx *Context) FormIntDef(key string, def int) int {
	params := ctx.FormAll()
	if values, ok := params[key]; ok {
		len := len(values)
		if len > 0 {
			intval, err := strconv.Atoi(values[len-1])
			if err != nil {
				return def
			}
			return intval
		}
	}
	return def
}

func (ctx *Context) FormStringDef(key string, def string) string {
	params := ctx.FormAll()
	if values, ok := params[key]; ok {
		len := len(values)
		if len > 0 {
			return values[len-1]
		}
	}
	return def
}

// 封装responseWriter

func (ctx *Context) Json(status int, obj interface{}) error {

	if ctx.GetIsTimeout() {
		return nil
	}

	ctx.responseWriter.Header().Set("Content-Type", "application/json")
	ctx.responseWriter.WriteHeader(status)

	byt, err := json.Marshal(obj)
	if err != nil {
		ctx.responseWriter.WriteHeader(500)
		return err
	}
	ctx.responseWriter.Write(byt)
	return nil
}
