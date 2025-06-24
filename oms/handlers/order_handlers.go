package handlers

import (
	"github.com/omniful/go_commons/http"
	"github.com/gin-gonic/gin"
	
	"github.com/angel-omniful/oms/model"
	"github.com/angel-omniful/oms/services"
	"context"
	"fmt"
	"os"
	"encoding/csv"
)


func CreateOrder(c *gin.Context) {
	var order model.Order
	if err := c.BindJSON(&order); err != nil {
		c.JSON(int(http.StatusBadRequest), gin.H{"error": "Invalid JSON"})
		return
	}
	err := services.CreateOrder(context.TODO(), &order)
	if err != nil {
		c.JSON(int(http.StatusInternalServerError), gin.H{"error": "Failed to create order"})
		return
	}
	c.JSON(int(http.StatusOK), gin.H{"message": "Order created successfully"})
}

func GetOrderById(c *gin.Context) {
	id := c.Param("id")
	orders, err := services.GetOrdersByField(context.TODO(), "id", id)
	if err != nil {
		c.JSON(int(http.StatusInternalServerError), gin.H{"error": "Failed to retrieve order"})
		return
	}
	c.JSON(int(http.StatusOK), orders)
}

func GetOrderByTenantId(c *gin.Context) {
	id := c.Param("tenant_id")
	orders, err := services.GetOrdersByField(context.TODO(), "tenant_id", id)
	if err != nil {
		c.JSON(int(http.StatusInternalServerError), gin.H{"error": "Failed to retrieve order"})
		return
	}
	c.JSON(int(http.StatusOK), orders)
}

func GetOrderBySellerId(c *gin.Context) {
	id := c.Param("seller_id")
	orders, err := services.GetOrdersByField(context.TODO(), "seller_id", id)
	if err != nil {
		c.JSON(int(http.StatusInternalServerError), gin.H{"error": "Failed to retrieve order"})
		return
	}
	c.JSON(int(http.StatusOK), orders)
}

func GetOrderByHubId(c *gin.Context) {
	id := c.Param("hub_id")
	orders, err := services.GetOrdersByField(context.TODO(), "hub_id", id)
	if err != nil {
		c.JSON(int(http.StatusInternalServerError), gin.H{"error": "Failed to retrieve order"})
		return
	}
	c.JSON(int(http.StatusOK), orders)
}

func GetOrderByStatus(c *gin.Context) {
	status := c.Param("status")
	orders, err := services.GetOrdersByField(context.TODO(), "status", status)
	if err != nil {
		c.JSON(int(http.StatusInternalServerError), gin.H{"error": "Failed to retrieve order"})
		return
	}
	c.JSON(int(http.StatusOK), orders)
}

func GetOrderByCreatedAt(c *gin.Context) {
	createdAt := c.Param("created_at")
	orders, err := services.GetOrdersByField(context.TODO(), "created_at", createdAt)
	if err != nil {
		c.JSON(int(http.StatusInternalServerError), gin.H{"error": "Failed to retrieve order"})
		return
	}
	c.JSON(int(http.StatusOK), orders)
}

func CreateFailedOrder(c *gin.Context) {
	var failed model.Order
	if err := c.BindJSON(&failed); err != nil {
		c.JSON(int(http.StatusBadRequest), gin.H{"error": "Invalid JSON"})
		return
	}
	err := services.CreateFailedOrder(context.TODO(), failed)
	if err != nil {
		c.JSON(int(http.StatusInternalServerError), gin.H{"error": "Failed to store failed order"})
		return
	}
	c.JSON(int(http.StatusOK), gin.H{"message": "Failed order logged"})
}

func GetAllFailedOrders(c *gin.Context) {
	failed, err := services.GetAllFailedOrders(context.TODO())
	if err != nil {
		c.JSON(int(http.StatusInternalServerError), gin.H{"error": "Failed to fetch failed orders"})
		return
	}
	c.JSON(int(http.StatusOK), failed)
}


func WriteOrdersToCSV(orders []model.Order, filepath string) error {
	// Create the file
	file, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("could not create file: %w", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header
	header := []string{"OrderID", "TenantID", "SellerID", "HubID","SkuId","Status","Quantity","CreatedAt"}
	if err := writer.Write(header); err != nil {
		return fmt.Errorf("could not write header to csv: %w", err)
	}

	// Write each order
	for _, order := range orders {
		row := []string{
			order.ID,
			order.TenantID,
			order.SellerID,
			order.HubID,
			order.SkuID,
			order.Status,
			fmt.Sprintf("%d", order.Quantity),
			fmt.Sprintf("%v", order.CreatedAt),
		}
		if err := writer.Write(row); err != nil {
			return fmt.Errorf("could not write row to csv: %w", err)
		}
	}

	return nil
}