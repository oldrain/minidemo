package v1

import (
	"github.com/oldrain/minigo"
	"minidemo/app/logics"
	"minidemo/constant"
	"minidemo/dto/entity"
	"minidemo/dto/input"
	"minidemo/lib"
	"minidemo/models"
)

func UerInfo(ctx *minigo.Context) {
	var logger = lib.ContextLogger(ctx)

	userIdIn := new(input.UserIdIn)
	err := ctx.BindJSON(userIdIn)
	if err != nil {
		logger.Error(err)
		ctx.Error(constant.ErrorSys, lib.CfgErrMsg(constant.ErrorSys))
		return
	}

	err = new(minigo.Validate).Do(userIdIn)
	if err != nil {
		logger.Error(err)
		ctx.Error(constant.ErrorIn, err.Error())
		return
	}

	var userEntity = new(entity.UserEntity)
	err = models.UserById(userIdIn.Id, userEntity)
	if err != nil {
		logger.Error(err)
		ctx.Error(constant.ErrorIn, err.Error())
		return
	}

	ctx.JSON(logics.RenderUserInfo(userEntity))
}
