package main

import (
	"net/http"
)

func (a *appContextHandler) middlewareRunner(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, mi := range a.preMiddleware {
			if mi != nil {

				mi.ServeHTTP(w, r)
			}
		}
		next.ServeHTTP(w, r)
		for _, mi := range a.postMiddleware {
			if mi != nil {

				mi.ServeHTTP(w, r)
			}
		}
	})
}

type loginMiddleware struct {
	appCtx *appContextHandler
}

func (l *loginMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	l.appCtx.log.Printf("LoginMiddleware: %s Path: %s\n", r.Method, r.Pattern)
}

type headerMiddleware struct {
	appCtx *appContextHandler
}

func (h *headerMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.appCtx.log.Printf("HeaderMiddleware: %s Path: %s\n", r.Method, r.Pattern)
}

type debugMiddleware struct {
	appCtx *appContextHandler
}

func (d *debugMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	d.appCtx.log.Printf("DebugMiddleware: %s Path: %s\n", r.Method, r.Pattern)
}

type afterVisitRouteMiddleware struct {
	appCtx *appContextHandler
}

func (l *afterVisitRouteMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	l.appCtx.log.Printf("AfterVisitRouteMiddleware: %s Path: %s\n", r.Method, r.Pattern)
}
