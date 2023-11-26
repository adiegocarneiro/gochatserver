package entities

type User struct {
	UserId            uint    `json:"id_usuario" gorm:"primaryKey;autoIncrement;column:id_usuario"`
	FullName          string  `json:"nome_completo" gorm:"nome_completo"`
	UserName          string  `json:"apelido" gorm:"apelido"`
	Password          string  `json:"senha" gorm:"senha"`
	ChatRooms         []*Room `json:"salas_conectadas" gorm:"many2many:usuarios_salas"`
	ProfilePictureUri string  `json:"-" gorm:"uri_foto_perfil"`
	LastLogin         string  `json:"ultimo_login" gorm:"ultimo_login"`
}

func (User) TableName() string {
	return "usuarios"
}
