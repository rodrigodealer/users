package handlers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnauthorizedToken(t *testing.T) {
	req, _ := http.NewRequest("GET", "/users/token", nil)
	res := httptest.NewRecorder()
	redisClient := new(redisMock)
	mysqlClient := new(mysqlMock)

	handler := http.HandlerFunc(TokenHandler(redisClient, mysqlClient))
	handler.ServeHTTP(res, req)

	assert.Equal(t, res.Code, 401)
	assert.Equal(t, res.Body.String(), "{\"status\":\"ERROR\",\"message\":\"Token expired or invalid\"}\n")
}

func TestSuccessfulTokenFromRedis(t *testing.T) {
	req, _ := http.NewRequest("GET", "/users/token?token=val", nil)
	res := httptest.NewRecorder()
	redisClient := new(redisMock)
	mysqlClient := new(mysqlMock)
	redisClient.On("Get").Return("bla", nil)

	handler := http.HandlerFunc(TokenHandler(redisClient, mysqlClient))
	handler.ServeHTTP(res, req)

	assert.Equal(t, 200, res.Code)
	assert.Equal(t, "\"bla\"\n", res.Body.String())
}

func TestSuccessfulTokenFromMySQL(t *testing.T) {
	req, _ := http.NewRequest("GET", "/users/token?token=val", nil)
	res := httptest.NewRecorder()
	redisClient := new(redisMock)
	mysqlClient := new(mysqlMock)
	redisClient.On("Get").Return("", nil)
	mysqlClient.On("GetToken").Return("myval", nil)

	handler := http.HandlerFunc(TokenHandler(redisClient, mysqlClient))
	handler.ServeHTTP(res, req)

	assert.Equal(t, 200, res.Code)
	assert.Equal(t, "\"myval\"\n", res.Body.String())
}

func TestUnsuccessfulTokenFromMySQL(t *testing.T) {
	req, _ := http.NewRequest("GET", "/users/token?token=val", nil)
	res := httptest.NewRecorder()
	redisClient := new(redisMock)
	mysqlClient := new(mysqlMock)
	redisClient.On("Get").Return("", nil)
	mysqlClient.On("GetToken").Return("", errors.New("bla"))

	handler := http.HandlerFunc(TokenHandler(redisClient, mysqlClient))
	handler.ServeHTTP(res, req)

	assert.Equal(t, res.Code, 404)
	assert.Equal(t, "\"\"\n", res.Body.String())
}
