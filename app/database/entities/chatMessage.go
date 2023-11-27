package entities

import (
	"time"

	"gorm.io/gorm"
)

type ChatMessage struct {
	ChatMessageId uint           `json:"id_mensagem_chat" gorm:"primaryKey;autoIncrement;column:id_mensagem_chat"`
	Message       string         `json:"mensagem" gorm:"column:mensagem"`
	UserId        *uint          `json:"id_usuario" gorm:"column:id_usuario"`
	RoomId        *uint          `json:"id_sala" gorm:"column:id_sala"`
	CreatedAt     time.Time      `json:"data_hora"`
	Deleted       gorm.DeletedAt `json:"-"`
}

func (ChatMessage) TableName() string {
	return "mensagens"
}
