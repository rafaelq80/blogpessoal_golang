package controllers

import (
	"blogpessoal/database"
	"blogpessoal/entities"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
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

	temaId := mux.Vars(r)["id"]

	if !checkIfUsuarioExists(temaId) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Usuario Não Encontrada!")
		return
	}

	var tema entities.Usuario

	database.Instance.Preload("Postagens").First(&tema, temaId)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tema)
}

// postUsuario godoc
// @Summary Criar Usuario
// @Description Cria um novo Usuario
// @Tags usuarios
// @Accept  json
// @Produce  json
// @Param tema body entities.Usuario true "Criar Usuario"
// @Success 201 {object} entities.Usuario
// @Router /usuarios [post]
func CreateUsuario(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var tema entities.Usuario
	json.NewDecoder(r.Body).Decode(&tema)

	validate := validator.New()

	err := validate.Struct(tema)
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

	database.Instance.Create(&tema)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(tema)
}

// putUsuario godoc
// @Summary Atualizar Usuario
// @Description Edita um Usuario
// @Tags usuarios
// @Accept  json
// @Produce  json
// @Param id path string true "Id do tema"
// @Param Usuario body entities.Usuario true "Atualizar Usuario"
// @Success 200 {object} entities.Usuario
// @Success 400 {object} errorResponse
// @Success 404 {object} errorResponse
// @Success 405 {object} errorResponse
// @Router /usuarios/{id} [put]
func UpdateUsuario(w http.ResponseWriter, r *http.Request) {

	temaId := mux.Vars(r)["id"]

	if !checkIfUsuarioExists(temaId) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Usuario Não Encontrado!")
		return
	}

	var tema entities.Usuario

	database.Instance.First(&tema, temaId)
	json.NewDecoder(r.Body).Decode(&tema)
	
	validate := validator.New()

	err := validate.Struct(tema)
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

	database.Instance.Save(&tema)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tema)
}

func checkIfUsuarioExists(temaId string) bool {

	var tema entities.Usuario
	database.Instance.First(&tema, temaId)

	return tema.ID != 0 

}