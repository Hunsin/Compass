package router

import (
	"context"
	"net/http"
	gpath "path"

	"github.com/julienschmidt/httprouter"
)

// Param returns the named URL parameter from a request context.
func Param(ctx context.Context, name string) string {
	if p := httprouter.ParamsFromContext(ctx); p != nil {
		return p.ByName(name)
	}
	return ""
}

// A Middleware chains http.Handlers.
type Middleware func(http.Handler) http.Handler

// A Router is a http.Handler which supports routing and middlewares.
type Router struct {
	middlewares []Middleware
	path        string
	root        *httprouter.Router
}

// New creates a new Router.
func New() *Router {
	return &Router{root: httprouter.New(), path: "/"}
}

// Group returns a new Router with given path and middlewares.
// It should be used for handlers which have same path prefix or
// common middlewares.
func (r *Router) Group(path string, m ...Middleware) *Router {
	return &Router{
		middlewares: append(m, r.middlewares...),
		path:        gpath.Join(r.path, path),
		root:        r.root,
	}
}

// Use appends new middlewares to current Router.
func (r *Router) Use(m ...Middleware) *Router {
	r.middlewares = append(m, r.middlewares...)
	return r
}

// Handle registers a new request handler combined with middlewares.
func (r *Router) Handle(method, path string, handler http.Handler) {
	for _, v := range r.middlewares {
		handler = v(handler)
	}
	r.root.Handler(method, gpath.Join(r.path, path), handler)
}

// GET is a shortcut for r.Handle("GET", path, handler)
func (r *Router) GET(path string, handler http.HandlerFunc) {
	r.Handle(http.MethodGet, path, handler)
}

// HEAD is a shortcut for r.Handle("HEAD", path, handler)
func (r *Router) HEAD(path string, handler http.HandlerFunc) {
	r.Handle(http.MethodHead, path, handler)
}

// OPTIONS is a shortcut for r.Handle("OPTIONS", path, handler)
func (r *Router) OPTIONS(path string, handler http.HandlerFunc) {
	r.Handle(http.MethodOptions, path, handler)
}

// POST is a shortcut for r.Handle("POST", path, handler)
func (r *Router) POST(path string, handler http.HandlerFunc) {
	r.Handle(http.MethodPost, path, handler)
}

// PUT is a shortcut for r.Handle("PUT", path, handler)
func (r *Router) PUT(path string, handler http.HandlerFunc) {
	r.Handle(http.MethodPut, path, handler)
}

// PATCH is a shortcut for r.Handle("PATCH", path, handler)
func (r *Router) PATCH(path string, handler http.HandlerFunc) {
	r.Handle(http.MethodPatch, path, handler)
}

// DELETE is a shortcut for r.Handle("DELETE", path, handler)
func (r *Router) DELETE(path string, handler http.HandlerFunc) {
	r.Handle(http.MethodDelete, path, handler)
}

// HandleFunc is an adapter for http.HandlerFunc.
func (r *Router) HandleFunc(method, path string, handler http.HandlerFunc) {
	r.Handle(method, path, handler)
}

// NotFound sets the handler which is called if the request path doesn't match
// any routes. It overwrites the previous setting.
func (r *Router) NotFound(handler http.Handler) {
	r.root.NotFound = handler
}

// Static serves files from given root directory.
func (r *Router) Static(path, root string) {
	if len(path) < 10 || path[len(path)-10:] != "/*filepath" {
		panic("path should end with '/*filepath' in path '" + path + "'.")
	}

	base := gpath.Join(r.path, path[:len(path)-9])
	fileServer := http.StripPrefix(base, http.FileServer(http.Dir(root)))

	r.Handle(http.MethodGet, path, fileServer)
}

// File serves the named file.
func (r *Router) File(path, name string) {
	r.HandleFunc(http.MethodGet, path, func(w http.ResponseWriter, req *http.Request) {
		http.ServeFile(w, req, name)
	})
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.root.ServeHTTP(w, req)
}
