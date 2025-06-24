package handlers

import (
	"github.com/omniful/go_commons/http"
	"github.com/gin-gonic/gin"
	"github.com/angel-omniful/ims/models"
	"github.com/angel-omniful/ims/services"
)

func CreateSKU(c *gin.Context) {
	var req models.SKU
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(int(http.StatusBadRequest), gin.H{"error": err.Error()})
		return
	}
	if err := services.CreateSKU(c, &req); err != nil {
		c.JSON(int(http.StatusInternalServerError), gin.H{"error": "Failed to create SKU"})
		return
	}
	c.JSON(int(http.StatusCreated), req)
}

func GetAllSKUs(c *gin.Context) {
	skus, err := services.GetAllSKUs(c)
	if err != nil {
		c.JSON(int(http.StatusInternalServerError), gin.H{"error": "Failed to fetch SKUs"})
		return
	}
	c.JSON(int(http.StatusOK), skus)
}

func GetSKUByID(c *gin.Context) {
	id := c.Param("id")
	sku, err := services.GetSKUByID(c, id)
	if err != nil {
		c.JSON(int(http.StatusNotFound), gin.H{"error": "SKU not found"})
		return
	}
	c.JSON(int(http.StatusOK), sku)
}

func UpdateSKU(c *gin.Context) {
	id := c.Param("id")
	var req models.SKU
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(int(http.StatusBadRequest), gin.H{"error": err.Error()})
		return
	}
	if err := services.UpdateSKU(c, id, &req); err != nil {
		c.JSON(int(http.StatusInternalServerError), gin.H{"error": "Failed to update SKU"})
		return
	}
	c.JSON(int(http.StatusOK), gin.H{"message": "Updated successfully"})
}

func DeleteSKU(c *gin.Context) {
	id := c.Param("id")
	if err := services.DeleteSKU(c, id); err != nil {
		c.JSON(int(http.StatusInternalServerError), gin.H{"error": "Failed to delete SKU"})
		return
	}
	c.JSON(int(http.StatusOK), gin.H{"message": "Deleted successfully"})
}

func GetSKUsByTenantID(c *gin.Context) {
	tenantID := c.Param("tenant_id")
	skus, err := services.GetSKUsByTenantID(c, tenantID)
	if err != nil {
		c.JSON(int(http.StatusInternalServerError), gin.H{"error": "Failed to fetch SKUs"})
		return
	}
	c.JSON(int(http.StatusOK), skus)
}

func GetSKUsBySellerID(c *gin.Context) {
	sellerID := c.Param("seller_id")
	skus, err := services.GetSKUsBySellerID(c, sellerID)
	if err != nil {
		c.JSON(int(http.StatusInternalServerError), gin.H{"error": "Failed to fetch SKUs"})
		return
	}
	c.JSON(int(http.StatusOK), skus)
}

func GetSKUsBySKUCode(c *gin.Context) {
	skuCode := c.Param("sku_code")
	skus, err := services.GetSKUsBySKUCode(c, skuCode)
	if err != nil {
		c.JSON(int(http.StatusInternalServerError), gin.H{"error": "Failed to fetch SKUs"})
		return
	}
	c.JSON(int(http.StatusOK), skus)
}
