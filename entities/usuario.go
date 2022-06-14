package entities

type Usuario struct {
	ID        	uint       `gorm:"primary_key, AUTO_INCREMENT" json:"id,omitempty"`
	Nome 		string     `gorm:"not null" json:"nome,omitempty" validate:"required"`
	Usuario 	string     `gorm:"not null" json:"usuario,omitempty" validate:"required"`
	Senha 		string     `gorm:"not null, min=8" json:"senha,omitempty" validate:"required"`
	Foto 		string     `json:"foto,omitempty"`
	//Postagens 	[]Postagem `gorm:"foreignkey:UsuarioID;references:ID;constraint:OnDelete:CASCADE;" json:"postagens,omitempty"`
}

func (Usuario) TableName() string {
	return "tb_usuarios"
}