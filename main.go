package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rodrigodealer/users/handlers"
	"github.com/rodrigodealer/users/mysql"
	"github.com/rodrigodealer/users/redis"
	"github.com/rodrigodealer/zipkin-tracing/tracing"
)

func main() {

	log.SetOutput(os.Stdout)
	log.Print("Starting server on port 8080")
	err := http.ListenAndServe(":8080", tracing.TrackerHandler(Server()))
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
	tracing.StartTracing("localhost:8080", "users")

	r.HandleFunc("/healthcheck", handlers.HealthcheckHandler(redis, mysql)).Name("/healthcheck").Methods("GET")
	r.HandleFunc("/users/token", handlers.TokenHandler(redis, mysql)).Name("/users/token").Methods("GET")
	return r
}
