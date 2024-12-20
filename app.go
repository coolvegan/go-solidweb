package main

import "net/http"

type appContextHandler struct {
	str            string
	router         *http.ServeMux
	log            Log
	preMiddleware  []http.Handler
	postMiddleware []http.Handler
}

func (a *appContextHandler) run(address string) {
	a.registerRoutes()

	a.log.Printf("Running Server: %s", address)
	err := http.ListenAndServe(address, a.router)
	if err != nil {
		a.log.Printf(err.Error())
	}
}
