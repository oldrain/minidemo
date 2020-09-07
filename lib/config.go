package lib

import (
	"flag"
	"github.com/spf13/viper"
	"minidemo/constant"
	"minidemo/util"
)

func init() {
	// Environment
	env := *flag.String("env", "dev", "go run main.go -env=[dev/sit/uat/prod]")
	flag.Parse()
	err := InitCfg("config", env)
	if err != nil {
		panic(err)
	}
}

func InitCfg(path, file string) error {
	viper.AddConfigPath(path)
	viper.SetConfigType("yml")
	viper.SetConfigName(file)
	return viper.ReadInConfig()
}

// base

func CfgBindObj(key string, obj interface{}) (err error) {
	err = util.JsonToObj(util.ObjToJson(viper.Get(key)), obj)
	return
}

func CfgGetString(key string) string {
	return viper.GetString(key)
}

func CfgGetInt(key string) int {
	return viper.GetInt(key)
}

func CfgGetInt32(key string) int32 {
	return viper.GetInt32(key)
}

func CfgGetInt64(key string) int64 {
	return viper.GetInt64(key)
}

// customization

func CfgErrMsg(err int) string {
	return CfgGetString(util.FormatString(constant.CfgError, err))
}

func CfgClientNameMp() string {
	return CfgGetString(constant.AppClientNameMp)
}

func CfgMeshPartnerId() string {
	return CfgGetString(constant.CfgMeshPartnerId)
}

func CfgMeshAppId() string {
	return CfgGetString(constant.CfgMeshAppId)
}

func CfgMpAppId() string {
	return CfgGetString(constant.CfgMpAppId)
}

func CfgMpAppKey() string {
	return CfgGetString(constant.CfgMpAppKey)
}
