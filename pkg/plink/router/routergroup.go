package router

import (
	"chatcser/pkg/plink/iface"
	"strconv"
	"strings"
)

// IRouteGroup interface for router group
type IRouteGroup interface {
	IRoutes
	//GroupWithNone(prefix string) IRouteGroup
	GroupWithMore(prefix string, middleware ...iface.Middleware) IRouteGroup
	Group(prefix string) IRouteGroup
}

// IRoutes interface for routes
type IRoutes interface {
	Use(...iface.Middleware)
	// Any(string, iface.HandlerFunc)
	// Get(string, iface.HandlerFunc)
	Post(string, iface.HandlerFunc)
	// Delete(string, iface.HandlerFunc)
	// Patch(string, iface.HandlerFunc)
	// Put(string, iface.HandlerFunc)
	// Options(string, iface.HandlerFunc)
	// Head(string, iface.HandlerFunc)
	// Connect(string, iface.HandlerFunc)
	// Trace(string, iface.HandlerFunc)
}

// routeGroup struct containing all fields and methods for use.
type routeGroup struct {
	prefix     string
	middleware []iface.Middleware
	pure       *Router
}

var _ iface.IRouter = &routeGroup{}

func (g *routeGroup) handle(method string, path string, handler iface.HandlerFunc) {

	if i := strings.Index(path, "//"); i != -1 {
		panic("Bad path '" + path + "' contains duplicate // at index:" + strconv.Itoa(i))
	}

	h := handler

	for i := len(g.middleware) - 1; i >= 0; i-- {
		h = g.middleware[i](h)
	}

	tree := g.pure.trees[method]

	if tree == nil {
		tree = new(node)
		g.pure.trees[method] = tree
	}

	pCount := tree.add(g.prefix+path, h)
	pCount++

	if pCount > g.pure.mostParams {
		g.pure.mostParams = pCount
	}
}

// Use adds a middleware handler to the group middleware chain.
func (g *routeGroup) Use(m ...iface.Middleware) {
	g.middleware = append(g.middleware, m...)
}

// Connect adds a CONNECT route & handler to the router.
func (g *routeGroup) Connect(path string, h iface.HandlerFunc) {
	g.handle(MethodConnect, path, h)
}

// Delete adds a DELETE route & handler to the router.
func (g *routeGroup) Delete(path string, h iface.HandlerFunc) {
	g.handle(MethodDelete, path, h)
}

// Get adds a GET route & handler to the router.
func (g *routeGroup) Get(path string, h iface.HandlerFunc) {
	g.handle(MethodGet, path, h)
}

// Head adds a HEAD route & handler to the router.
func (g *routeGroup) Head(path string, h iface.HandlerFunc) {
	g.handle(MethodHead, path, h)
}

// Options adds an OPTIONS route & handler to the router.
func (g *routeGroup) Options(path string, h iface.HandlerFunc) {
	g.handle(MethodOptions, path, h)
}

// Patch adds a PATCH route & handler to the router.
func (g *routeGroup) Patch(path string, h iface.HandlerFunc) {
	g.handle(MethodPatch, path, h)
}

// Post adds a POST route & handler to the router.
func (g *routeGroup) Post(path string, h iface.HandlerFunc) {
	g.handle(MethodPost, path, h)
}

// Put adds a PUT route & handler to the router.
func (g *routeGroup) Put(path string, h iface.HandlerFunc) {
	g.handle(MethodPut, path, h)
}

// Trace adds a TRACE route & handler to the router.
func (g *routeGroup) Trace(path string, h iface.HandlerFunc) {
	g.handle(MethodTrace, path, h)
}

// Handle allows for any method to be registered with the given
// route & handler. Allows for non standard methods to be used
// like CalDavs PROPFIND and so forth.
func (g *routeGroup) Handle(method string, path string, h iface.HandlerFunc) {
	g.handle(method, path, h)
}

// Any adds a route & handler to the router for all HTTP methods.
func (g *routeGroup) Any(path string, h iface.HandlerFunc) {
	g.Connect(path, h)
	g.Delete(path, h)
	g.Get(path, h)
	g.Head(path, h)
	g.Options(path, h)
	g.Patch(path, h)
	g.Post(path, h)
	g.Put(path, h)
	g.Trace(path, h)
}

// Match adds a route & handler to the router for multiple HTTP methods provided.
func (g *routeGroup) Match(methods []string, path string, h iface.HandlerFunc) {
	for _, m := range methods {
		g.handle(m, path, h)
	}
}

// GroupWithNone creates a new sub router with specified prefix and no middleware attached.
func (g *routeGroup) GroupWithNone(prefix string) iface.IRouter {
	return &routeGroup{
		prefix:     g.prefix + prefix,
		pure:       g.pure,
		middleware: make([]iface.Middleware, 0),
	}
}

// GroupWithMore creates a new sub router with specified prefix, retains existing middleware and adds new middleware.
func (g *routeGroup) GroupWithMore(prefix string, middleware ...iface.Middleware) iface.IRouter {
	rg := &routeGroup{
		prefix:     g.prefix + prefix,
		pure:       g.pure,
		middleware: make([]iface.Middleware, len(g.middleware)),
	}
	copy(rg.middleware, g.middleware)
	rg.Use(middleware...)
	return rg
}

// Group creates a new sub router with specified prefix and retains existing middleware.
func (g *routeGroup) Group(prefix string) iface.IRouter {
	rg := &routeGroup{
		prefix:     g.prefix + prefix,
		pure:       g.pure,
		middleware: make([]iface.Middleware, len(g.middleware)),
	}
	copy(rg.middleware, g.middleware)
	return rg
}

func (g *routeGroup) GetHandlerWithUrl(url string) iface.HandlerFunc {
	tree := g.pure.trees["POST"]
	h, _ := tree.Find(url, g.pure)
	return h
}
