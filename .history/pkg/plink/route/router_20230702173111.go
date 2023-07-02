package route

import (
	"errors"
	//"net/http"
)

// Router represents the router which handles routing.
type Router struct {
	Tree                    *tree
	NotFoundHandler         HandlerFun
	MethodNotAllowedHandler HandlerFun
	DefaultOPTIONSHandler   HandlerFun
	globalMiddlewares       middlewares
}

// route represents the route which has data for a routing.
type route struct {
	//methods     []string
	path        string
	handler     HandlerFun
	middlewares middlewares
}

var (
	tmpRoute = &route{}

	// Error for not found.
	ErrNotFound = errors.New("no matching route was found")
	// Error for method not allowed.
	ErrMethodNotAllowed = errors.New("methods is not allowed")
)

// NewRouter creates a new router.
func NewRouter() *Router {
	return &Router{
		Tree: map[string]*tree{},
	}
}

func (r *Router) UseGlobal(mws ...middleware) {
	nm := NewMiddlewares(mws)
	r.globalMiddlewares = nm
}

// Use sets middlewares.
func (r *Router) Use(mws ...middleware) *Router {
	nm := NewMiddlewares(mws)
	tmpRoute.middlewares = nm
	return r
}

// Use sets methods.
func (r *Router) Methods(methods ...string) *Router {
	tmpRoute.methods = append(tmpRoute.methods, methods...)
	return r
}
func (r *Router) Post() *Router {
	tmpRoute.methods = append(tmpRoute.methods, "POST")
	return r
}

// Handler sets a handler.
func (r Router) Handler(path string, handler HandlerFun) {
	tmpRoute.handler = handler
	tmpRoute.path = path
	r.Handle()
}

// Handle handles a route.
func (r *Router) Handle() {
	for i := 0; i < len(tmpRoute.methods); i++ {
		_, ok := r.Tree[tmpRoute.methods[i]]
		if !ok {
			r.Tree[tmpRoute.methods[i]] = newTree()
		}
		r.Tree[tmpRoute.methods[i]].Insert(tmpRoute.path, tmpRoute.handler, tmpRoute.middlewares)
	}
	tmpRoute = &route{}
}

// ServeHTTP dispatches the request to the handler whose
// pattern most closely matches the request URL.
// func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
// 	method := req.Method
// 	if method == http.MethodOptions {
// 		if r.DefaultOPTIONSHandler != nil {
// 			r.DefaultOPTIONSHandler.ServeHTTP(w, req)
// 			return
// 		}
// 	}

// 	t, ok := r.tree[method]
// 	if !ok {
// 		if r.MethodNotAllowedHandler == nil {
// 			methodNotAllowedHandler().ServeHTTP(w, req)
// 			return
// 		}
// 		r.MethodNotAllowedHandler.ServeHTTP(w, req)
// 		return
// 	}

// 	action, params, err := t.Search(req.URL.Path)
// 	if err == ErrNotFound {
// 		if r.NotFoundHandler == nil {
// 			http.NotFoundHandler().ServeHTTP(w, req)
// 			return
// 		}
// 		r.NotFoundHandler.ServeHTTP(w, req)
// 		return
// 	}

// 	h := action.handler
// 	// append globalMiddlewares to head of middlewares.
// 	mws := append(r.globalMiddlewares, action.middlewares...)
// 	if mws != nil {
// 		h = mws.then(h)
// 	}
// 	if params != nil {
// 		ctx := context.WithValue(req.Context(), ParamsKey, params)
// 		req = req.WithContext(ctx)
// 	}
// 	h.ServeHTTP(w, req)
// }

func (r *Router) Find(url string) HandlerFun {
	method := "POST" //req.Method

	t, ok := r.Tree[method]
	if !ok {
		if r.MethodNotAllowedHandler == nil {
			//methodNotAllowedHandler().ServeHTTP(w, req)
			return nil
		}
		//r.MethodNotAllowedHandler.ServeHTTP(w, req)
		return nil
	}

	action, params, err := t.Search(url)
	if err == ErrNotFound {
		if r.NotFoundHandler == nil {
			//http.NotFoundHandler().ServeHTTP(w, req)
			return nil
		}
		//r.NotFoundHandler.ServeHTTP(w, req)
		return nil
	}

	h := action.handler
	// append globalMiddlewares to head of middlewares.
	mws := append(r.globalMiddlewares, action.middlewares...)
	if mws != nil {
		h = mws.then(h)
	}
	if params != nil {
		//ctx := context.WithValue(req.Context(), ParamsKey, params)
		//req = req.WithContext(ctx)
	}
	return h
	//h(ctx)
}

// methodNotAllowedHandler is a default handler when status code is 405.
func methodNotAllowedHandler() HandlerFun {
	return HandlerFun(func(ctx any) {
		//w.WriteHeader(http.StatusMethodNotAllowed)
	})
}
