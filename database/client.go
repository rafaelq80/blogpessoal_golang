package database

import (
	"blogpessoal/entities"
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
	Instance.AutoMigrate(&entities.Postagem{})
	Instance.AutoMigrate(&entities.Tema{})
	Instance.AutoMigrate(&entities.Usuario{})
	log.Println("Criação das Tabelas Finalizada...")
}
