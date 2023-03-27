package controllers

import (
	"blogpessoal/database"
	"blogpessoal/entities"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

// getAll godoc
// @Summary Listar Usuarios
// @Description Lista todos os Usuarios
// @Tags usuarios
// @Accept  json
// @Produce  json
// @Success 200 {array} entities.Usuario
// @Router /usuarios [get]
func GetUsuarios(w http.ResponseWriter, r *http.Request) {

	var usuarios []entities.Usuario

	database.Instance.Preload("Postagens").Find(&usuarios)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(usuarios)
}

// getById godoc
// @Summary Listar Usuario por id
// @Description Lista um Usuario por id
// @Tags usuarios
// @Accept  json
// @Produce  json
// @Param id path string true "Id do Usuario"
// @Success 200 {array} entities.Usuario
// @Success 400 {object} errorResponse
// @Success 404 {object} errorResponse
// @Success 405 {object} errorResponse
// @Router /usuarios/{id} [get]
func GetUsuarioById(w http.ResponseWriter, r *http.Request) {

	usuarioId := mux.Vars(r)["id"]

	if !checkIfUsuarioExists(usuarioId) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Usuario Não Encontrada!")
		return
	}

	var usuario entities.Usuario

	database.Instance.Preload("Postagens").First(&usuario, usuarioId)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(usuario)
}

// postUsuario godoc
// @Summary Criar Usuario
// @Description Cria um novo Usuario
// @Tags usuarios
// @Accept  json
// @Produce  json
// @Param usuario body entities.Usuario true "Criar Usuario"
// @Success 201 {object} entities.Usuario
// @Router /usuarios [post]
func CreateUsuario(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var usuario entities.Usuario
	json.NewDecoder(r.Body).Decode(&usuario)

	validate := validator.New()

	err := validate.Struct(usuario)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		responseBody := map[string]string{"error": validationErrors.Error()}
		if err := json.NewEncoder(w).Encode(responseBody); err != nil {
			log.Fatalf("Erro: %s", err)
		}
		return
	}

	hash,_ := HashPassword(usuario.Senha)
	usuario.Senha = hash

	database.Instance.Create(&usuario)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(usuario)
}

// putUsuario godoc
// @Summary Atualizar Usuario
// @Description Edita um Usuario
// @Tags usuarios
// @Accept  json
// @Produce  json
// @Param id path string true "Id do usuario"
// @Param Usuario body entities.Usuario true "Atualizar Usuario"
// @Success 200 {object} entities.Usuario
// @Success 400 {object} errorResponse
// @Success 404 {object} errorResponse
// @Success 405 {object} errorResponse
// @Router /usuarios/{id} [put]
func UpdateUsuario(w http.ResponseWriter, r *http.Request) {

	usuarioId := mux.Vars(r)["id"]

	if !checkIfUsuarioExists(usuarioId) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Usuario Não Encontrado!")
		return
	}

	var usuario entities.Usuario

	database.Instance.First(&usuario, usuarioId)
	json.NewDecoder(r.Body).Decode(&usuario)
	
	validate := validator.New()

	err := validate.Struct(usuario)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		responseBody := map[string]string{"error": validationErrors.Error()}
		if err := json.NewEncoder(w).Encode(responseBody); err != nil {
			log.Fatalf("Erro: %s", err)
		}
		return
	}

	hash,_ := HashPassword(usuario.Senha)
	usuario.Senha = hash
	
	database.Instance.Save(&usuario)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(usuario)
}

func checkIfUsuarioExists(usuarioId string) bool {

	var usuario entities.Usuario
	database.Instance.First(&usuario, usuarioId)

	return usuario.ID != 0 

}

func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}