package cache

import (
	"minidemo/constant"
	"minidemo/lib"
	"minidemo/lib/db"
	"minidemo/util"
	"time"
)

func GetAppRedisClient() *db.RedisClient {
	return db.NewRedisClient(constant.CfgRedisApp)
}

// test client
func GetTestSingleRedisClient() *db.RedisClient {
	return db.NewRedisClient(constant.CfgRedisTestSingle)
}

// test client
func GetTestClusterRedisClient() *db.RedisClient {
	return db.NewRedisClient(constant.CfgRedisTestCluster)
}

func GetKeyExpire(exp int) time.Duration {
	return time.Second * time.Duration(util.RandomNumber(exp, int(float64(exp)*constant.RedisExpireFlexRate)))
}

func CreateCacheKey(key string, args ...interface{}) string {
	prefix := util.FormatString(constant.RedisPrefix, lib.CfgMeshPartnerId(), lib.CfgMeshAppId())
	return util.FormatString(prefix + key, args...)
}

func SetJson(client *db.RedisClient, key string, obj interface{}, exp time.Duration) {
	client.Set(key, util.ObjToJson(obj), exp)
}

// fyi: obj signed to nil
func GetObj(client *db.RedisClient, key string, obj interface{}) error {
	s := client.Get(key).Val()
	err := util.JsonToObj(s, obj)
	if err != nil {
		obj = nil
	}
	return err
}
