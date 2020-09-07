package middleware

import (
	"minidemo/model/usersmodel"
	"time"

	"github.com/oldrain/minigo"

	"minidemo/constant"
	"minidemo/lib"
	"minidemo/util"
)

func NeedLogin(ctx *minigo.Context) {
	var logger = lib.ContextLogger(ctx)

	var tokenStr = ctx.GetInHeader(constant.HeaderToken)
	if tokenStr == "" {
		ctx.AbortWithError(constant.ErrorNeedLogin, lib.CfgErrMsg(constant.ErrorNeedLogin))
		return
	}

	jwtToken := &util.JwtToken{
		Token: tokenStr,
	}
	jwtClaims := new(util.JwtClaims)

	// parse
	err := util.JwtParse(jwtToken, jwtClaims, lib.CfgGetString(constant.CfgJwtSecretKey))
	if err != nil {
		ctx.AbortWithError(constant.ErrorNeedLogin, lib.CfgErrMsg(constant.ErrorNeedLogin))
		return
	}

	// check
	tokenModel := usersmodel.CustomerTokenGetByCode(jwtClaims.CustomerCode)
	if tokenModel == nil {
		logger.Error(util.FormatString("token not in storage: %s, claims: %s",  err, util.ObjToJson(jwtClaims)))
		ctx.AbortWithError(constant.ErrorNeedLogin, lib.CfgErrMsg(constant.ErrorNeedLogin))
		return
	}

	now := time.Now().Unix()
	if tokenModel.LoginTime.Unix() - now < lib.CfgGetInt64(constant.CfgJwtExpireSec) {
		logger.Error(util.FormatString("login time expired, claims: %s", util.ObjToJson(jwtClaims)))
		ctx.AbortWithError(constant.ErrorLoginExpired, lib.CfgErrMsg(constant.ErrorLoginExpired))
		return
	}

	if jwtClaims.VerifyExpiresAt(now, true) {
		logger.Error(util.FormatString("jwt expired, claims: %s", util.ObjToJson(jwtClaims)))
		ctx.AbortWithError(constant.ErrorLoginExpired, lib.CfgErrMsg(constant.ErrorLoginExpired))
		return
	}

	ctx.SetInHeader(constant.HeaderPartnerId, jwtClaims.PartnerId)
	ctx.SetInHeader(constant.HeaderAppId, jwtClaims.AppId)
	ctx.SetInHeader(constant.HeaderCustomerCode, jwtClaims.CustomerCode)
	ctx.SetOutHeader(constant.HeaderToken, jwtToken.Token)

	ctx.Continue()
}
