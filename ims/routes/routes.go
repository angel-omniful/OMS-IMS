package routes

import "github.com/omniful/go_commons/http"

func RegisterAllRoutes(s *http.Server) {
	RegisterHealthRoutes(s)
	RegisterSKURoutes(s)
	RegisterHubRoutes(s)
	RegisterInventoryRoutes(s)
	RegisterSellerRoutes(s)
	RegisterTenantRoutes(s)
}
