package main

import (
	"learn_fiber/controller"
	"learn_fiber/database"
	"learn_fiber/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func main() {
	app := fiber.New()
	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		Format:   "${pid} ${locals:requestid} ${status} - ${method} ${path} \n",
		TimeZone: "Asia/Seoul",
	}))

	db := database.NewDatabase()
	db.Migrate()

	us := service.NewUserService()
	uc := controller.NewUserController(us)

	app.Get("/", uc.Greeting)
	app.Get("/metrics", monitor.New())

	app.Listen(":3000")
}
