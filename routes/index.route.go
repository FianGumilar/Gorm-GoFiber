package routes

import (
	"gorm-fiber/controllers"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(app *fiber.App) {
	app.Get("/users", controllers.UserGetAll)
	app.Get("/users/:id", controllers.UserGetById)
	app.Post("/users", controllers.CreateUser)
	app.Put("/users/:id", controllers.UpdateUser)
	app.Delete("/users/:id", controllers.DeleteUser)

	app.Get("/lockers", controllers.LockerGetAll)
	app.Get("/lockers/:id", controllers.LockerGetById)
	app.Post("/lockers", controllers.CreateLocker)
	app.Put("/lockers/:id", controllers.UpdateLocker)
	app.Delete("/lockers/:id", controllers.DeleteLocker)

	app.Get("/posts", controllers.GetAllPost)
	app.Post("/posts", controllers.CreatePost)

	app.Get("/tags", controllers.GetAllTag)
	app.Post("/tags", controllers.CreateTag)
}
