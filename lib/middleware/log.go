package middleware

import (
	"github.com/oldrain/minigo"
	"minidemo/constant"
	"minidemo/lib"
	"minidemo/utils"
)

func LogInput(ctx *minigo.Context) {
	var logger = lib.ContextLogger(ctx)

	logger.Erase()

	logId := utils.UUIDPure()
	logger.AppendHead(utils.FormatString("%s%s%s", "[", logId, "]"))
	ctx.Set(constant.LogContext, logger)
	ctx.SetOutHeader(constant.ContextHeaderLogID, logId)

	logger.Info(ctx.Input.Request.Header)
	logger.Info(ctx.GetInBody())

	// next middleware
	ctx.Continue()
}

func LogOutput(ctx *minigo.Context) {
	var logger = lib.ContextLogger(ctx)
	logger.Info(ctx.GetOutBody())
}
