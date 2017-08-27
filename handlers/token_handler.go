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
			var key = redis.Get(token[0])
			if key == "" {
				var dbToken, err = mysql.GetToken(token[0])
				if err != nil {
					w.WriteHeader(http.StatusNotFound)
				} else {
					key = dbToken
					redis.SetXX(token[0], dbToken)
				}
			}
			json.NewEncoder(w).Encode(key)
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
