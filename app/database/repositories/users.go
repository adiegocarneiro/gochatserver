package repositories

import (
	"gochatserver/app/database/entities"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
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
