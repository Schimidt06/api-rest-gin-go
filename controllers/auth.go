package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/guilhermeonrails/api-go-gin/models"
)

func Login(c *gin.Context) {
	var usuario models.Aluno

	if err := c.ShouldBindJSON(&usuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Mocking validation - In real app check DB/Hash
	if usuario.CPF == "12345678901" && usuario.RG == "123456789" {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"cpf": usuario.CPF,
			"rg":  usuario.RG,
		})

		tokenString, err := token.SignedString([]byte("secret"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Erro ao gerar token",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"token": tokenString,
		})
		return
	}

	c.JSON(http.StatusUnauthorized, gin.H{
		"error": "Credenciais inválidas",
	})
}
