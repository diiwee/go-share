package main

import (
	"framedemo/framework"
	"framedemo/framework/middleware"
	"net/http"
)

func main() {

	core := framework.NewCore()

	core.Use(middleware.Cost())
	//注册路由
	resistRouter(core)

	server := &http.Server{
		Addr:    ":8080",
		Handler: core,
	}

	server.ListenAndServe()

}
