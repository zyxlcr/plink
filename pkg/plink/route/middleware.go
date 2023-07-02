package route

type HandlerFun func(ctx any)

// middleware represents the singular of middleware.
type middleware func(HandlerFun) HandlerFun

// middlewares represents the plural of middleware.
type middlewares []middleware

// NewMiddlewares creates a new middlewares.
func NewMiddlewares(mws middlewares) middlewares {
	return append([]middleware(nil), mws...)
}

// then executes middlewares.
func (mws middlewares) then(h HandlerFun) HandlerFun {
	l := len(mws)
	for i := 0; i < l; i++ {
		h = mws[l-1-i](h)
	}
	return h
}
