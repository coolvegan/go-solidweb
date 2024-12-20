package main

import (
	"fmt"
	"net/http"
)

func NewAppContext() *appContextHandler {
	mux := http.NewServeMux()
	preMiddleware := make([]http.Handler, 0, 5)
	postMiddleware := make([]http.Handler, 0, 5)
	appCtx := &appContextHandler{str: "hiho", router: mux}
	appCtx.preMiddleware = append(preMiddleware, &debugMiddleware{appCtx: appCtx}, &loginMiddleware{appCtx: appCtx}, &headerMiddleware{appCtx: appCtx})
	appCtx.postMiddleware = append(postMiddleware, &afterVisitRouteMiddleware{appCtx: appCtx})
	return appCtx
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	//Für den Fall, dass wir jede 404 redirecten wollen
	//http.Redirect(w, r, "/", http.StatusPermanentRedirect)
	http.NotFound(w, r)
}

func (a appContextHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.log.Printf("App Design")
	fmt.Fprintf(w, "HandlerVariante über ServeHTTP %s", a.str)
}

func (a *appContextHandler) MuxFuncHandler(w http.ResponseWriter, r *http.Request) {
	a.log.Printf("Method: %s Path: %s", r.Method, r.Pattern)
	fmt.Fprintf(w, "HandlerVariante über mux.HandleFunc %s", a.str)
}

// Ggf. es soll 404 allgemein abgefangen werden
func NotFoundHandler() http.Handler { return http.HandlerFunc(NotFound) }
