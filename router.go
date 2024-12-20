package main

import (
	"fmt"
	"net/http"
)

func (a *appContextHandler) registerRoutes() {
	a.router.Handle("/api/{$}", a)
	//Variante D über Closures
	a.router.Handle("/{$}", http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Willkommen")
		}))
	a.router.Handle("/api/v1", a.middlewareRunner(http.HandlerFunc(a.MuxFuncHandler)))
	//Jede undefinierte Route 404 zuweisen
	a.router.Handle("/", http.HandlerFunc(NotFound))

}
