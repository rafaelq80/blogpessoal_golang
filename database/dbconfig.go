package database

import (
	"blogpessoal/model"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Instance *gorm.DB
var err error

func Connect(connectionString string) {
	Instance, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		panic("Não foi possível conectar ao banco de dados!")
	}

	log.Println("Conectando ao Banco de Dados...")
}

func Migrate() {
	Instance.AutoMigrate(&model.Postagem{})
	Instance.AutoMigrate(&model.Tema{})
	Instance.AutoMigrate(&model.Usuario{})
	log.Println("Criação das Tabelas Finalizada...")
}
