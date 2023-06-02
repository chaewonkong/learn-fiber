package service_test

import (
	"context"
	"learn_fiber/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGreeting(t *testing.T) {
	t.Run("if name given, returns pre-defined string containing the name", func(t *testing.T) {
		s := service.NewUserService()
		// given
		name := "Leon"

		// when
		res, err := s.Greeting(context.Background(), name)

		// then
		assert.NoError(t, err)
		assert.Contains(t, res, name)
	})
}
