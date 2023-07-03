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

/* Descomentar as próximas 3 linhas
   apenas ao atualizar o Swagger
*/
// type errorResponse struct {
// 	Message    string
// }

// getAll godoc
// @Summary Listar Postagens
// @Description Lista todas as Postagens
// @Tags postagens
// @Accept  json
// @Produce  json
// @Success 200 {array} model.Postagem
// @Router /postagens [get]
// @Security Bearer
func GetPostagens(w http.ResponseWriter, _ *http.Request) {
	var postagens []model.Postagem

	database.Instance.Joins("Tema").Joins("Usuario").Find(&postagens)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(postagens)
}

// getById godoc
// @Summary Listar Postagem por id
// @Description Lista uma Postagem por id
// @Tags postagens
// @Accept  json
// @Produce  json
// @Param id path string true "Id da Postagem"
// @Success 200 {array} model.Postagem
// @Success 400 {object} errorResponse
// @Success 404 {object} errorResponse
// @Success 405 {object} errorResponse
// @Router /postagens/{id} [get]
// @Security Bearer
func GetPostagemById(w http.ResponseWriter, r *http.Request) {

	postagemId := mux.Vars(r)["id"]

	if !checkIfPostagemExists(postagemId) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Postagem Não Encontrada!")
		return
	}

	var postagem model.Postagem

	database.Instance.Joins("Tema").Joins("Usuario").First(&postagem, postagemId)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(postagem)
}

// getByTitulo godoc
// @Summary Listar Postagens por título
// @Description Lista todas as Postagem por título
// @Tags postagens
// @Accept  json
// @Produce  json
// @Param titulo path string true "Título da Postagem"
// @Success 200 {array} model.Postagem
// @Success 400 {object} errorResponse
// @Success 405 {object} errorResponse
// @Router /postagens/titulo/{titulo} [get]
// @Security Bearer
func GetPostagemByTitulo(w http.ResponseWriter, r *http.Request) {

	postagemTitulo := mux.Vars(r)["titulo"]

	var postagens []model.Postagem

	database.Instance.Joins("Tema").Joins("Usuario").Where("titulo LIKE ?", "%"+postagemTitulo+"%").Find(&postagens)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(postagens)
}

// postPostagem godoc
// @Summary Criar Postagem
// @Description Cria uma nova Postagem
// @Tags postagens
// @Accept  json
// @Produce  json
// @Param postagem body model.Postagem true "Criar Postagem"
// @Success 201 {object} model.Postagem
// @Router /postagens [post]
// @Security Bearer
func CreatePostagem(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var postagem model.Postagem
	json.NewDecoder(r.Body).Decode(&postagem)

	validate := validator.New()

	err := validate.Struct(postagem)

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

	var temaId string = strconv.FormatUint(uint64(postagem.TemaID), 10)

	if !checkIfTemaExists(temaId) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Tema Não Encontrado!")
		return
	}

	database.Instance.Create(&postagem)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(postagem)
}

// putPostagem godoc
// @Summary Atualizar Postagem
// @Description Edita uma Postagem
// @Tags postagens
// @Accept  json
// @Produce  json
// @Param id path string true "Id da Postagem"
// @Param postagem body model.Postagem true "Atualizar Postagem"
// @Success 200 {object} model.Postagem
// @Success 400 {object} errorResponse
// @Success 404 {object} errorResponse
// @Success 405 {object} errorResponse
// @Router /postagens [put]
// @Security Bearer
func UpdatePostagem(w http.ResponseWriter, r *http.Request) {

	var postagem model.Postagem
	json.NewDecoder(r.Body).Decode(&postagem)

	validate := validator.New()

	err := validate.Struct(postagem)

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

	var id = strconv.FormatUint(uint64(postagem.ID), 10)

	if !checkIfPostagemExists(id) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Postagem Não Encontrada!")
		return
	}

	var temaId string = strconv.FormatUint(uint64(postagem.TemaID), 10)

	if !checkIfTemaExists(temaId) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Tema Não Encontrado!")
		return
	}

	database.Instance.Save(&postagem)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(postagem)
}

// deletePostagem godoc
// @Summary Deletar Postagem
// @Description Apaga uma Postagem
// @Tags postagens
// @Accept  json
// @Produce  json
// @Param id path string true "Id da Postagem"
// @Success 204 {object} errorResponse
// @Success 400 {object} errorResponse
// @Success 404 {object} errorResponse
// @Success 405 {object} errorResponse
// @Router /postagens/{id} [delete]
// @Security Bearer
func DeletePostagem(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	postagemId := mux.Vars(r)["id"]

	if !checkIfPostagemExists(postagemId) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Postagem Não Encontrada!")
		return
	}

	var postagem model.Postagem

	database.Instance.Delete(&postagem, postagemId)
	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode("Postagem Deletada!")
}

func checkIfPostagemExists(postagemId string) bool {

	var postagem model.Postagem

	database.Instance.First(&postagem, postagemId)

	return postagem.ID != 0
}
