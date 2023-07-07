package main

import (
	"gorm-fiber/database"
	"gorm-fiber/database/migrations"
	"gorm-fiber/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// CONNECT TO DATABASE
	database.DatabaseInit()

	// MIGRATION
	migrations.Migration()

	app.Get("/user", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Hello from user route",
		})
	})

	//ROUTNG
	routes.RouteInit(app)

	app.Listen(":8080")
}
