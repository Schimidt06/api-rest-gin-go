package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	"github.com/golang-jwt/jwt/v5"
	"github.com/guilhermeonrails/api-go-gin/database"
	"github.com/guilhermeonrails/api-go-gin/models"
	"github.com/guilhermeonrails/api-go-gin/routes"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var r *gin.Engine

func SetupDasRotasDeTeste() {
	gin.SetMode(gin.ReleaseMode)
	rotas := routes.SetupRotas()
	r = rotas
}

func GeraTokenMock() string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"cpf": "12345678901",
		"rg":  "123456789",
	})
	tokenString, _ := token.SignedString([]byte("secret"))
	return "Bearer " + tokenString
}

func CriaAlunoMock() {
	aluno := models.Aluno{Nome: "Nome do Aluno Teste", CPF: "12345678901", RG: "123456789"}
	database.DB.Create(&aluno)
}

func DeletaAlunoMock() {
	var aluno models.Aluno
	database.DB.Delete(&aluno, "cpf = ?", "12345678901")
}

func SetupBancoDeDados() {
	// Use SQLite in-memory for testing
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("Falha ao conectar ao banco de dados")
	}
	database.DB = db
	database.DB.AutoMigrate(&models.Aluno{})
}

func TestVerificaStatusCodeDaSaudacaoComParametro(t *testing.T) {
	SetupDasRotasDeTeste()
	req, _ := http.NewRequest("GET", "/gui", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code, "Deveriam ser iguais")
	mockDaResposta := `{"API diz":"E ai gui, Tudo beleza?"}`
	respostaBody, _ := ioutil.ReadAll(resposta.Body)
	assert.Equal(t, mockDaResposta, string(respostaBody))
}

func TestListandoTodosOsAlunosHandler(t *testing.T) {
	SetupBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	SetupDasRotasDeTeste()
	req, _ := http.NewRequest("GET", "/alunos", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestBuscandoAlunoPorCPFHandler(t *testing.T) {
	SetupBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	SetupDasRotasDeTeste()
	req, _ := http.NewRequest("GET", "/alunos/cpf/12345678901", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestBuscandoAlunoPorIDHandler(t *testing.T) {
	SetupBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	SetupDasRotasDeTeste()
	// Need to find the ID of the created student first or just rely on it being the first if DB is fresh
	var aluno models.Aluno
	database.DB.Where("cpf = ?", "12345678901").First(&aluno)
	pathDaBusca := "/alunos/" + fmt.Sprint(aluno.ID)

	req, _ := http.NewRequest("GET", pathDaBusca, nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)

	var alunoMock models.Aluno
	json.Unmarshal(resposta.Body.Bytes(), &alunoMock)

	assert.Equal(t, "Nome do Aluno Teste", alunoMock.Nome)
	assert.Equal(t, "12345678901", alunoMock.CPF)
	assert.Equal(t, "123456789", alunoMock.RG)
	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestDeletandoAlunoHandler(t *testing.T) {
	SetupBancoDeDados()
	CriaAlunoMock()
	SetupDasRotasDeTeste()

	var aluno models.Aluno
	database.DB.Where("cpf = ?", "12345678901").First(&aluno)
	pathDaBusca := "/alunos/" + fmt.Sprint(aluno.ID)

	req, _ := http.NewRequest("DELETE", pathDaBusca, nil)
	req.Header.Add("Authorization", GeraTokenMock())
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)

	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestEditandoAlunoHandler(t *testing.T) {
	SetupBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	SetupDasRotasDeTeste()

	var aluno models.Aluno
	database.DB.Where("cpf = ?", "12345678901").First(&aluno)
	pathDaBusca := "/alunos/" + fmt.Sprint(aluno.ID)

	aluno.Nome = "Nome do Aluno Alterado"
	alunoJson, _ := json.Marshal(aluno)

	req, _ := http.NewRequest("PATCH", pathDaBusca, bytes.NewBuffer(alunoJson))
	req.Header.Add("Authorization", GeraTokenMock())
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)

	var alunoMockAtualizado models.Aluno
	json.Unmarshal(resposta.Body.Bytes(), &alunoMockAtualizado)

	assert.Equal(t, http.StatusOK, resposta.Code)
	assert.Equal(t, "Nome do Aluno Alterado", alunoMockAtualizado.Nome)
	assert.Equal(t, "12345678901", alunoMockAtualizado.CPF)
	assert.Equal(t, "123456789", alunoMockAtualizado.RG)
}
