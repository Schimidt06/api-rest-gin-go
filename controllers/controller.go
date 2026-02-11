package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/guilhermeonrails/api-go-gin/database"
	"github.com/guilhermeonrails/api-go-gin/models"
)

// Saudacoes exibe uma mensagem de boas-vindas
// @Summary Exibe uma mensagem de boas-vindas
// @Description Rota de saudação
// @Tags Geral
// @Accept  json
// @Produce  json
// @Param nome path string true "Nome da pessoa"
// @Success 200 {object} string
// @Router /{nome} [get]
func Saudacoes(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"API diz": "E ai " + nome + ", Tudo beleza?",
	})
}

// TodosAlunos exibe todos os alunos
// @Summary Exibe todos os alunos
// @Description Rota para listar todos os alunos
// @Tags Alunos
// @Accept  json
// @Produce  json
// @Param page query int false "Página da lista"
// @Param limit query int false "Limite de itens por página"
// @Success 200 {array} models.Aluno
// @Router /alunos [get]
func TodosAlunos(c *gin.Context) {
	var alunos []models.Aluno
	page := c.Query("page")
	limit := c.Query("limit")

	if page == "" && limit == "" {
		database.DB.Find(&alunos)
		c.JSON(200, alunos)
		return
	}

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	database.DB.Limit(intLimit).Offset(offset).Find(&alunos)
	c.JSON(200, alunos)
}

// CriarNovoAluno cria um novo aluno
// @Summary Cria um novo aluno
// @Description Rota para criar um novo aluno
// @Tags Alunos
// @Accept  json
// @Produce  json
// @Param aluno body models.Aluno true "Modelo de aluno"
// @Success 200 {object} models.Aluno
// @Router /alunos [post]
func CriarNovoAluno(c *gin.Context) {
	var aluno models.Aluno
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error()})
		return
	}
	if err := models.ValidaDadosDeAluno(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error()})
		return
	}
	database.DB.Create(&aluno)
	c.JSON(http.StatusOK, aluno)
}

// BuscarAlunoPorID busca um aluno pelo ID
// @Summary Busca um aluno pelo ID
// @Description Rota para buscar um aluno pelo ID
// @Tags Alunos
// @Accept  json
// @Produce  json
// @Param id path int true "ID do aluno"
// @Success 200 {object} models.Aluno
// @Failure 404 {object} string
// @Router /alunos/{id} [get]
func BuscarAlunoPorID(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.First(&aluno, id)
	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not Found": "Aluno não encontrado"})
		return
	}
	c.JSON(http.StatusOK, aluno)
}

// DeletarAluno deleta um aluno
// @Summary Deleta um aluno
// @Description Rota para deletar um aluno
// @Tags Alunos
// @Accept  json
// @Produce  json
// @Param id path int true "ID do aluno"
// @Success 200 {object} string
// @Failure 404 {object} string
// @Router /alunos/{id} [delete]
func DeletarAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.Delete(&aluno, id)
	c.JSON(http.StatusOK, gin.H{"data": "Aluno deletado com sucesso"})
}

// EditarAluno edita um aluno
// @Summary Edita um aluno
// @Description Rota para editar um aluno
// @Tags Alunos
// @Accept  json
// @Produce  json
// @Param id path int true "ID do aluno"
// @Param aluno body models.Aluno true "Modelo de aluno"
// @Success 200 {object} models.Aluno
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Router /alunos/{id} [patch]
func EditarAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.First(&aluno, id)
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error()})
		return
	}
	if err := models.ValidaDadosDeAluno(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error()})
		return
	}
	database.DB.Save(&aluno)
	c.JSON(http.StatusOK, aluno)
}

// BuscaAlunoPorCPF busca um aluno pelo CPF
// @Summary Busca um aluno pelo CPF
// @Description Rota para buscar um aluno pelo CPF
// @Tags Alunos
// @Accept  json
// @Produce  json
// @Param cpf path string true "CPF do aluno"
// @Success 200 {object} models.Aluno
// @Failure 404 {object} string
// @Router /alunos/cpf/{cpf} [get]
func BuscaAlunoPorCPF(c *gin.Context) {
	var aluno models.Aluno
	cpf := c.Param("cpf")
	database.DB.Where(&models.Aluno{CPF: cpf}).First(&aluno)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not Found": "Aluno não encontrado"})
		return
	}
	c.JSON(http.StatusOK, aluno)
}

// ExibePaginaIndex exibe a página principal
func ExibePaginaIndex(c *gin.Context) {
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"alunos": alunos,
	})
}

// RotaNaoEncontrada exibe a página de erro 404
func RotaNaoEncontrada(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.html", nil)
}
