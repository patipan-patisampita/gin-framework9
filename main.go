package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/patipan-patisampita/gin-framework9/configs"
	"github.com/patipan-patisampita/gin-framework9/routes"
)

func main() {
	router := SetupRouter()
	router.Run(":" + os.Getenv("GO_PORT")) //listen and server on 0.0.0.0:8080
}

func SetupRouter() *gin.Engine{

	//Load .env
	godotenv.Load(".env")
	
	//Connection db
	configs.Connection()
	router := gin.Default()

	apiV1 := router.Group("/api/v1")//127.0.0.1:3000/api/v1

	routes.InitHomeRoutes(apiV1)
	routes.InitUserRoutes(apiV1)//127.0.0.1:3000/api/v1/users/register
	return router
}
