package handlers

import (
	"gochatserver/app/config"
	"gochatserver/app/database/entities"
	"gochatserver/app/database/repositories"

	"github.com/gofiber/fiber/v2"
)

func CreateRoomHandler(repo *repositories.Repository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		response := config.Response{}
		roomData := entities.Room{}
		if err := c.BodyParser(&roomData); err != nil {
			response.Success = false
			response.Message = "Parâmetros incorretos!"
			return c.Status(400).JSON(response)
		}

		createdUser := repo.CreateRoom(c, &roomData)
		return c.Status(200).JSON(createdUser)
	}
}

func GetAllRoomsHandler(repo *repositories.Repository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		users := repo.GetAllRooms(c)
		if users == nil {
			message := map[string]string{"message": "Not Found"}
			return c.Status(404).JSON(message)
		}
		return c.Status(200).JSON(users)
	}
}

func GetRoomHandler(repo *repositories.Repository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")

		if id == "" {
			message := map[string]string{"message": "Parâmetros incorretos."}
			return c.Status(400).JSON(message)
		}

		user := repo.GetRoomById(c, id)
		if user == nil {
			message := map[string]string{"message": "Sala não encontrada."}
			return c.Status(404).JSON(message)
		}
		return c.Status(200).JSON(user)
	}
}

func IngressChatRoom(repo *repositories.Repository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ingressRequest := config.IngressRequest{}

		if err := c.BodyParser(&ingressRequest); err != nil {
			response := config.Response{}
			response.Success = false
			response.Message = "Parâmetros incorretos!"
			return c.Status(400).JSON(response)
		}

		response := repo.RoomIngress(c, ingressRequest.UserId, ingressRequest.RoomId)
		return c.Status(200).JSON(response)
	}
}

func ExitChatRoom(repo *repositories.Repository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ingressRequest := config.IngressRequest{}

		if err := c.BodyParser(&ingressRequest); err != nil {
			response := config.Response{}
			response.Success = false
			response.Message = "Parâmetros incorretos!"
			return c.Status(400).JSON(response)
		}

		response := repo.RoomIngress(c, ingressRequest.UserId, ingressRequest.RoomId)
		return c.Status(200).JSON(response)
	}
}
