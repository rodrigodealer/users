package redis

import (
	"log"
	"time"

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
	var client, err = redis.Dial("tcp", "127.0.0.1:6379")

	if err != nil {
		log.Printf("Redis error: %s", err.Error())
	}
	r.Conn = client
}

func (r *RedisConnection) getPool() *redis.Pool {
	pool := &redis.Pool{
		MaxIdle:     50,              // Maximum number of idle connections in the pool.
		MaxActive:   100,             // Maximum number of connections allocated by the pool at a given time.
		IdleTimeout: 2 * time.Second, // Close connections after remaining idle for this duration. Applications should set the timeout to a value less than the server's timeout.
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "127.0.0.1:6379")
			if err != nil {
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
	return pool
}

func (r *RedisConnection) Get(key string) string {
	var pool = r.getPool().Get()
	var result, err = redis.String(r.getPool().Get().Do("GET", key))
	if err != nil {
		log.Printf("Error trying to get key: %s\n %s", key, err.Error())
	}
	defer pool.Close()
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
