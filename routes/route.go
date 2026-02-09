package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/guilhermeonrails/api-go-gin/controllers"
	"github.com/guilhermeonrails/api-go-gin/middleware"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func HandleRequest() {
	r := SetupRotas()
	r.Run()
}

func SetupRotas() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.POST("/login", controllers.Login)

	protected := r.Group("/")
	protected.Use(middleware.Autentica())
	{
		protected.POST("/alunos", controllers.CriarNovoAluno)
		protected.DELETE("/alunos/:id", controllers.DeletarAluno)
		protected.PATCH("/alunos/:id", controllers.EditarAluno)
	}

	r.GET("/:nome", controllers.Saudacoes)
	r.GET("/alunos", controllers.TodosAlunos)
	r.GET("/alunos/:id", controllers.BuscarAlunoPorID)
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoPorCPF)
	r.GET("/index", controllers.ExibePaginaIndex)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return r
}
