package hxxp

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Context struct {
	Writer  http.ResponseWriter
	Request *http.Request
	status  int
}

type HandlerFunc func(ctx *Context)

type Router struct {
	chi chi.Router
}

func (r *Router) Get(pattern string, handler HandlerFunc) {
	r.build(http.MethodGet, pattern, handler)
}

func (r *Router) Post(pattern string, handler HandlerFunc) {
	r.build(http.MethodPost, pattern, handler)
}

func (r *Router) Put(pattern string, handler HandlerFunc) {
	r.build(http.MethodPut, pattern, handler)
}

func (r *Router) Patch(pattern string, handler HandlerFunc) {
	r.build(http.MethodPatch, pattern, handler)
}

func (r *Router) Delete(pattern string, handler HandlerFunc) {
	r.build(http.MethodDelete, pattern, handler)
}

func (r *Router) Options(pattern string, handler HandlerFunc) {
	r.build(http.MethodOptions, pattern, handler)
}

func (r *Router) build(method string, pattern string, handler HandlerFunc) {
	r.chi.Method(method, pattern,
		http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			ctx := &Context{
				Writer:  writer,
				Request: request,
			}
			handler(ctx)
		}),
	)
}

func (r *Router) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	r.chi.ServeHTTP(writer, request)
}
