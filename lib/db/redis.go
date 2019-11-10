package db

import (
	"github.com/go-redis/redis"
	"minidemo/constant"
	"minidemo/dto"
	"minidemo/lib"
	"strings"
	"time"
)

type RedisClient struct {
	tp string
	single *redis.Client
	cluster *redis.ClusterClient
}

var redisConnections = make(map[string]*RedisClient)

// test client
func GetTestSingleRedisClient() *RedisClient {
	return NewRedisClient(constant.ConfigRedisTestSingle)
}

// test client
func GetTestClusterRedisClient() *RedisClient {
	return NewRedisClient(constant.ConfigRedisTestCluster)
}

func NewRedisClient(conn string) *RedisClient {
	if exists, ok := redisConnections[conn]; ok && (exists != nil) {
		return exists
	}

	var cfg = new(dto.RedisCfg)
	err := lib.CfgBindObj(conn, cfg)
	if err != nil {
		lib.GetLogger(constant.LogRedis).Error(err)
	}

	var client = &RedisClient{
		tp: constant.RedisTypeSingle,
	}

	if cfg.Type == constant.RedisTypeCluster {
		client.tp = constant.RedisTypeCluster
		client.cluster = RedisClusterClient(cfg)
	} else {
		client.single = RedisSingleClient(cfg)
	}

	_, err = client.Ping()
	if err != nil {
		lib.GetLogger(constant.LogRedis).Error(err)
	}

	redisConnections[conn] = client

	return client
}

func RedisSingleClient(cfg *dto.RedisCfg) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: cfg.Addr,
		Password: cfg.Password,
		DB: cfg.Db,
		PoolSize: cfg.PoolSize,
		PoolTimeout: cfg.PoolTimeout * time.Second,
		IdleTimeout: cfg.IdleTimeout * time.Second,
		MaxRetries: cfg.MaxRetries,
	})
}

func RedisClusterClient(cfg *dto.RedisCfg) *redis.ClusterClient {
	return redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: strings.Split(cfg.Addr, ","),
		Password: cfg.Password,
		PoolSize: cfg.PoolSize,
		PoolTimeout: cfg.PoolTimeout * time.Second,
		IdleTimeout: cfg.IdleTimeout * time.Second,
		MaxRedirects: cfg.MaxRedirects,
	})
}

// Delegate commands

func (client *RedisClient) Ping() (string, error) {
	if client.isCluster(client.tp) {
		return client.cluster.Ping().Result()
	}

	return client.single.Ping().Result()
}

func (client *RedisClient) Set(key string, value interface{}, exp time.Duration) *redis.StatusCmd {
	if client.isCluster(client.tp) {
		return client.cluster.Set(key, value, exp)
	}

	return client.single.Set(key, value, exp)
}

func (client *RedisClient) Get(key string) *redis.StringCmd {
	if client.isCluster(client.tp) {
		return client.cluster.Get(key)
	}

	return client.single.Get(key)
}

// ...

func (client *RedisClient) isCluster(tp string) bool {
	return constant.RedisTypeCluster == tp
}
