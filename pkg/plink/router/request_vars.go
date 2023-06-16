package router

import (
	"context"
)

// ReqVars is the interface of request scoped variables
// tracked by pure
type ReqVars interface {
	URLParam(pname string) string
}

type requestVars struct {
	ctx    context.Context // holds a copy of it's parent requestVars
	params urlParams
	//formParsed bool
}

// Params returns the current routes Params
func (r *requestVars) URLParam(pname string) string {
	return r.params.Get(pname)
}

type urlParam struct {
	key   string
	value string
}

type urlParams []urlParam

// Get returns the URL parameter for the given key, or blank if not found
func (p urlParams) Get(key string) (param string) {
	for i := 0; i < len(p); i++ {
		if p[i].key == key {
			param = p[i].value
			return
		}
	}
	return
}

// Middleware is pure's middleware definition
//type Middleware func(h iface.HandlerFunc) iface.HandlerFunc
