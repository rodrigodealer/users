package handlers

import (
	"net/http"

	"github.com/rodrigodealer/users/mysql"
	"github.com/rodrigodealer/users/redis"
)

func TokenHandler(redis redis.RedisConn, mysql mysql.MySQLConn) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// var header = "x-user-token"
		// if token not present in header
		// if token present and its in redis
		// var key = redis.Get("mykey")
		// log.Printf(key)

	}
}
