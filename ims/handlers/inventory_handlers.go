package handlers

import (
	"github.com/omniful/go_commons/http"
	"github.com/gin-gonic/gin"
	"github.com/angel-omniful/ims/models"
	"github.com/angel-omniful/ims/services"
)

func CreateInventory(c *gin.Context) {
	var req models.Inventory
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(int(http.StatusBadRequest), gin.H{"error": err.Error()})
		return
	}
	if err := services.CreateInventory(c, &req); err != nil {
		c.JSON(int(http.StatusInternalServerError), gin.H{"error": "Failed to create inventory"})
		return
	}
	c.JSON(int(http.StatusCreated), req)
}

func GetAllInventory(c *gin.Context) {
	inventory, err := services.GetAllInventory(c)
	if err != nil {
		c.JSON(int(http.StatusInternalServerError), gin.H{"error": "Failed to fetch inventory"})
		return
	}
	c.JSON(int(http.StatusOK), inventory)
}

func GetInventoryByID(c *gin.Context) {
	id := c.Param("id")
	inv, err := services.GetInventoryByID(c, id)
	if err != nil {
		c.JSON(int(http.StatusNotFound), gin.H{"error": "Inventory item not found"})
		return
	}
	c.JSON(int(http.StatusOK), inv)
}

func UpdateInventory(c *gin.Context) {
	id := c.Param("id")
	var req models.Inventory
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(int(http.StatusBadRequest), gin.H{"error": err.Error()})
		return
	}
	if err := services.UpdateInventory(c, id, &req); err != nil {
		c.JSON(int(http.StatusInternalServerError), gin.H{"error": "Failed to update inventory"})
		return
	}
	c.JSON(int(http.StatusOK), gin.H{"message": "Updated successfully"})
}

func DeleteInventory(c *gin.Context) {
	id := c.Param("id")
	if err := services.DeleteInventory(c, id); err != nil {
		c.JSON(int(http.StatusInternalServerError), gin.H{"error": "Failed to delete inventory"})
		return
	}
	c.JSON(int(http.StatusOK), gin.H{"message": "Deleted successfully"})
}

func Validate(c *gin.Context){
	hubid:=c.Param("hubid")
	skuid:=c.Param("skuid")

	bool,err:=services.ValidateInventory(c,skuid,hubid)
	if err!=nil{
		c.JSON(int(http.StatusInternalServerError), gin.H{"status":500 ,"error": "Failed to search inventory"})
		return
	}

	if bool{
		c.JSON(int(http.StatusOK), gin.H{"status":200 ,"message": "Found successfully"})
	}else
	{
		c.JSON(int(http.StatusNotFound), gin.H{"status":404 ,"error": "Not found in inventory"})
	}
}

func CheckInventory(c *gin.Context) {
	var req models.Inventory

	// Validate input
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(int(http.StatusBadRequest), gin.H{
			"status": http.StatusBadRequest,
			"error":  "invalid input",
		})
		return
	}

	// Check inventory from service
	found, err := services.CheckInventory(c, req.SkuID, req.HubID, req.Quantity)
	if err != nil {
		c.JSON(int(http.StatusInternalServerError), gin.H{
			"status": http.StatusInternalServerError,
			"error":  "internal error checking inventory",
		})
		return
	}

	if found {
		c.JSON(int(http.StatusOK), gin.H{
			"status":  http.StatusOK,
			"message": "inventory available",
		})
	} else {
		c.JSON(int(http.StatusNotFound), gin.H{
			"status": http.StatusNotFound,
			"error":  "inventory not found or insufficient",
		})
	}
}

