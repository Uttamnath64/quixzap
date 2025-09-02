package v1

import (
	"github.com/Uttamnath64/quixzap/internal/handlers"
	"github.com/Uttamnath64/quixzap/internal/middlewares"
)

func (routes *RoutesV1) SuperAdminRoutes() {
	handler := handlers.NewSuperAdmin(routes.container)
	middle := middlewares.New(routes.container)

	// Super admin routes
	routes.rGroup.POST("/superadmin/login", middle.NoAuthMiddleware(), handler.SuperAdminLogin)
	superAdminGroup := routes.rGroup.Group("/superadmin").Use(middle.SuperAdminAuthMiddleware())
	{
		superAdminGroup.GET("/customers", handler.ListCustomers)
		superAdminGroup.GET("/customers/:customer_id", handler.GetCustomerDetails)
		superAdminGroup.POST("/customers/:customer_id/block", handler.BlockCustomer)
		superAdminGroup.GET("/sessions", handler.ListAllSessions)
		superAdminGroup.POST("/sessions/:session_id/block", handler.BlockSession)
		superAdminGroup.GET("/usage", handler.GetUsageStats)
	}
}
