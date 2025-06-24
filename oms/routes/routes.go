package routes

import "github.com/omniful/go_commons/http"

func RegisterAllRoutes(s *http.Server) {
	RegisterOrderRoutes(s)
	RegisterWebhookRoutes(s)
	RegisterCsvRoutes(s)

}
