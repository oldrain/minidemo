package v1

import (
	"github.com/oldrain/minigo"
	"minidemo/app/logic"
	"minidemo/constant"
	"minidemo/dto/input"
	"minidemo/lib"
	"minidemo/model/usersmodel"
)

func UerInfo(ctx *minigo.Context) {
	var logger = lib.ContextLogger(ctx)

	in := new(input.UserIdIn)
	err := ctx.BindJSON(in)
	if err != nil {
		logger.Error(err)
		ctx.Error(constant.ErrorSys, lib.CfgErrMsg(constant.ErrorSys))
		return
	}

	err = new(minigo.Validate).Do(in)
	if err != nil {
		logger.Error(err)
		ctx.Error(constant.ErrorIn, err.Error())
		return
	}

	var userModel = usersmodel.UserGetById(in.Id)
	if userModel == nil {
		logger.Error(err)
		ctx.Error(constant.ErrorIn, err.Error())
		return
	}

	ctx.JSON(logic.RenderUserInfo(userModel))
}
