// Code generated by hertz generator.

package main

import (
	handler "hertz_demo/biz/handler"

	"github.com/cloudwego/hertz/pkg/app/server"
)

// customizeRegister registers customize routers.
func customizedRegister(r *server.Hertz) {
	r.GET("/ping", handler.Ping)
	ui := handler.NewUserImpl()
	// 区分路由组，以v1开头
	rg := r.Group("/v1")
	rg.POST("/login", ui.Login)
	rg.POST("/logout", ui.LogOut)
	rg.GET("/users", ui.GetUsers)
}
