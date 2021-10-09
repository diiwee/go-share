package main

import (
	"framedemo/framework"
	"net/http"
)

func main() {

	core := framework.NewCore()
	core.Get("ping", TestHandler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: core,
	}

	server.ListenAndServe()

}
