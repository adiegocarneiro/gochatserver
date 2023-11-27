package repositories

import (
	"gochatserver/app/config"
	"gochatserver/app/database/entities"

	"github.com/gofiber/fiber/v2"
)

func (repo Repository) CreateUser(ctx *fiber.Ctx, userData *entities.User) config.Response {
	if results := repo.DB.Create(&userData); results.Error != nil {
		return config.Response{
			Success: false,
			Message: "Erro ao cadastrar usuário",
		}
	}
	return config.Response{
		Success: true,
		Message: "Usuário cadastrado com sucesso!",
		Object:  userData,
	}
}

func (repo Repository) GetAllUsers(ctx *fiber.Ctx) *[]entities.User {
	users := []entities.User{}
	if results := repo.DB.Find(&users); results.Error != nil {
		return nil
	}
	return &users
}

func (repo Repository) GetUserById(ctx *fiber.Ctx, id string) *entities.User {
	user := entities.User{}
	if results := repo.DB.First(&user, id); results.Error != nil {
		return nil
	}
	return &user
}
