package routes

import (
	"github.com/angel-omniful/ims/handlers"
	"github.com/omniful/go_commons/http"
)

func RegisterSKURoutes(s *http.Server) {
	r := s.Engine.Group("/api/ims/skus")
	
		r.POST("/", handlers.CreateSKU)
		r.GET("/", handlers.GetAllSKUs)
		r.GET("/:id", handlers.GetSKUByID)
		r.PUT("/:id", handlers.UpdateSKU)
		r.DELETE("/:id", handlers.DeleteSKU)
		r.GET("/filter/tenant/:tenant_id", handlers.GetSKUsByTenantID)
		r.GET("/filter/seller/:seller_id", handlers.GetSKUsBySellerID)
		r.GET("/filter/code/:sku_code", handlers.GetSKUsBySKUCode)
	
}
