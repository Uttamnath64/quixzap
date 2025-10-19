package v1

// func (routes *RoutesV1) AdminRoutes() {
// 	handler := handlers.NewAdmin(routes.container)
// 	middle := middlewares.New(routes.container)
// 	routes.rGroup.POST("/admin/login", middle.NoAuthMiddleware(), handler.AdminLogin)
// 	adminGroup := routes.rGroup.Group("/admin").Use(middle.Middleware())
// 	{
// 		adminGroup.GET("/profile", handler.GetAdminProfile)
// 		adminGroup.PUT("/profile", handler.UpdateAdminProfile)
// 		adminGroup.GET("/chats", handler.ListAssignedChats)
// 		adminGroup.GET("/chats/:chat_id", handler.GetChatDetails)
// 		adminGroup.POST("/chats/:chat_id/close", handler.CloseChatSession)
// 	}
// }
