package handlers

import (
	"github.com/omniful/go_commons/http"
	"github.com/gin-gonic/gin"
	"github.com/angel-omniful/ims/models"
	"github.com/angel-omniful/ims/services"
)

func CreateTenant(c *gin.Context) {
	var req models.Tenant
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(int(http.StatusBadRequest), gin.H{"error": err.Error()})
		return
	}
	if err := services.CreateTenant(c, &req); err != nil {
		c.JSON(int(http.StatusInternalServerError), gin.H{"error": "Failed to create tenant"})
		return
	}
	c.JSON(int(http.StatusCreated), req)
}

func GetAllTenants(c *gin.Context) {
	tenants, err := services.GetAllTenants(c)
	if err != nil {
		c.JSON(int(http.StatusInternalServerError), gin.H{"error": "Failed to fetch tenants"})
		return
	}
	c.JSON(int(http.StatusOK), tenants)
}

func GetTenantByID(c *gin.Context) {
	id := c.Param("id")
	tenant, err := services.GetTenantByID(c, id)
	if err != nil {
		c.JSON(int(http.StatusNotFound), gin.H{"error": "Tenant not found"})
		return
	}
	c.JSON(int(http.StatusOK), tenant)
}

func UpdateTenant(c *gin.Context) {
	id := c.Param("id")
	var req models.Tenant
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(int(http.StatusBadRequest), gin.H{"error": err.Error()})
		return
	}
	if err := services.UpdateTenant(c, id, &req); err != nil {
		c.JSON(int(http.StatusInternalServerError), gin.H{"error": "Failed to update tenant"})
		return
	}
	c.JSON(int(http.StatusOK), gin.H{"message": "Updated successfully"})
}

func DeleteTenant(c *gin.Context) {
	id := c.Param("id")
	if err := services.DeleteTenant(c, id); err != nil {
		c.JSON(int(http.StatusInternalServerError), gin.H{"error": "Failed to delete tenant"})
		return
	}
	c.JSON(int(http.StatusOK), gin.H{"message": "Deleted successfully"})
}
