package lib

import (
	"github.com/oldrain/golog"
	"github.com/oldrain/minigo"
	"strings"
	"minidemo/constant"
	"minidemo/util"
)

func ContextLogger(ctx *minigo.Context) golog.Logger {
	if value, ok := ctx.Get(constant.LogContext); ok {
		logger, _ := value.(golog.Logger)
		return logger
	}

	var path = ctx.Path()

	if path == constant.ContextPathRoot {
		path = CfgGetString(constant.CfgAppName)
	} else {
		path = strings.Trim(path, "/")
		path = strings.ReplaceAll(path, "/", " ")
		path = strings.Title(path)
		path = strings.ReplaceAll(path, " ", "")
		path = util.LcFirst(path)
	}

	return GetLogger(path)
}

func GetLogger(moduleName string) golog.Logger {
	env := CfgGetString(constant.CfgAppEnv)
	if env == constant.EnvDev {
		cfg := new(golog.Config)
		cfg.SetLevel(golog.LevelAll)
		return golog.ConsoleLogger(moduleName, cfg)
	}
	return golog.GetFileLogger(moduleName)
}
