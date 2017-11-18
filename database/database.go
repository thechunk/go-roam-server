package database

import (
	"github.com/garyburd/redigo/redis"
	"os"
	"time"
)

func newPool(addr string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial:        func() (redis.Conn, error) { return redis.DialURL(addr) },
	}
}

var (
	pool        *redis.Pool
	redisServer = os.Getenv("REDIS_URL")
)

func Conn() redis.Conn {
	if pool == nil {
		pool = newPool(redisServer)
	}
	return pool.Get()
}
