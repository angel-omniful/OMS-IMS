package routes

import (
	"github.com/angel-omniful/ims/handlers"
	"github.com/omniful/go_commons/http"
	
)

func RegisterSellerRoutes(s *http.Server) {
	
	r := s.Engine.Group("/api/ims/sellers")

	r.POST("/", handlers.CreateSeller)
	r.GET("/", handlers.GetAllSellers)
	r.GET("/:id", handlers.GetSellerByID)
	r.PUT("/:id", handlers.UpdateSeller)
	r.DELETE("/:id", handlers.DeleteSeller)
	
}
