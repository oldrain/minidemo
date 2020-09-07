package entry

import (
	"github.com/oldrain/minigo"
	v1 "minidemo/app/controller/v1"
	"minidemo/lib/middleware"
)

func InitV1(router *minigo.Api) {
	// v1
	v1Router := router.Group("/v1")

	// user
	user := v1Router.Group("/user")
	user.Post("/info", v1.UerInfo)
	user.LastUse(middleware.LogOutput)

}
