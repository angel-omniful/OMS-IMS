package routes

import (
	"github.com/angel-omniful/ims/handlers"
	"github.com/omniful/go_commons/http"
)

func RegisterHealthRoutes(s *http.Server) {
	r:=s.Engine.Group("/api/ims")
	r.POST("/health", handlers.HealthCheck) // Health check route

}
