package framework

type IGroup interface {
	Get(string, ...FrameHandler)
	Post(string, ...FrameHandler)
	Delete(string, ...FrameHandler)
	Put(string, ...FrameHandler)

	// Group 嵌套
	Group(string) IGroup
}

type GroupRouter struct {
	core        *Core
	prefix      string
	parent      *GroupRouter
	middlewares []FrameHandler
}

func (g *GroupRouter) Use(middlewares ...FrameHandler) {
	g.middlewares = middlewares
}

func NewGroupRouter(prefix string, core *Core) *GroupRouter {
	return &GroupRouter{
		prefix:      prefix,
		core:        core,
		parent:      nil,
		middlewares: []FrameHandler{},
	}
}

// 获取当前group的绝对路径
func (g *GroupRouter) getAbsolutePrefix() string {
	if g.parent == nil {
		return g.prefix
	}
	return g.parent.getAbsolutePrefix() + g.prefix
}

// 获取某个group的middleware
// 这里就是获取除了Get/Post/Put/Delete之外设置的middleware
func (g *GroupRouter) getMiddlewares() []FrameHandler {
	if g.parent == nil {
		return g.middlewares
	}

	return append(g.parent.getMiddlewares(), g.middlewares...)
}

func (g *GroupRouter) Get(uri string, handlers ...FrameHandler) {
	uri = g.getAbsolutePrefix() + uri
	allHandlers := append(g.getMiddlewares(), handlers...)
	g.core.Get(uri, allHandlers...)
}

func (g *GroupRouter) Post(uri string, handlers ...FrameHandler) {
	uri = g.getAbsolutePrefix() + uri
	allHandlers := append(g.getMiddlewares(), handlers...)
	g.core.Post(uri, allHandlers...)
}

func (g *GroupRouter) Delete(uri string, handlers ...FrameHandler) {
	uri = g.getAbsolutePrefix() + uri
	allHandlers := append(g.getMiddlewares(), handlers...)
	g.core.Delete(uri, allHandlers...)
}

func (g *GroupRouter) Put(uri string, handlers ...FrameHandler) {
	uri = g.getAbsolutePrefix() + uri
	allHandlers := append(g.getMiddlewares(), handlers...)
	g.core.Put(uri, allHandlers...)
}

func (g *GroupRouter) Group(uri string) IGroup {
	cgroup := NewGroupRouter(uri, g.core)
	cgroup.parent = g
	return cgroup
}
