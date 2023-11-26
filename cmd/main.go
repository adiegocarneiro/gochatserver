package main

import (
	"gochatserver/app/config"
	"gochatserver/app/database"
	"gochatserver/app/server/routes"
	"log"

	"github.com/gofiber/fiber"
)

func main() {
	config, err := config.NewConfig()
	app := fiber.New()

	if err != nil {
		log.Panic("Impossível acessar configurações de ambiente.")
	}
	db, err := database.NewPostgreConnection(config)
	if err != nil {
		log.Panic("Impossível conectar-se ao banco de dados.")
	}
	database.ExecuteMigrations(db)
	routes.SetupRoutes(db, app)
	app.Listen(":3000")
}
