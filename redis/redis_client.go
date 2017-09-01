package redis

import (
	"log"
	"time"

	"github.com/go-redis/redis"
)

type RedisConnection struct {
	Conn *redis.Client
}

type RedisConn interface {
	Connect()
	Ping() (bool, error)
	Get(key string) string
	SetXX(key string, value string)
}

func (r *RedisConnection) Connect() {
	r.Conn = r.getPool()
}

func (r *RedisConnection) getPool() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:        ":6379",
		PoolSize:    3000,
		PoolTimeout: 1,
		Password:    "",
		DB:          0,
	})

	return client
}

func (r *RedisConnection) Get(key string) string {
	var result, err = r.Conn.Get(key).Result()
	if err != nil {
		log.Printf("Error trying to get key: %s\n %s", key, err.Error())
	}
	return result
}

func (r *RedisConnection) SetXX(key string, value string) {
	const ttl = 120
	var _, err = r.Conn.SetNX(key, value, ttl*time.Second).Result()
	if err != nil {
		log.Printf("Error trying to set key: %s\n %s", key, err.Error())
	} else {
		log.Printf("Set key %s with TTL %d", key, ttl)
	}
}

func (r *RedisConnection) Ping() (bool, error) {
	var redisWorking = false
	if r.Conn != nil {
		var _, err = r.Conn.Ping().Result()
		if err != nil {
			log.Printf("Error performing ping: %s", err.Error())
		}
		redisWorking = true
		return redisWorking, err
	}
	return redisWorking, nil
}
