package entities

import (
	"time"
)

type Postagem struct {
	ID        uint      `gorm:"primary_key, AUTO_INCREMENT" json:"id" example:"1"`
	Titulo    string    `gorm:"not null;size:100" json:"titulo" validate:"required,min=5,max=100" example:"Minha primeira postagem"`
	Texto     string    `gorm:"not null;size:1000" json:"texto" validate:"required,min=10,max=1000" example:"Texto da primeira postagem"`
	UpdatedAt time.Time `gorm:"column:data;autoUpdateTime:mili" json:"data" example:"2022-04-09T21:21:46+00:00"`
	TemaID    uint      `gorm:"column:tema_id;not null" json:"tema_id" validate:"required" example:"1"`
	Tema      Tema      `gorm:"ForeignKey:TemaID;association_foreignkey:ID" json:"tema" validate:"-"`
}

func (Postagem) TableName() string {
	return "tb_postagens"
}
