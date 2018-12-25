package redis

import (
	"fmt"
	"log"
	"time"

	"github.com/garyburd/redigo/redis"
)

type RedisPool struct {
	rp *redis.Pool
}

func NewRedisPool(host, port, auth string) (*RedisPool, error) {

	addr := fmt.Sprintf("%s:%s", host, port)
	log.Println(addr)

	RedisClient := &redis.Pool{
		MaxIdle: 30,
		//MaxActive:   30,
		IdleTimeout: 60 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", addr)
			if err != nil {
				log.Println("redis open fail:", err)
				return nil, err
			}
			if auth != "" {
				if _, err := c.Do("AUTH", auth); err != nil {
					c.Close()
					return nil, err
				}
			}
			// 选择db
			//c.Do("SELECT", REDIS_DB)
			return c, nil
		},
	}
	//懒加载
	return &RedisPool{rp: RedisClient}, nil
}

func (ts *RedisPool) GetRedis() redis.Conn {
	rc := ts.rp.Get()
	return rc
}
func (ts *RedisPool) CloseRedis(rc redis.Conn) {
	rc.Close()
}

func Test(ts *RedisPool) error {
	rc := ts.rp.Get()
	_, err := redis.String(rc.Do("get", "key1"))
	rc.Close()
	if err != nil {
		return err
	}
	return nil
}
