package constant

const (
	// Redis type
	RedisTypeSingle  = "single"
	RedisTypeCluster = "cluster"

	// Expire
	RedisExpireFlexRate = 1.5
	RedisExpire5m       = 300
	RedisExpire30m      = 1800
	RedisExpire1h       = 3600
	RedisExpire2h       = 7200
	RedisExpire1d       = 86400
	RedisExpire1w       = 604800

	// Prefix
	RedisPrefix = "%s:%s:"

	// Redis key
	RedisKeyUserDailyView = "userDailyView:%s:%d"

)
