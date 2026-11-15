package routes

import (
	"football-club/internal/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()
	//Allow all origins (in dev)
	r.Use(cors.Default())

	r.StaticFile("/", "./index.html")

	//	api
	api := r.Group("/api")
	{
		api.GET("/players", handlers.GetPlayers)
		api.GET("/players/:id", handlers.GetPlayerById)
		api.POST("/players", handlers.CreatePlayer)
		api.PUT("/players/:id", handlers.UpdatePlayer)
		api.DELETE("/players/:id", handlers.DeletePlayer)
	}
	return r
}
