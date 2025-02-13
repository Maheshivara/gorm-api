package main

import (
	"fmt"
	"gormCompose/src/driver"
	"gormCompose/src/routers"
	"gormCompose/src/services"
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
	domainUrl := os.Getenv("APP_DOMAIN_URL")
	serverEnv := os.Getenv("SERVER_ENV")
	if serverEnv == "DEV" {
		domainUrl = fmt.Sprintf("localhost:%s", serverPort)
	}
	driver.Migrate()
	docs.SwaggerInfo.Host = domainUrl

	app := gin.Default()
	app.SetTrustedProxies([]string{"localhost", "127.0.0.1", "proxy"})

	api := app.Group("/api")

	api.GET("/docs", func(c *gin.Context) {
		c.Redirect(301, "/api/docs/index.html")
	})
	api.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	foodRouter := routers.FoodRouter(services.FoodService())
	foodGroup := api.Group("/foods")
	foodGroup.POST("", foodRouter.Create)
	foodGroup.PUT("/:id", foodRouter.Update)
	foodGroup.GET("", foodRouter.List)
	foodGroup.DELETE("/:id", foodRouter.Delete)

	app.Run(fmt.Sprintf(":%s", serverPort))
}
