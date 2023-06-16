package router4

import (
	"net/http"
	"strings"
)

type HandlerFunc func(http.ResponseWriter, *http.Request, map[string]string)

type Router struct {
	routes map[string]HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		routes: make(map[string]HandlerFunc),
	}
}

func (r *Router) AddRoute(method string, path string, handler HandlerFunc) {
	key := method + "_" + path
	r.routes[key] = handler
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for k, v := range r.routes {
		parts := strings.Split(k, "_")
		if len(parts) != 2 {
			continue
		}
		if req.Method == parts[0] {
			if matchPath(req.URL.Path, parts[1]) {
				vars := extractVars(req.URL.Path, parts[1])
				v(w, req, vars)
				return
			}
		}
	}
	http.NotFound(w, req)
}

func matchPath(path string, pattern string) bool {
	pathParts := strings.Split(path, "/")
	patternParts := strings.Split(pattern, "/")

	if len(pathParts) != len(patternParts) {
		return false
	}

	for i, part := range patternParts {
		if part != "" && part[0] == ':' {
			continue
		}
		if pathParts[i] != part {
			return false
		}
	}

	return true
}

func extractVars(path string, pattern string) map[string]string {
	vars := make(map[string]string)

	pathParts := strings.Split(path, "/")
	patternParts := strings.Split(pattern, "/")

	for i, part := range patternParts {
		if part != "" && part[0] == ':' {
			vars[part[1:]] = pathParts[i]
		}
	}

	return vars
}

//func helloHandler(w http.ResponseWriter, req *http.Request, vars map[string]string) {
//	name := "World"
//	if val, ok := vars["name"]; ok {
//		name = val
//	}
//	fmt.Fprintf(w, "Hello, %s!", name)
//}

//func main() {
//	r := NewRouter()
//	r.AddRoute("GET", "/", func(w http.ResponseWriter, req *http.Request, vars map[string]string) {
//		fmt.Fprint(w, "Welcome to my website!")
//	})
//	r.AddRoute("GET", "/hello/:name", helloHandler)
//
//	http.ListenAndServe(":8000", r)
//}
