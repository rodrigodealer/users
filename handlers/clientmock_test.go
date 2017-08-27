package handlers

import (
	"github.com/stretchr/testify/mock"
)

type redisMock struct {
	mock.Mock
}

func (o redisMock) Ping() (bool, error) {
	args := o.Called()
	return args.Bool(0), nil
}

func (o redisMock) Get(key string) string {
	args := o.Called()
	return args.String(0)
}

func (o redisMock) SetXX(key string, value string) {
}

func (o redisMock) Connect() {
}

type mysqlMock struct {
	mock.Mock
}

func (o mysqlMock) Ping() (bool, error) {
	args := o.Called()
	return args.Bool(0), args.Error(1)
}

func (o mysqlMock) GetToken(token string) (string, error) {
	args := o.Called()
	return args.String(0), args.Error(1)
}

func (o mysqlMock) Connect() {
}
