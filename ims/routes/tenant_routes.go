package routes

import (
	"github.com/angel-omniful/ims/handlers"
	"github.com/omniful/go_commons/http"
	"log"
)

func RegisterTenantRoutes(s *http.Server) {
	r := s.Engine.Group("/api/ims/tenants")

	r.POST("/", handlers.CreateTenant)
	r.GET("/", handlers.GetAllTenants)
	r.GET("/:id", handlers.GetTenantByID)
	r.PUT("/:id", handlers.UpdateTenant)
	r.DELETE("/:id", handlers.DeleteTenant)
	log.Println("Registered Seller routes")
}
