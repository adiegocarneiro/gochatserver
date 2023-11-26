package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB, app *fiber.App) *fiber.App {
	userRouter(db, app)
	wsRouter(app)
	return app
}
