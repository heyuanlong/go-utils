package redis

import (
	"fmt"
	"log"
	"time"

	"github.com/garyburd/redigo/redis"

	klog "github.com/heyuanlong/go-utils/common/log"
)

func NewRedisPool(host, port, auth string) (*redis.Pool, error) {

	addr := fmt.Sprintf("%s:%s", host, port)
	klog.Info.Println(addr)

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
	return RedisClient, nil
}

func GetRedis(RedisPool *redis.Pool) redis.Conn {
	rc := RedisPool.Get()
	return rc
}
func CloseRedis(rc redis.Conn) {
	rc.Close()
}

func Test(RedisPool *redis.Pool) error {
	rc := RedisPool.Get()
	_, err := redis.String(rc.Do("get", "key1"))
	rc.Close()
	if err != nil {
		return err
	}
	return nil
}
