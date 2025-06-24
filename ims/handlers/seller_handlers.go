package handlers

import (
	"github.com/omniful/go_commons/http"
	"github.com/gin-gonic/gin"
	"github.com/angel-omniful/ims/models"
	"github.com/angel-omniful/ims/services"
)

func CreateSeller(c *gin.Context) {
	var req models.Seller
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(int(http.StatusBadRequest), gin.H{"error": err.Error()})
		return
	}
	if err := services.CreateSeller(c, &req); err != nil {
		c.JSON(int(http.StatusInternalServerError), gin.H{"error": "Failed to create seller"})
		return
	}
	c.JSON(int(http.StatusCreated), req)
}

func GetAllSellers(c *gin.Context) {
	sellers, err := services.GetAllSellers(c)
	if err != nil {
		c.JSON(int(http.StatusInternalServerError), gin.H{"error": "Failed to fetch sellers"})
		return
	}
	c.JSON(int(http.StatusOK), sellers)
}

func GetSellerByID(c *gin.Context) {
	id := c.Param("id")
	seller, err := services.GetSellerByID(c, id)
	if err != nil {
		c.JSON(int(http.StatusNotFound), gin.H{"error": "Seller not found"})
		return
	}
	c.JSON(int(http.StatusOK), seller)
}

func UpdateSeller(c *gin.Context) {
	id := c.Param("id")
	var req models.Seller
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(int(http.StatusBadRequest), gin.H{"error": err.Error()})
		return
	}
	if err := services.UpdateSeller(c, id, &req); err != nil {
		c.JSON(int(http.StatusInternalServerError), gin.H{"error": "Failed to update seller"})
		return
	}
	c.JSON(int(http.StatusOK), gin.H{"message": "Updated successfully"})
}

func DeleteSeller(c *gin.Context) {
	id := c.Param("id")
	if err := services.DeleteSeller(c, id); err != nil {
		c.JSON(int(http.StatusInternalServerError), gin.H{"error": "Failed to delete seller"})
		return
	}
	c.JSON(int(http.StatusOK), gin.H{"message": "Deleted successfully"})
}
