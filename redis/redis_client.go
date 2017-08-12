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
}

func (r *RedisConnection) Connect() {
	var client, err = redis.Dial("tcp", ":6379")

	if err != nil {
		log.Printf("Redis error: %s", err.Error())
	}
	r.Conn = client
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
