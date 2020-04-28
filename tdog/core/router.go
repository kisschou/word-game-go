package core

import (
	"net/http"
	"path"

	"github.com/julienschmidt/httprouter"
)

type (
	// Handlerfunc.
	HandlerFunc func(*Context)

	// Context.
	Context struct {
		Req     *http.Request
		Writer  http.ResponseWriter
		Params  httprouter.Params
		Handler HandlerFunc
		Engine  *HttpEngine
	}

	// RouterGroup.
	RouterGroup struct {
		Handler HandlerFunc
		Prefix  string
		Parent  *RouterGroup
		Engine  *HttpEngine
	}

	// HttpEngine.
	HttpEngine struct {
		*RouterGroup
		Router *httprouter.Router
	}
)

func New() *HttpEngine {
	engine := &HttpEngine{}
	engine.RouterGroup = &RouterGroup{nil, "", nil, engine}
	engine.Router = httprouter.New()
	return engine
}

func (engine *HttpEngine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	engine.Router.ServeHTTP(w, req)
}

func (engine *HttpEngine) Run(addr string) {
	http.ListenAndServe(addr, engine)
}

func (group *RouterGroup) createContext(w http.ResponseWriter, req *http.Request, params httprouter.Params, handler HandlerFunc) *Context {
	return *Context{
		Writer:  w,
		Req:     req,
		Engine:  group.Engine,
		Params:  params,
		Handler: handler,
	}
}

func (group *RouterGroup) Group(component string) *RouterGroup {
	prefix := path.Join(group.Prefix, component)
	return &RouterGroup{
		Handler: nil,
		Parent:  group,
		Prefix:  prefix,
		Engine:  group.Engine,
	}
}

func (group *RouterGroup) Handle(method, p string, handler HandlerFunc) {
	p = path.Join(group.Prefix, p)
	group.Engine.Router.Handle(method, p, func(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
		group.createContext(w, req, params, handler).Next()
	})
}

// POST is a shortcut for router.Handle("POST", path, handle)
func (group *RouterGroup) POST(path string, handler HandlerFunc) {
	group.Handle("POST", path, handler)
}

// GET is a shortcut for router.Handle("GET", path, handle)
func (group *RouterGroup) GET(path string, handler HandlerFunc) {
	group.Handle("GET", path, handler)
}

// DELETE is a shortcut for router.Handle("DELETE", path, handle)
func (group *RouterGroup) DELETE(path string, handler HandlerFunc) {
	group.Handle("DELETE", path, handler)
}

// PATCH is a shortcut for router.Handle("PATCH", path, handle)
func (group *RouterGroup) PATCH(path string, handler HandlerFunc) {
	group.Handle("PATCH", path, handler)
}

// PUT is a shortcut for router.Handle("PUT", path, handle)
func (group *RouterGroup) PUT(path string, handler HandlerFunc) {
	group.Handle("PUT", path, handler)
}

// Next .
func (c *Context) Next() {
	c.Handler(c)
}

// Writes the given string into the response body and sets the Content-Type to "text/plain"
func (c *Context) String(code int, msg string) {
	c.Writer.Header().Set("Content-Type", "text/plain")
	c.Writer.WriteHeader(code)
	c.Writer.Write([]byte(msg))
}
