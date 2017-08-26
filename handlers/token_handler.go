package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/rodrigodealer/users/mysql"
	"github.com/rodrigodealer/users/redis"
)

func TokenHandler(redis redis.RedisConn, mysql mysql.MySQLConn) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := r.URL.Query()
		token, isTokenPresent := vars["token"]
		if !isTokenPresent {
			var response = &TokenError{Status: "ERROR", Message: "Token expired or invalid"}
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(response)
		} else {
			log.Printf("Token for: %s", token[0])
			w.WriteHeader(http.StatusOK)
		}

	}
}

type TokenError struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
