# Projeto Blog Pessoal (Em desenvolvimento)

<br />

<div align="center">
    <img src="https://i.imgur.com/YC6Av6e.png" title="source: imgur.com" /> 
</div>

<br /><br />

## Diagrama de Classes

```mermaid
classDiagram
class Tema {
  - id : Long
  - descricao : String
  - postagem : List ~Postagem~
  + getAll()
  + getById(Long id)
  + getByDescricao(String descricao)
  + postTema(Tema tema)
  + putTema(Tema tema)
  + deleteTema(Long id)
}
class Postagem {
  - id : Long
  - titulo : String
  - texto: String
  - data: LocalDateTime
  - tema : Tema
  - usuario : Usuario
  + getAll()
  + getById(Long id)
  + getByTitulo(String titulo)
  + postPostagem(Postagem postagem)
  + putPostagem(Postagem postagem)
  + deleteTema(Long id)
}
class Usuario {
  - id : Long
  - nome : String
  - usuario : String
  - senha : String
  - foto : String
  - postagem : List ~Postagem~
  + getAll()
  + getById(Long id)
  + autenticarUsuario(UsuarioLogin usuarioLogin)
  + cadastrarUsuario(Usuario usuario)
  + atualizarUsuario(Usuario usuario)
}
class UsuarioLogin{
  - id : Long
  - nome : String
  - usuario : String
  - senha : String
  - foto : String
  - token : String
}
Tema --> Postagem
Usuario --> Postagem
```

<br /><br />

# Referências sobre Golang

<br />

<a href="https://go.dev/" target="_blank">Site Oficial - Golang</a>

<a href="https://go.dev/doc/" target="_blank">Documentação Oficial - Golang</a>

<a href="https://pkg.go.dev/" target="_blank">Repositório de pacotes Oficial - Golang</a>

<a href="https://gorm.io/" target="_blank">Biblioteca GORM - Mapeamento Objeto Relacional - Golang</a>

<a href="https://github.com/spf13/viper" target="_blank">Pacote Viper - Gerenciador de configurações da API - Golang</a>

<a href="https://pkg.go.dev/encoding/json" target="_blank">Pacote JSON - Golang</a>

<a href="https://github.com/gorilla/mux" target="_blank">Pacote MUX - Rotas e Endpoints - Golang</a>

<a href="https://github.com/go-playground/validator" target="_blank">Go Validator V10 - Validação de dados - Golang</a>

<a href="https://github.com/swaggo/swag" target="_blank">Swag - Documentação com o Swagger 2.0 - Golang</a>
