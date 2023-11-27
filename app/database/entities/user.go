package entities

type User struct {
	UserId            uint   `json:"id_usuario" gorm:"primaryKey;autoIncrement;column:id_usuario"`
	FullName          string `json:"nome_completo" gorm:"column:nome_completo"`
	UserName          string `json:"apelido" gorm:"column:apelido"`
	Password          string `json:"senha" gorm:"column:senha"`
	ProfilePictureUri string `json:"-" gorm:"column:uri_foto_perfil"`
	LastLogin         string `json:"ultimo_login" gorm:"column:ultimo_login"`
}

func (User) TableName() string {
	return "usuarios"
}
