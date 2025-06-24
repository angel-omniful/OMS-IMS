package routes


import(
	 "github.com/omniful/go_commons/http"
	"github.com/angel-omniful/oms/handlers"
)

func RegisterWebhookRoutes(s *http.Server) {
	r := s.Engine.Group("/api/oms/webhooks")

	r.POST("/", handlers.HandleOrderWebhook)
}

