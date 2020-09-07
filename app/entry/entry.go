package entry

import (
	"github.com/oldrain/minigo"
	"minidemo/app/controller"
	"minidemo/lib/middleware"
)

func InitEntry(router *minigo.Api) {
	// middleware
	router.Use(middleware.LogInput)

	// root
	router.Post("/test", middleware.NeedLogin, controller.Test)

	// v1
	InitV1(router)
}
