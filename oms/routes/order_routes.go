package routes


import(
	 "github.com/omniful/go_commons/http"
	"github.com/angel-omniful/oms/handlers"
	"github.com/omniful/go_commons/config"
)

func RegisterOrderRoutes(s *http.Server) {
	r := s.Engine.Group("/api/oms/orders/get")
	r.Use(config.Middleware())
	r.POST("/create", handlers.CreateOrder)
	r.GET("/:id", handlers.GetOrderById)
	r.GET("/tenant/:tenant_id", handlers.GetOrderByTenantId)
	r.GET("/seller/:seller_id", handlers.GetOrderBySellerId)
	r.GET("/hub/:hub_id", handlers.GetOrderByHubId)
	r.GET("/status/:status", handlers.GetOrderByStatus)
	r.GET("/createdat/:created_at", handlers.GetOrderByCreatedAt)
	
	r.POST("/errors", handlers.CreateFailedOrder)
	r.GET("/errors", handlers.GetAllFailedOrders)
}

