package router

import (
	"context"
	"net/http"
	"regexp"
	"strings"
)

type Router struct {
	routes []route
}

type route struct {
	method    string
	path      *regexp.Regexp
	pathParam []string
	handler   http.HandlerFunc
}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) Add(method string, pattern string, handler http.HandlerFunc) *Router {
	newRoute := route{method, regexp.MustCompile("^" + pattern + "$"), nil, handler}
	r.routes = append(r.routes, newRoute)
	return r
}

type PathParamCtxKey struct{}

func isMatchPath(r *Router, req *http.Request) []route {
	var matches []route
	for _, route := range r.routes {
		match := route.path.FindStringSubmatch(req.URL.Path)
		if 0 < len(match) {
			route.pathParam = match[:1]
			matches = append(matches, route)
		}
	}
	return matches
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var allow []string
	matches := isMatchPath(r, req)
	if len(matches) == 0 {
		http.NotFound(w, req)
		return
	}
	for _, route := range matches {
		if route.method == req.Method {
			ctx := context.WithValue(req.Context(), PathParamCtxKey{}, route.pathParam)
			route.handler(w, req.WithContext(ctx))
			return
		}
		allow = append(allow, route.method)
	}
	w.Header().Set("Allow", strings.Join(allow, ", "))
	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
}

func PathParam(r *http.Request, index int) string {
	params := r.Context().Value(PathParamCtxKey{}).([]string)
	return params[index]
}
