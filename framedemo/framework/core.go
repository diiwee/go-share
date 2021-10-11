package framework

import (
	"log"
	"net/http"
	"strings"
)

type FrameHandler func(c *Context) error

type Core struct {
	router      map[string]*Tree
	middlewares []FrameHandler
}

// NewCore 初始化框架核心
func NewCore() *Core {
	router := map[string]*Tree{}
	router["GET"] = NewTree()
	router["PUT"] = NewTree()
	router["POST"] = NewTree()
	router["DELETE"] = NewTree()
	return &Core{
		router: router,
	}
}

// Use 注册中间件
func (c *Core) Use(middlewares ...FrameHandler) {
	c.middlewares = middlewares
}

// Group 初始化路由组
func (c *Core) Group(prefix string) IGroup {
	return NewGroupRouter(prefix, c)
}

// Get 方法
func (c *Core) Get(url string, handlers ...FrameHandler) {
	allHandlers := append(c.middlewares, handlers...)
	if err := c.router["GET"].AddRouter(url, allHandlers); err != nil {
		log.Fatal("add router error: ", err)
	}
}

// Post 方法
func (c *Core) Post(url string, handlers ...FrameHandler) {
	allHandlers := append(c.middlewares, handlers...)
	if err := c.router["POST"].AddRouter(url, allHandlers); err != nil {
		log.Fatal("add router error: ", err)
	}
}

// Put 方法
func (c *Core) Put(url string, handlers ...FrameHandler) {
	allHandlers := append(c.middlewares, handlers...)
	if err := c.router["PUT"].AddRouter(url, allHandlers); err != nil {
		log.Fatal("add router error: ", err)
	}
}

// Delete 方法
func (c *Core) Delete(url string, handlers ...FrameHandler) {
	allHandlers := append(c.middlewares, handlers...)
	if err := c.router["DELETE"].AddRouter(url, allHandlers); err != nil {
		log.Fatal("add router error: ", err)
	}
}

// FindRouter 通过请求匹配路由
func (c *Core) FindRouter(req *http.Request) []FrameHandler {
	uri := req.URL.Path
	method := strings.ToUpper(req.Method)

	if m, ok := c.router[method]; ok {
		return m.FindHandler(uri)
	}
	return nil
}

func (c *Core) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	ctx := NewContext(w, req)

	handlers := c.FindRouter(req)

	if handlers == nil {
		ctx.Json(http.StatusNotFound, "Not Found")
		return
	}

	ctx.SetHandlers(handlers)

	if err := ctx.Next(); err != nil {
		ctx.Json(http.StatusInternalServerError, "Internal Error")
		return
	}
}
