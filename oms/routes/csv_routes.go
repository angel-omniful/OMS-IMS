package routes

import (
	"github.com/angel-omniful/oms/handlers"
	"github.com/omniful/go_commons/http"
)

func RegisterCsvRoutes(s *http.Server) {
	r := s.Engine.Group("/api/oms/csv")
	r.POST("/upload", handlers.UploadCSVFileToS3)
	r.GET("/:filename",handlers.GenerateCsvUrl)
}






