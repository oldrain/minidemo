package lib

import (
	"github.com/oldrain/golog"
	"github.com/oldrain/minigo"
	"minidemo/constant"
	"minidemo/utils"
	"strings"
)

func ContextLogger(ctx *minigo.Context) golog.Logger {
	var path = ctx.Path()

	if path == constant.ContextPathRoot {
		path = CfgGetString(constant.ConfigAppName)
	} else {
		path = strings.Trim(path, "/")
		path = strings.ReplaceAll(path, "/", " ")
		path = strings.Title(path)
		path = strings.ReplaceAll(path, " ", "")
		path = utils.LcFirst(path)
	}

	if value, ok := ctx.Get(path); ok {
		logger, _ := value.(golog.Logger)
		return logger
	}

	return GetLogger(path)
}

func GetLogger(moduleName string) golog.Logger {
	env := CfgGetString(constant.ConfigAppEnv)
	if env == constant.EnvDev {
		cfg := new(golog.Config)
		cfg.SetLevel(golog.LevelAll)
		return golog.ConsoleLogger(moduleName, cfg)
	}
	return golog.GetFileLogger(moduleName)
}
