package handlers

import (
	"gochatserver/app/database/repositories"

	"github.com/gofiber/fiber"
)

func GetAllUsersHandler(repo *repositories.Repository) fiber.Handler {
	return func(c *fiber.Ctx) {
		users := repo.GetAllUsers(c)
		if users == nil {
			message := map[string]string{"message": "Not Found"}
			c.Status(404).JSON(message)
		}
		c.Status(200).JSON(users)
	}
}

func GetUserHandler(repo *repositories.Repository) fiber.Handler {
	return func(c *fiber.Ctx) {
		id := c.Params("id")

		if id != "" {
			message := map[string]string{"message": "Parâmetros incorretos."}
			c.Status(400).JSON(message)
		}

		user := repo.GetUserById(c, id)
		if user == nil {
			message := map[string]string{"message": "Not Found"}
			c.Status(404).JSON(message)
		}
		c.Status(200).JSON(user)
	}
}
