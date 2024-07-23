package db

import "github.com/xochat/TangSengDaoDaoServerLib/pkg/redis"

func NewRedis(addr string, password string, db int) *redis.Conn {
	return redis.New(addr, password, db)
}
