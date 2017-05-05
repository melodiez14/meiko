package conn

import (
	"log"
	"time"

	"github.com/garyburd/redigo/redis"
)

type RedisConfig struct {
	Address string `json:"address"`
}

var Redis *redis.Pool

func InitRedis(cfg RedisConfig) {
	log.Println("Initializing Redis")

	Redis = &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 10 * time.Second,
		Dial:        func() (redis.Conn, error) { return redis.Dial("tcp", cfg.Address) },
	}

	conn := Redis.Get()
	_, err := conn.Do("PING")
	if err != nil {
		log.Fatalln("Redis connection failed")
		return
	}

	log.Println("Redis successfully connected")
}
