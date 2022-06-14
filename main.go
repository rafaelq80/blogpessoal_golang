package main

import (
	"blogpessoal/controllers"
	"blogpessoal/database"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"

	_ "blogpessoal/docs"

	httpSwagger "github.com/swaggo/http-swagger"

)

var DB *gorm.DB

// @title Blog Pessoal
// @version 1.0
// @description Projeto Blog Pessoal
// @contact.name Rafael Queiróz
// @contact.email rafaelproinfo@gmail.com
// @contact.url https://github.com/rafaelq80
// @license.name Apache 2.0
// @license.url https://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func main() {

	// Load Configurations from config.json using Viper
	LoadAppConfig()

	// Initialize Database
	database.Connect(AppConfig.ConnectionString)
	database.Migrate()

	// Initialize the router
	router := mux.NewRouter().StrictSlash(true)

	// Register Routes
	RegisterPostagemRoutes(router)
	RegisterTemaRoutes(router)
	RegisterSwaggerRoutes(router)
	//handler := cors.Default().Handler(router)

	// Start the server
	log.Printf("Iniciando o Servidor na porta %s", AppConfig.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("127.0.0.1:%v", AppConfig.Port), router))
}

func RegisterPostagemRoutes(router *mux.Router) {
	router.HandleFunc("/postagens", controllers.GetPostagens).Methods("GET")
	router.HandleFunc("/postagens/{id}", controllers.GetPostagemById).Methods("GET")
	router.HandleFunc("/postagens/titulo/{titulo}", controllers.GetPostagemByTitulo).Methods("GET")
	router.HandleFunc("/postagens", controllers.CreatePostagem).Methods("POST")
	router.HandleFunc("/postagens/{id}", controllers.UpdatePostagem).Methods("PUT")
	router.HandleFunc("/postagens/{id}", controllers.DeletePostagem).Methods("DELETE")
}

func RegisterTemaRoutes(router *mux.Router) {
	router.HandleFunc("/temas", controllers.GetTemas).Methods("GET")
	router.HandleFunc("/temas", controllers.CreateTema).Methods("POST")
	router.HandleFunc("/temas/{id}", controllers.GetTemaById).Methods("GET")
	router.HandleFunc("/temas/descricao/{descricao}", controllers.GetTemaByDescricao).Methods("GET")
	router.HandleFunc("/temas/{id}", controllers.UpdateTema).Methods("PUT")
	router.HandleFunc("/temas/{id}", controllers.DeleteTema).Methods("DELETE")
}

func RegisterSwaggerRoutes(router *mux.Router) {
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
}

