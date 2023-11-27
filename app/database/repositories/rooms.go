package repositories

import (
	"gochatserver/app/config"
	"gochatserver/app/database/entities"

	"github.com/gofiber/fiber/v2"
)

func (repo Repository) CreateRoom(ctx *fiber.Ctx, roomData *entities.Room) config.Response {
	if results := repo.DB.Create(&roomData); results.Error != nil {
		return config.Response{
			Success: false,
			Message: "Erro ao cadastrar sala",
		}
	}
	return config.Response{
		Success: true,
		Message: "Sala cadastrada com sucesso!",
		Object:  roomData,
	}
}

func (repo Repository) GetAllRooms(ctx *fiber.Ctx) *config.Response {
	rooms := []entities.Room{}
	if results := repo.DB.Preload("Users").Preload("ChatMessage").Find(&rooms); results.Error != nil {
		return &config.Response{
			Success: false,
			Message: "Erro ao obter lista de salas",
		}
	}
	return &config.Response{
		Success: true,
		Message: "Salas",
		Object:  rooms,
	}
}

func (repo Repository) GetRoomById(ctx *fiber.Ctx, id string) *config.Response {
	room := entities.Room{}
	if results := repo.DB.Preload("Users").Preload("ChatMessage").Preload("Users").First(&room, id); results.Error != nil {
		return &config.Response{
			Success: false,
			Message: "Erro ao obter sala",
		}
	}
	return &config.Response{
		Success: true,
		Message: "Sala",
		Object:  room,
	}
}

func (repo Repository) RoomIngress(ctx *fiber.Ctx, userId string, roomId string) *config.Response {
	user := entities.User{}
	room := entities.Room{}

	if results := repo.DB.First(&user, userId); results.Error != nil {
		return &config.Response{
			Success: false,
			Message: "Usuário não encontrado!",
		}
	}

	if results := repo.DB.First(&room, roomId); results.Error != nil {
		return &config.Response{
			Success: false,
			Message: "Sala não encontrada!",
		}
	}
	room.Users = append(room.Users, &user)

	if results := repo.DB.Save(&room); results.Error != nil {
		return &config.Response{
			Success: false,
			Message: "Erro ao atualizar sala!",
		}
	}

	return &config.Response{
		Success: true,
		Message: "Usuário " + user.UserName + " entrou na sala!",
	}
}

func (repo Repository) RoomEgress(ctx *fiber.Ctx, userId string, roomId string) *config.Response {
	user := entities.User{}
	room := entities.Room{}

	if results := repo.DB.First(&user, userId); results.Error != nil {
		return &config.Response{
			Success: false,
			Message: "Usuário não encontrado!",
		}
	}

	if results := repo.DB.First(&room, roomId); results.Error != nil {
		return &config.Response{
			Success: false,
			Message: "Sala não encontrada!",
		}
	}
	room.Users = append(room.Users, &user)

	if results := repo.DB.Save(&room); results.Error != nil {
		return &config.Response{
			Success: false,
			Message: "Erro ao atualizar sala!",
		}
	}

	return &config.Response{
		Success: true,
		Message: "Usuário " + user.UserName + " entrou na sala!",
	}
}
