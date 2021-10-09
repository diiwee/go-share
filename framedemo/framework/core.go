package framework

import (
	"log"
	"net/http"
)

type FrameHandler func(c *Context) error

type Core struct {
	router map[string]FrameHandler
}

// NewCore 初始化框架核心
func NewCore() *Core {
	return &Core{
		router: map[string]FrameHandler{},
	}
}

func (c *Core) Get(url string, handler FrameHandler) {
	c.router[url] = handler
}

func (c *Core) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	log.Println("framework.serveHTTP")
	ctx := NewContext(w, req)

	//写死测试一下
	router := c.router["ping"]

	if router == nil {
		return
	}
	log.Println("framework.router")
	router(ctx)
}
