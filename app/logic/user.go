package logic

import (
	"minidemo/constant"
	"minidemo/dto/output"
	"minidemo/lib/cache"
	"minidemo/model/usersmodel"
	"minidemo/util"
)

func RenderUserInfo(user *usersmodel.User) *output.UserInfoOut {
	if user == nil {
		return nil
	}
	userInfoOut := new(output.UserInfoOut)
	userInfoOut.Id = user.ID
	userInfoOut.Name = user.Name
	userInfoOut.Email = user.Email
	var key = util.FormatString(constant.RedisKeyUserDailyView, util.NowDateClose(), user.ID)
	count, _ := cache.GetAppRedisClient().Get(key).Int()
	userInfoOut.DailyView = count
	return userInfoOut
}
