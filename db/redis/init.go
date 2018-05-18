package redis

import (
	"fmt"
	"time"
	"github.com/garyburd/redigo/redis"

	kconf "github.com/heyuanlong/go-utils/common/conf"
	klog "github.com/heyuanlong/go-utils/common/log"
)

var (
	RedisClient     *redis.Pool
)

func init() {
	
}

func InitRedis()  {

	host,_ := kconf.GetString("redis","host")
	port,_ := kconf.GetString("redis","port")
	auth,_ := kconf.GetString("redis","auth")
	addr := fmt.Sprintf("%s:%s",host,port)
	klog.Info.Println(addr)

	RedisClient = &redis.Pool{
		MaxIdle:    1,
		MaxActive:   30,
		IdleTimeout: 1800 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", addr )
			if err != nil {
				klog.Warn.Fatal(err.Error())
				return nil, err
			}
			if _, err := c.Do("AUTH", auth); err != nil {
				c.Close()
				return nil, err
			}
			// 选择db
			//c.Do("SELECT", REDIS_DB)
			return c, nil
		},
	}
}

func GetRedis() redis.Conn  {
	rc := RedisClient.Get()
	return rc
}

func CloseRedis(rc redis.Conn )  {
	rc.Close()
}

func Test()  {
	rc := RedisClient.Get()
	v, err := redis.String(rc.Do("get", "key1"))
	rc.Close()
	if err !=nil {
		klog.Klog.Fatal(err.Error())
	}
	klog.Klog.Println(v)
}
