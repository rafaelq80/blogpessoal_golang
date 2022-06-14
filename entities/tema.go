package entities

type Tema struct {
	ID        uint       `gorm:"primary_key, AUTO_INCREMENT" json:"id,omitempty"`
	Descricao string     `gorm:"not null" json:"descricao,omitempty" validate:"required"`
	Postagens []Postagem `gorm:"foreignkey:TemaID;references:ID;constraint:OnDelete:CASCADE;" json:"postagens,omitempty"`
}

func (Tema) TableName() string {
	return "tb_temas"
}