package routes

import (
	"github.com/angel-omniful/ims/handlers"
	"github.com/omniful/go_commons/http"
)

func RegisterInventoryRoutes(s *http.Server) {
	r := s.Engine.Group("/api/ims/inventory")

	r.POST("/", handlers.CreateInventory)
	r.GET("/", handlers.GetAllInventory)
	r.GET("/:id", handlers.GetInventoryByID)
	r.PUT("/:id", handlers.UpdateInventory)
	r.DELETE("/:id", handlers.DeleteInventory)
	r.GET("/hubsku/:hubid/:skuid",handlers.Validate)
	r.POST("/check",handlers.CheckInventory)
}
