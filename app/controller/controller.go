package controller

import (
	"github.com/oldrain/golog"
	"github.com/oldrain/minigo"
	"minidemo/lib"
)

func Test(ctx *minigo.Context) {
	ctx.JSON(nil)
}

func BindAndValidate(ctx *minigo.Context, in interface{}) error {
	err := ctx.BindJSON(in)
	if err != nil {
		return err
	}
	return new(minigo.Validate).Do(in)
}

func JsonWithError(ctx *minigo.Context, logger golog.Logger, errNo int, err error) {
	logger.Error(err)
	ctx.Error(errNo, lib.CfgErrMsg(errNo))
}
