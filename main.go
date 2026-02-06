package main

import (
	"github.com/guilhermeonrails/api-go-gin/database"
	"github.com/guilhermeonrails/api-go-gin/routes"
	"github.com/joho/godotenv"
)

// @title API de Alunos
// @version 1.0
// @description API para gerenciamento de alunos
// @host localhost:8080
// @BasePath /
func main() {
	godotenv.Load()
	database.ConectaComBancoDeDados()
	routes.HandleRequest()
}
