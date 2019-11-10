package logics

import (
	"minidemo/constant"
	"minidemo/dto/entity"
	"minidemo/dto/output"
	"minidemo/lib/db"
	"minidemo/utils"
)

func RenderUserInfo(entity *entity.UserEntity) *output.UserInfoOut {
	if entity == nil {
		return nil
	}
	userInfoOut := new(output.UserInfoOut)
	userInfoOut.Id = entity.ID
	userInfoOut.Name = entity.Name
	userInfoOut.Email = entity.Email
	var key = utils.FormatString(constant.RedisKeyUserDailyView, utils.NowDateClose(), entity.ID)
	count, _ := db.GetTestSingleRedisClient().Get(key).Int()
	userInfoOut.DailyView = count
	return userInfoOut
}
