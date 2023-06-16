package router

import (
	"chatcser/pkg/plink/iface"
	"context"
	"sync"
)

/*
	路由接口， 这里面路由是 使用框架者给该链接自定的 处理业务方法
	路由里的IRequest 则包含用该链接的链接信息和该链接的请求数据信息
*/

// Mux is the main request multiplexer
type Router struct {
	routeGroup
	trees map[string]*node

	// pool is used for reusable request scoped RequestVars content
	pool sync.Pool

	// http404     iface.HandlerFunc // 404 Not Found
	// http405     iface.HandlerFunc // 405 Method Not Allowed
	// httpOPTIONS iface.HandlerFunc

	// mostParams used to keep track of the most amount of
	// params in any URL and this will set the default capacity
	// of each Params
	mostParams uint8

	// Enables automatic redirection if the current route can't be matched but a
	// handler for the path with (without) the trailing slash exists.
	// For example if /foo/ is requested but a route only exists for /foo, the
	// client is redirected to /foo with http status code 301 for GET requests
	// and 307 for all other request methods.
	redirectTrailingSlash bool

	// If enabled, the router checks if another method is allowed for the
	// current route, if the current request can not be routed.
	// If this is the case, the request is answered with 'Method Not Allowed'
	// and HTTP status code 405.
	// If no other Method is allowed, the request is delegated to the NotFound
	// handler.
	handleMethodNotAllowed bool

	// if enabled automatically handles OPTION requests; manually configured OPTION
	// handlers take presidence. default true
	automaticallyHandleOPTIONS bool
}

// New Creates and returns a new Pure instance
func NewRouter() *Router {
	p := &Router{
		routeGroup: routeGroup{
			middleware: make([]iface.Middleware, 0),
		},
		trees:                      make(map[string]*node),
		mostParams:                 0,
		redirectTrailingSlash:      true,
		handleMethodNotAllowed:     false,
		automaticallyHandleOPTIONS: false,
	}
	p.routeGroup.pure = p
	p.pool.New = func() interface{} {

		rv := &requestVars{
			params: make(urlParams, p.mostParams),
		}

		rv.ctx = context.WithValue(context.Background(), defaultContextIdentifier, rv)

		return rv
	}
	return p
}

// func (m *Router) GetHandler(req *iface.Request) iface.HandlerFunc {
// 	tree := m.trees[req.Method]
// 	h, _ := tree.Find(req.URL.Path, m)
// 	return h
// }

func (m *Router) GetHandlerWithUrl(url string) iface.HandlerFunc {
	tree := m.trees["POST"]
	h, _ := tree.Find(url, m)
	return h
}
