package core

import (
	"github.com/julienschmidt/httprouter"

	"net/http"
	"path"
	"strings"
)

type (
	// HandlerFunc .
	HandlerFunc func()

	// Context .
	Context struct {
		Req            *http.Request
		Writer         http.ResponseWriter
		Params         httprouter.Params
		handler        HandlerFunc
		engine         *HttpEngine
		BaseController *Controller
	}

	// RouterGroup .
	RouterGroup struct {
		Handler        HandlerFunc
		prefix         string
		parent         *RouterGroup
		engine         *HttpEngine
		BaseController *Controller
	}

	// HttpEngine .
	HttpEngine struct {
		*RouterGroup
		router *httprouter.Router
	}
)

// New HttpEngine
func NewEngine() *HttpEngine {
	engine := &HttpEngine{}
	engine.RouterGroup = &RouterGroup{nil, "", nil, engine, nil}
	engine.router = httprouter.New()
	return engine
}

// ServeHTTP makes the router implement the http.Handler interface.
func (engine *HttpEngine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	engine.router.ServeHTTP(w, req)
}

// Run .
func (engine *HttpEngine) Run(addr string) {
	http.ListenAndServe(addr, engine)
}

/************************************/
/********** ROUTES GROUPING *********/
/************************************/

func (group *RouterGroup) createContext(w http.ResponseWriter, req *http.Request, params httprouter.Params, handler HandlerFunc) *Context {
	if _, ok := req.Header["Content-Type"]; ok {
		if strings.Contains(req.Header["Content-Type"][0], "x-www-form-urlencoded") {
			req.ParseForm()
		}

		if strings.Contains(req.Header["Content-Type"][0], "form-data") {
			req.ParseMultipartForm(32 << 20)
		}
	}

	return &Context{
		Writer:         w,
		Req:            req,
		engine:         group.engine,
		Params:         params,
		handler:        handler,
		BaseController: group.BaseController,
	}
}

// Group .
func (group *RouterGroup) Group(component string, baseController *Controller) *RouterGroup {
	prefix := path.Join(group.prefix, component)
	return &RouterGroup{
		Handler:        nil,
		parent:         group,
		prefix:         prefix,
		engine:         group.engine,
		BaseController: baseController,
	}
}

// Handle .
func (group *RouterGroup) Handle(method, p string, handler HandlerFunc) {
	p = path.Join(group.prefix, p)
	group.engine.router.Handle(method, p, func(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
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

// OPTIONS
func (group *RouterGroup) OPTIONS(path string, handler HandlerFunc) {
	group.Handle("OPTIONS", path, handler)
}

// Next .
func (c *Context) Next() {
	var req Request
	req.Recv(c)
	c.handler()
}

// Writes the given string into the response body and sets the Content-Type to "text/plain"
func (c *Context) String(code int, msg string) {
	c.Writer.Header().Set("Content-Type", "text/plain")
	c.Writer.WriteHeader(code)
	c.Writer.Write([]byte(msg))
}
