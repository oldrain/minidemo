package controllers

import (
	"github.com/oldrain/minigo"
	"minidemo/app/controllers/v1"
	"minidemo/lib/middleware"
)

func InitCtl(router *minigo.Api) {
	// middleware
	router.Use(middleware.LogInput)

	// root
	// ..

	// v1
	v1Router := router.Group("/v1")
	user := v1Router.Group("/user")
	user.Post("/info", v1.UerInfo)
	user.LastUse(middleware.LogOutput)

	// v2
	// ...
}
