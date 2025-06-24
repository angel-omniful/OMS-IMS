package handlers


import (
	"log"
	"github.com/omniful/go_commons/http"
	"github.com/gin-gonic/gin"
	"github.com/angel-omniful/oms/model"
	"github.com/angel-omniful/oms/services"


)

// POST /api/oms/webhooks
func HandleOrderWebhook(c *gin.Context) {
	var req model.Webhook

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(int(http.StatusBadRequest), gin.H{"error": "Invalid payload", "details": err.Error()})
		return
	}

	if err := services.RegisterWebhook(c.Request.Context(), &req); err != nil {
		log.Println("Failed to register webhook: ", err)
		c.JSON(int(http.StatusInternalServerError), gin.H{"error": "Webhook registration failed"})
		return
	}

	c.JSON(int(http.StatusCreated), gin.H{"message": "Webhook registered successfully"})
}
