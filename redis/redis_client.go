package redis

import (
	"log"

	"github.com/garyburd/redigo/redis"
)

type RedisConnection struct {
	Conn redis.Conn
}

type RedisConn interface {
	Connect()
	Ping() (bool, error)
	Get(key string) string
	SetXX(key string, value string)
}

func (r *RedisConnection) Connect() {
	var client, err = redis.Dial("tcp", ":6379")

	if err != nil {
		log.Printf("Redis error: %s", err.Error())
	}
	r.Conn = client
}

func (r *RedisConnection) Get(key string) string {
	var result, err = redis.String(r.Conn.Do("GET", key))
	if err != nil {
		log.Printf("Error trying to get key: %s\n %s", key, err.Error())
	}
	return result
}

func (r *RedisConnection) SetXX(key string, value string) {
	const ttl = 120
	var _, err = r.Conn.Do("SETEX", key, ttl, value)
	if err != nil {
		log.Printf("Error trying to set key: %s\n %s", key, err.Error())
	} else {
		log.Printf("Set key %s with TTL %d", key, ttl)
	}
}

func (r *RedisConnection) Ping() (bool, error) {
	var redisWorking = false
	if r.Conn != nil {
		var _, err = r.Conn.Do("PING")
		if err != nil {
			log.Printf("Error performing ping: %s", err.Error())
		}
		redisWorking = true
		return redisWorking, err
	}
	return redisWorking, nil
}
