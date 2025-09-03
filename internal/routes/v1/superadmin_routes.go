package v1

// func (routes *RoutesV1) SuperAdminRoutes() {
// 	handler := handlers.NewSuperAdmin(routes.container)
// 	middle := middlewares.New(routes.container)
// 	routes.rGroup.POST("/superadmin/login", middle.NoAuthMiddleware(), handler.SuperAdminLogin)
// 	superAdminGroup := routes.rGroup.Group("/superadmin").Use(middle.Middleware())
// 	{
// 		superAdminGroup.GET("/customers", handler.ListCustomers)
// 		superAdminGroup.GET("/customers/:customer_id", handler.GetCustomerDetails)
// 		superAdminGroup.POST("/customers/:customer_id/block", handler.BlockCustomer)
// 		superAdminGroup.GET("/sessions", handler.ListAllSessions)
// 		superAdminGroup.POST("/sessions/:session_id/block", handler.BlockSession)
// 		superAdminGroup.GET("/usage", handler.GetUsageStats)
// 	}
// }
