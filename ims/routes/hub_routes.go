package routes

import (
	"github.com/angel-omniful/ims/handlers"
	"github.com/omniful/go_commons/http"
)

func RegisterHubRoutes(s *http.Server) {
	r := s.Engine.Group("/api/ims/hubs")

	r.POST("/", handlers.CreateHub)
	r.GET("/", handlers.GetAllHubs)
	r.GET("/:id", handlers.GetHubByID)
	r.PUT("/:id", handlers.UpdateHub)
	r.DELETE("/:id", handlers.DeleteHub)
}
