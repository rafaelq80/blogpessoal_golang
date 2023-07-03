package main

import (
	"blogpessoal/auth"
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
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
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
	RegisterUsuarioRoutes(router)
	RegisterSwaggerRoutes(router)
	//handler := cors.Default().Handler(router)

	// Start the server
	log.Printf("Iniciando o Servidor na porta %s", AppConfig.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("127.0.0.1:%v", AppConfig.Port), router))
}

func RegisterPostagemRoutes(router *mux.Router) {
	router.HandleFunc("/postagens", auth.SetMiddlewareJSON(auth.SetMiddlewareAuthentication(controllers.GetPostagens))).Methods("GET")
	router.HandleFunc("/postagens/{id}", auth.SetMiddlewareJSON(auth.SetMiddlewareAuthentication(controllers.GetPostagemById))).Methods("GET")
	router.HandleFunc("/postagens/titulo/{titulo}", auth.SetMiddlewareJSON(auth.SetMiddlewareAuthentication(controllers.GetPostagemByTitulo))).Methods("GET")
	router.HandleFunc("/postagens", auth.SetMiddlewareJSON(auth.SetMiddlewareAuthentication(controllers.CreatePostagem))).Methods("POST")
	router.HandleFunc("/postagens", auth.SetMiddlewareJSON(auth.SetMiddlewareAuthentication(controllers.UpdatePostagem))).Methods("PUT")
	router.HandleFunc("/postagens/{id}", auth.SetMiddlewareJSON(auth.SetMiddlewareAuthentication(controllers.DeletePostagem))).Methods("DELETE")
}

func RegisterTemaRoutes(router *mux.Router) {
	router.HandleFunc("/temas", auth.SetMiddlewareJSON(auth.SetMiddlewareAuthentication(controllers.GetTemas))).Methods("GET")
	router.HandleFunc("/temas", auth.SetMiddlewareJSON(auth.SetMiddlewareAuthentication(controllers.CreateTema))).Methods("POST")
	router.HandleFunc("/temas/{id}", auth.SetMiddlewareJSON(auth.SetMiddlewareAuthentication(controllers.GetTemaById))).Methods("GET")
	router.HandleFunc("/temas/descricao/{descricao}", auth.SetMiddlewareJSON(auth.SetMiddlewareAuthentication(controllers.GetTemaByDescricao))).Methods("GET")
	router.HandleFunc("/temas", auth.SetMiddlewareJSON(auth.SetMiddlewareAuthentication(controllers.UpdateTema))).Methods("PUT")
	router.HandleFunc("/temas/{id}", auth.SetMiddlewareJSON(auth.SetMiddlewareAuthentication(controllers.DeleteTema))).Methods("DELETE")
}

func RegisterUsuarioRoutes(router *mux.Router) {
	router.HandleFunc("/usuarios/all", auth.SetMiddlewareJSON(auth.SetMiddlewareAuthentication(controllers.GetUsuarios))).Methods("GET")
	router.HandleFunc("/usuarios/cadastrar", auth.SetMiddlewareJSON(controllers.CreateUsuario)).Methods("POST")
	router.HandleFunc("/usuarios/{id}", auth.SetMiddlewareJSON(auth.SetMiddlewareAuthentication(controllers.GetUsuarioById))).Methods("GET")
	router.HandleFunc("/usuarios/atualizar", auth.SetMiddlewareJSON(controllers.UpdateUsuario)).Methods("PUT")
	router.HandleFunc("/usuarios/logar", auth.SetMiddlewareJSON(controllers.Authetication)).Methods("POST")
}

func RegisterSwaggerRoutes(router *mux.Router) {
	router.PathPrefix("/").Handler(httpSwagger.WrapHandler)

}
