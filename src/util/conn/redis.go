package conn

import (
	"log"
	"time"

	"github.com/garyburd/redigo/redis"
)

type RedisConfig struct {
	Address  string `json:"address"`
	Password string `json:"password"`
}

var Redis *redis.Pool

func InitRedis(cfg RedisConfig) {
	log.Println("Initializing Redis")

	Redis = &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 10 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", cfg.Address)
			if err != nil || len(cfg.Password) < 1 {
				return c, err
			}

			c.Do("AUTH", cfg.Password)
			return c, err
		},
	}

	conn := Redis.Get()
	_, err := conn.Do("PING")
	if err != nil {
		log.Fatalln("Redis connection failed")
		return
	}

	log.Println("Redis successfully connected")
}
