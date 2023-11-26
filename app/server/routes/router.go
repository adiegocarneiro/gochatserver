package routes

import (
	"github.com/gofiber/fiber"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB, app *fiber.App) *fiber.App {
	userRouter(db, app)

	return app
}
