package controller

import (
	"learn_fiber/service"

	"github.com/gofiber/fiber/v2"
)

type (
	// User 핸들러
	User interface {
		Greeting(c *fiber.Ctx) error
	}
	user struct {
		svc service.User
	}
)

// NewUserController 핸들러 생성자
func NewUserController(u service.User) User {
	return &user{u}
}

func (uc *user) Greeting(c *fiber.Ctx) error {
	name := c.Query("name")
	resp, err := uc.svc.Greeting(c.Context(), name)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.SendString(resp)
}
