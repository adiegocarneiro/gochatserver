package routes

import (
	"gochatserver/app/database/repositories"
	"gochatserver/app/server/handlers"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func userRouter(db *gorm.DB, app *fiber.App) *fiber.App {
	repo := &repositories.Repository{
		DB: db,
	}

	app.Get("users/", handlers.GetAllUsersHandler(repo))
	app.Get("users/:id", handlers.GetUserHandler(repo))
	app.Post("users/", handlers.CreateUserHandler(repo))
	return app
}
