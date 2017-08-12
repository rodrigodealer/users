package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/gorilla/mux"
)

func main() {
	var client, err = redis.Dial("tcp", ":6379")

	if err != nil {
		log.Printf("Redis error: %s", err.Error())
	}

	var ping, _ = client.Do("PING")
	log.Print(ping)

	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	http.Handle("/", r)

	srv := &http.Server{
		Handler: r,
		Addr:    ":8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category: %v\n", vars["category"])
}
