package dto

import "time"

type MysqlCfg struct {
	Url string
	MaxOpenConn int
	MaxIdleConn int
	ConnMaxLifeTime time.Duration
}

type RedisCfg struct {
	Addr string
	Password string
	Type string
	Db int
	PoolSize int
	PoolTimeout time.Duration
	IdleTimeout time.Duration
	MaxRetries int
	MaxRedirects int
}
