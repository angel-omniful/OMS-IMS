package handlers

import(

	"github.com/omniful/go_commons/http"
	"github.com/gin-gonic/gin"
	
)

func HealthCheck(c *gin.Context) {
	c.JSON(int(http.StatusOK), gin.H{
		"status": "ok",
	})
}