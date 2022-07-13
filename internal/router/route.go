package router

import (
	"net/http"
)

type (
	Route struct {
		Method  string
		Pattern string
		Handler http.HandlerFunc
	}

	Router struct {
		routes []Route
		mux    *http.ServeMux
	}

	Handler http.HandlerFunc
)

func (r *Router) AddRouter(method, path string, handler http.HandlerFunc) {
	r.routes = append(r.routes, Route{
		Method:  method,
		Pattern: path,
		Handler: handler,
	})
}

// NewRouter -.
func NewRouter() *Router {
	return &Router{
		routes: make([]Route, 0),
		mux:    http.NewServeMux(),
	}
}

func (r *Router) Mux() *http.ServeMux {
	// Set up 404 handler
	return r.mux
}

func (r *Router) Static(path string) *Router {
	r.mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(path))))
	return r
}

func (r *Router) GET(path string, handler http.HandlerFunc) *Router {
	r.mux.HandleFunc(path, func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == "GET" {
			handler(writer, request)
		}
	})
	return r
}

func (r *Router) POST(path string, handler http.HandlerFunc) *Router {
	r.mux.HandleFunc(path, func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == "POST" {
			handler(writer, request)
		}
	})
	return r
}

func (r *Router) PUT(path string, handler http.HandlerFunc) *Router {
	r.mux.HandleFunc(path, func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == "PUT" {
			handler(writer, request)
		}
	})
	return r
}

func (r *Router) DELETE(path string, handler http.HandlerFunc) *Router {
	r.mux.HandleFunc(path, func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == "DELETE" {
			handler(writer, request)
		}
	})
	return r
}
