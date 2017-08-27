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
			tokenNotPresent(w)
		} else {
			var key = redis.Get(token[0])
			if key != "" {
				json.NewEncoder(w).Encode(key)
			} else {
				key, err := mysql.GetToken(token[0])
				if err != nil {
					log.Printf("Error fetching token %s from MySQL: %s", token[0], err.Error())
					w.WriteHeader(http.StatusNotFound)
				} else {
					log.Printf("Found token %s in MySQL", token[0])
					redis.SetXX(token[0], key)
					json.NewEncoder(w).Encode(key)
				}
			}
			// w.WriteHeader(http.StatusOK)
		}
	}
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
