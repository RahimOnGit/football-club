package handlers

import (
	"football-club/internal/database"
	"football-club/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetPlayers Get all players
func GetPlayers(c *gin.Context) {
	var players []models.Player
	database.DB.Find(&players)
	c.JSON(http.StatusOK, players)
}

// GetPlayerById find a player by id
func GetPlayerById(c *gin.Context) {
	id := c.Param("id")
	var player models.Player
	result := database.DB.Find(&player, id)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, player)

}

// CreatePlayer create new player
func CreatePlayer(c *gin.Context) {
	var newPlayer models.Player
	if err := c.ShouldBind(&newPlayer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	database.DB.Create(&newPlayer)
	c.JSON(http.StatusOK, newPlayer)
}

// UpdatePlayer update player
func UpdatePlayer(c *gin.Context) {
	id := c.Param("id")
	//bind
	var payload models.Player
	if err := c.ShouldBind(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//	ensure the row exists
	var player models.Player
	if err := database.DB.Find(&player, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// update using model(&).updates(&)
	if err := database.DB.Model(&player).Updates(&payload).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, player)

}

func DeletePlayer(c *gin.Context) {
	id := c.Param("id")
	var player models.Player
	if err := database.DB.Find(&player, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	result := database.DB.Delete(&player)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error, failed to Delete": result.Error.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"message": "Player deleted"})

}
