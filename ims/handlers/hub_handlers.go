package handlers

import (
	"github.com/omniful/go_commons/http"
	"github.com/gin-gonic/gin"
	"github.com/angel-omniful/ims/models"
	"github.com/angel-omniful/ims/services"
)

func CreateHub(c *gin.Context) {
	var req models.Hub
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(int(http.StatusBadRequest), gin.H{"error": err.Error()})
		return
	}
	if err := services.CreateHub(c, &req); err != nil {
		c.JSON(int(http.StatusInternalServerError), gin.H{"error": "Failed to create Hub"})
		return
	}
	c.JSON(int(http.StatusCreated), req)
}

func GetAllHubs(c *gin.Context) {
	hubs, err := services.GetAllHubs(c)
	if err != nil {
		c.JSON(int(http.StatusInternalServerError), gin.H{"error": "Failed to fetch Hubs"})
		return
	}
	c.JSON(int(http.StatusOK), hubs)
}

func GetHubByID(c *gin.Context) {
	id := c.Param("id")
	hub, err := services.GetHubByID(c, id)
	if err != nil {
		c.JSON(int(http.StatusNotFound), gin.H{"error": "Hub not found"})
		return
	}
	c.JSON(int(http.StatusOK), hub)
}

func UpdateHub(c *gin.Context) {
	id := c.Param("id")
	var req models.Hub
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(int(http.StatusBadRequest), gin.H{"error": err.Error()})
		return
	}
	if err := services.UpdateHub(c, id, &req); err != nil {
		c.JSON(int(http.StatusInternalServerError), gin.H{"error": "Failed to update Hub"})
		return
	}
	c.JSON(int(http.StatusOK), gin.H{"message": "Updated successfully"})
}

func DeleteHub(c *gin.Context) {
	id := c.Param("id")
	if err := services.DeleteHub(c, id); err != nil {
		c.JSON(int(http.StatusInternalServerError), gin.H{"error": "Failed to delete Hub"})
		return
	}
	c.JSON(int(http.StatusOK), gin.H{"message": "Deleted successfully"})
}
