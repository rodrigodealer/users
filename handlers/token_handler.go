package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/rodrigodealer/users/mysql"
	"github.com/rodrigodealer/users/redis"
)

func TokenHandler(redis redis.RedisConn, mysql mysql.MySQLConn) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := r.URL.Query()
		token, isTokenPresent := vars["token"]
		if !isTokenPresent {
			tokenNotPresent(w)
		} else {
			tokenFromMySQL(w, token[0], redis, mysql)
		}
	}
}

func tokenFromMySQL(w http.ResponseWriter, token string,
	redis redis.RedisConn, mysql mysql.MySQLConn) {
	// var key = redis.Get("bla")
	key := redis.Get(token)
	if key == "" {
		var dbToken, err = mysql.GetToken(token)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
		} else {
			key = dbToken
			redis.SetXX(token, dbToken)
		}
	}
	json.NewEncoder(w).Encode(key)
}

func tokenNotPresent(w http.ResponseWriter) {
	var response = &TokenError{Status: "ERROR", Message: "Token expired or invalid"}
	w.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(w).Encode(response)
}

type TokenError struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
