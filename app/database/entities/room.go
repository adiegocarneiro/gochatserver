package entities

import (
	"time"

	"gorm.io/gorm"
)

type Room struct {
	RoomId      uint           `json:"id_sala" gorm:"primaryKey;autoIncrement;column:id_sala"`
	Name        string         `json:"nome_da_sala" gorm:"column:nome_da_sala"`
	Users       []*User        `json:"participantes,omitempty" gorm:"many2many:usuarios_salas"`
	ChatMessage []*ChatMessage `json:"mensagens,omitempty" gorm:"foreignKey:RoomId"`
	CreatedAt   time.Time      `json:"data_criacao"`
	Deleted     gorm.DeletedAt `json:"-"`
}

func (Room) TableName() string {
	return "salas"
}
