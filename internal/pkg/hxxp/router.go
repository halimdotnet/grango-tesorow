package hxxp

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type HandlerFunc func(ctx *Context)

type Router struct {
	chi chi.Router
}

func (r *Router) Use(middlewares ...func(http.Handler) http.Handler) {
	r.chi.Use(middlewares...)
}

func (r *Router) Get(pattern string, handler HandlerFunc) {
	r.register(http.MethodGet, pattern, handler)
}

func (r *Router) Post(pattern string, handler HandlerFunc) {
	r.register(http.MethodPost, pattern, handler)
}

func (r *Router) Put(pattern string, handler HandlerFunc) {
	r.register(http.MethodPut, pattern, handler)
}

func (r *Router) Patch(pattern string, handler HandlerFunc) {
	r.register(http.MethodPatch, pattern, handler)
}

func (r *Router) Delete(pattern string, handler HandlerFunc) {
	r.register(http.MethodDelete, pattern, handler)
}

func (r *Router) Options(pattern string, handler HandlerFunc) {
	r.register(http.MethodOptions, pattern, handler)
}

func (r *Router) register(method string, pattern string, handler HandlerFunc) {
	r.chi.Method(method, pattern,
		http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			ctx := &Context{
				Writer:  writer,
				Request: request,
				Ctx:     request.Context(),
			}
			handler(ctx)
		}),
	)
}

func (r *Router) notFound(writer http.ResponseWriter, request *http.Request) {
	ctx := &Context{
		Writer:  writer,
		Request: request,
		Ctx:     request.Context(),
	}

	ctx.Response(http.StatusNotFound, Response{
		Error:   true,
		Message: "Resource Not Found",
	})
}

func (r *Router) methodNotAllowed(writer http.ResponseWriter, request *http.Request) {
	ctx := &Context{
		Writer:  writer,
		Request: request,
		Ctx:     request.Context(),
	}

	ctx.Response(http.StatusMethodNotAllowed, Response{
		Error:   true,
		Message: "Method Not Allowed",
	})
}

func (r *Router) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	r.chi.ServeHTTP(writer, request)
}
