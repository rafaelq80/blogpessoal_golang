package controllers

import (
	"blogpessoal/auth"
	"blogpessoal/database"
	"blogpessoal/model"
	"encoding/json"
	"net/http"
)

// postUsuario godoc
// @Summary Autenticar Usuario
// @Description Autentica um Usuario
// @Tags usuarios
// @Accept  json
// @Produce  json
// @Param usuario body model.UsuarioLogin true "Autenticar Usuario"
// @Success 200 {object} model.UsuarioLogin
// @Router /usuarios/logar [post]
func Authetication(w http.ResponseWriter, r *http.Request) {
	
	var usuario model.Usuario
	var err error
	var token string
	
	w.Header().Set("Content-Type", "application/json")
	var usuarioLogin model.UsuarioLogin
	json.NewDecoder(r.Body).Decode(&usuarioLogin)

	if !checkIfUsuarioEmailExists(usuarioLogin.Usuario){
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Usuario Inválido!")
		return
	}
	
	database.Instance.Where("usuario = ?", usuarioLogin.Usuario).Find(&usuario) 

	if !CheckPasswordHash(usuarioLogin.Senha, usuario.Senha){
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Usuario Inválido!")
		return
	}

	token, err = auth.CreateToken(usuarioLogin.Usuario)

	if err != nil{
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode("Usuario Inválido!")
		return
	}

	usuarioLogin.ID = usuario.ID
	usuarioLogin.Nome = usuario.Nome
	usuarioLogin.Foto = usuario.Foto
	usuarioLogin.Senha = ""
	usuarioLogin.Token = "Bearer " + token

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(usuarioLogin)
}