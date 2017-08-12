package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/rodrigodealer/users/mysql"
	"github.com/rodrigodealer/users/redis"
)

type HealthcheckServiceStatus struct {
	Working bool
	Service string
}

type HealthcheckStatus struct {
	Status   string
	Services []*HealthcheckServiceStatus
}

func HealthcheckHandler(redis redis.RedisConn, mysql mysql.MySQLConn) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		healthcheckStatus := &HealthcheckStatus{Status: "WORKING"}
		healthcheckStatus.Services = append(healthcheckStatus.Services, redisCheck(redis))
		healthcheckStatus.Services = append(healthcheckStatus.Services, mysqlCheck(mysql))

		for _, service := range healthcheckStatus.Services {
			if !service.Working {
				healthcheckStatus.Status = "FAILED"
				w.WriteHeader(http.StatusInternalServerError)
			}
		}

		json.NewEncoder(w).Encode(healthcheckStatus)
	}
}

func redisCheck(redis redis.RedisConn) *HealthcheckServiceStatus {
	redis.Connect()
	var working, _ = redis.Ping()
	return &HealthcheckServiceStatus{Working: working, Service: "Redis"}
}

func mysqlCheck(mysql mysql.MySQLConn) *HealthcheckServiceStatus {
	var working, _ = mysql.Ping()
	return &HealthcheckServiceStatus{Working: working, Service: "MySQL"}
}
