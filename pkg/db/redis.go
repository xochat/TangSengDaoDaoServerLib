package db

import "git.wkwork.xyz/im/TangSengDaoDaoServerLib/pkg/redis"

func NewRedis(addr string, password string, db int) *redis.Conn {
	return redis.New(addr, password, db)
}
