package controllers

import (
	"blogpessoal/database"
	"blogpessoal/model"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

// getAll godoc
// @Summary Listar Temas
// @Description Lista todos os Temas
// @Tags temas
// @Accept  json
// @Produce  json
// @Success 200 {array} model.Tema
// @Router /temas [get]
// @Security Bearer
func GetTemas(w http.ResponseWriter, _ *http.Request) {

	var temas []model.Tema

	database.Instance.Preload("Postagens").Find(&temas)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(temas)
}

// getById godoc
// @Summary Listar Tema por id
// @Description Lista um Tema por id
// @Tags temas
// @Accept  json
// @Produce  json
// @Param id path string true "Id do Tema"
// @Success 200 {array} model.Tema
// @Success 400 {object} errorResponse
// @Success 404 {object} errorResponse
// @Success 405 {object} errorResponse
// @Router /temas/{id} [get]
// @Security Bearer
func GetTemaById(w http.ResponseWriter, r *http.Request) {

	temaId := mux.Vars(r)["id"]

	if !checkIfTemaExists(temaId) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Tema Não Encontrada!")
		return
	}

	var tema model.Tema

	database.Instance.Preload("Postagens").First(&tema, temaId)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tema)
}

// getByDescricao godoc
// @Summary Listar Temas por descrição
// @Description Lista todos os Temas por descrição
// @Tags temas
// @Accept  json
// @Produce  json
// @Param descricao path string true "Descrição do Tema"
// @Success 200 {array} model.Tema
// @Success 400 {object} errorResponse
// @Success 405 {object} errorResponse
// @Router /temas/descricao/{descricao} [get]
// @Security Bearer
func GetTemaByDescricao(w http.ResponseWriter, r *http.Request) {

	temaDescricao := mux.Vars(r)["descricao"]

	var temas []model.Tema

	database.Instance.Preload("Postagens").Where("descricao LIKE ?", "%"+temaDescricao+"%").Find(&temas)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(temas)
}

// postTema godoc
// @Summary Criar Tema
// @Description Cria um novo Tema
// @Tags temas
// @Accept  json
// @Produce  json
// @Param tema body model.Tema true "Criar Tema"
// @Success 201 {object} model.Tema
// @Router /temas [post]
// @Security Bearer
func CreateTema(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var tema model.Tema
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

// putTema godoc
// @Summary Atualizar Tema
// @Description Edita um Tema
// @Tags temas
// @Accept  json
// @Produce  json
// @Param id path string true "Id do tema"
// @Param Tema body model.Tema true "Atualizar Tema"
// @Success 200 {object} model.Tema
// @Success 400 {object} errorResponse
// @Success 404 {object} errorResponse
// @Success 405 {object} errorResponse
// @Router /temas [put]
// @Security Bearer
func UpdateTema(w http.ResponseWriter, r *http.Request) {

	var tema model.Tema
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

	var id = strconv.FormatUint(uint64(tema.ID), 10)

	if !checkIfTemaExists(id) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Tema Não Encontrado!")
		return
	}

	database.Instance.Save(&tema)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tema)
}

// deleteTema godoc
// @Summary Deletar Tema
// @Description Apaga uma Tema
// @Tags temas
// @Accept  json
// @Produce  json
// @Param id path string true "Id do Tema"
// @Success 204 {object} errorResponse
// @Success 400 {object} errorResponse
// @Success 404 {object} errorResponse
// @Success 405 {object} errorResponse
// @Router /temas/{id} [delete]
// @Security Bearer
func DeleteTema(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	temaId := mux.Vars(r)["id"]

	if !checkIfTemaExists(temaId) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Tema Não Encontrado!")
		return
	}

	var tema model.Tema

	database.Instance.Delete(&tema, temaId)
	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode("Tema Deletado!")
}

func checkIfTemaExists(temaId string) bool {

	var tema model.Tema
	database.Instance.First(&tema, temaId)

	return tema.ID != 0

}
