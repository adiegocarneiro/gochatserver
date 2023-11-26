package handlers

import (
	"gochatserver/app/database/repositories"

	"github.com/gofiber/fiber/v2"
)

func GetAllUsersHandler(repo *repositories.Repository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		users := repo.GetAllUsers(c)
		if users == nil {
			message := map[string]string{"message": "Not Found"}
			return c.Status(404).JSON(message)
		}
		return c.Status(200).JSON(users)
	}
}

func GetUserHandler(repo *repositories.Repository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")

		if id != "" {
			message := map[string]string{"message": "Parâmetros incorretos."}
			return c.Status(400).JSON(message)
		}

		user := repo.GetUserById(c, id)
		if user == nil {
			message := map[string]string{"message": "Not Found"}
			return c.Status(404).JSON(message)
		}
		return c.Status(200).JSON(user)
	}
}
