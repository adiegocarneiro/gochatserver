package routes

import (
	"gochatserver/app/database/repositories"
	"gochatserver/app/server/handlers"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func roomRouter(db *gorm.DB, app *fiber.App) *fiber.App {
	repo := &repositories.Repository{
		DB: db,
	}

	app.Get("rooms/", handlers.GetAllRoomsHandler(repo))
	app.Get("rooms/:id", handlers.GetRoomHandler(repo))
	app.Post("rooms/", handlers.CreateRoomHandler(repo))
	app.Post("rooms/ingress", handlers.IngressChatRoom(repo))
	return app
}
