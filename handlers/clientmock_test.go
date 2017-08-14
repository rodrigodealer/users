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

func (o redisMock) Connect() {
}

type mysqlMock struct {
	mock.Mock
}

func (o mysqlMock) Ping() (bool, error) {
	args := o.Called()
	return args.Bool(0), nil
}

func (o mysqlMock) Connect() {
}
