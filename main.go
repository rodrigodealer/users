package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rodrigodealer/realtime/tracing"
	"github.com/rodrigodealer/users/handlers"
	"github.com/rodrigodealer/users/mysql"
	"github.com/rodrigodealer/users/redis"
)

func main() {

	log.SetOutput(os.Stdout)
	tracing.StartTracing("localhost:8080", "users")
	log.Print("Starting server on port 8080")
	err := http.ListenAndServe(":8080", Server())
	if err != nil {
		log.Panic("Something is wrong : " + err.Error())
	}
}

func Server() http.Handler {
	redis := &redis.RedisConnection{}
	redis.Connect()
	mysql := &mysql.MySQLConnection{}
	mysql.Connect()

	r := mux.NewRouter()
	r.HandleFunc("/healthcheck", handlers.HealthcheckHandler(redis, mysql)).Name("/healthcheck").Methods("GET")
	return r
}
