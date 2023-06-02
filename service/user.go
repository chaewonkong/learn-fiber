package service

import (
	"context"
	"fmt"
)

type (
	// User 서비스
	User interface {
		Greeting(c context.Context, name string) (resp string, err error)
	}

	user struct {
	}
)

// NewUserService 서비스 생성자
func NewUserService() User {
	return &user{}
}

func (u *user) Greeting(c context.Context, name string) (resp string, err error) {
	resp = fmt.Sprintf("Hello, %s!", name)
	return
}
