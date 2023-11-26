package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FullName          string `json:"nome_completo" gorm:"nome_completo"`
	UserName          string `json:"apelido" gorm:"apelido"`
	Password          string `json:"senha" gorm:"senha"`
	ProfilePictureUri string `json:"-" gorm:"uri_foto_perfil"`
	LastLogin         string `json:"ultimo_login" gorm:"ultimo_login"`
}

func (User) TableName() string {
	return "usuarios"
}
