package lib

import (
	"github.com/spf13/viper"
	"minidemo/utils"
)

func InitCfg(path, file string) error {
	viper.AddConfigPath(path)
	viper.SetConfigType("yml")
	viper.SetConfigName(file)
	return viper.ReadInConfig()
}

func CfgBindObj(key string, obj interface{}) (err error) {
	err = utils.JsonToObj(utils.ObjToJson(viper.Get(key)), obj)
	return
}

func CfgGetString(key string) string {
	return viper.GetString(key)
}

func CfgErrMsg(err int) string {
	return CfgGetString(string(err))
}
