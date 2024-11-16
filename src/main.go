package main

import (
	"fmt"
	"gormCompose/src/driver"
	"gormCompose/src/routers"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	docs "gormCompose/src/docs"
)

//	@title			API de Comida.
//	@version		1.0
//	@description	API simples para demonstrar o uso de containers Docker e suas integrações usando Docker Compose.

//	@contact.name	Luis Gabriel da Costa Silva
//	@contact.email	lgcs10@aluno.ifal.edu.br

//	@BasePath	/api

// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
func main() {
	godotenv.Load(".env")
	serverPort := os.Getenv("SERVER_PORT")
	driver.Migrate()
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%s", serverPort)

	app := gin.Default()
	app.SetTrustedProxies([]string{"localhost", "127.0.0.1"})

	app.GET("/docs", func(c *gin.Context) {
		c.Redirect(301, "/docs/index.html")
	})
	app.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := app.Group("/api")

	foodRouter := routers.FoodRouter()
	foodGroup := api.Group("/foods")
	foodGroup.POST("", foodRouter.Create)
	foodGroup.GET("", foodRouter.List)

	app.Run(fmt.Sprintf(":%s", serverPort))
}
