package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuccessfulToken(t *testing.T) {
	req, _ := http.NewRequest("GET", "/healthcheck", nil)
	res := httptest.NewRecorder()
	redisClient := new(redisMock)
	mysqlClient := new(mysqlMock)

	redisClient.On("Ping").Return(true, nil)
	mysqlClient.On("Ping").Return(true, nil)

	handler := http.HandlerFunc(HealthcheckHandler(redisClient, mysqlClient))
	handler.ServeHTTP(res, req)

	assert.Equal(t, res.Code, 200)
	assert.Equal(t, res.Body.String(), "{\"Status\":\"WORKING\",\"Services\":[{\"Working\":true,\"Service\":\"Redis\"},{\"Working\":true,\"Service\":\"MySQL\"}]}\n")
}
