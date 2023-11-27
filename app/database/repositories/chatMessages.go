package repositories

import (
	"gochatserver/app/config"
	"gochatserver/app/database/entities"

	"github.com/gofiber/fiber/v2"
)

func (repo Repository) CreateMessage(ctx *fiber.Ctx, messageData *entities.ChatMessage) config.Response {
	room := entities.Room{}
	if results := repo.DB.Model(&entities.Room{}).Preload("Users").First(&room, messageData.RoomId); results.Error != nil {
		return config.Response{
			Success: false,
			Message: "Erro ao enviar mensagem: Sala não encontrada!",
		}
	}

	noSuchUser := true //

	for _, user := range room.Users {
		if user.UserId == *messageData.UserId {
			noSuchUser = false
			break
		}
	}

	if noSuchUser {
		return config.Response{
			Success: false,
			Message: "Erro ao enviar mensagem: Usuário não está nesta sala.",
		}
	}

	if results := repo.DB.Create(&messageData); results.Error != nil {
		return config.Response{
			Success: false,
			Message: "Erro ao enviar mensagem",
		}
	}
	return config.Response{
		Success: true,
	}
}

func (repo Repository) GetMessages(ctx *fiber.Ctx) *[]entities.ChatMessage {
	chatMessages := []entities.ChatMessage{}
	if results := repo.DB.Preload("Users").Find(&chatMessages); results.Error != nil {
		return nil
	}
	return &chatMessages
}
