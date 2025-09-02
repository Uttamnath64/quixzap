package v1

import (
	"github.com/Uttamnath64/quixzap/internal/handlers"
	"github.com/Uttamnath64/quixzap/internal/middlewares"
)

func (routes *RoutesV1) AdminRoutes() {
	handler := handlers.NewAdmin(routes.container)
	middle := middlewares.New(routes.container)

	// Admin routes
	routes.rGroup.POST("/admin/login", middle.NoAuthMiddleware(), handler.AdminLogin)
	adminGroup := routes.rGroup.Group("/admin").Use(middle.AdminAuthMiddleware())
	{
		adminGroup.GET("/profile", handler.GetAdminProfile)
		adminGroup.PUT("/profile", handler.UpdateAdminProfile)
		adminGroup.GET("/chats", handler.ListAssignedChats)
		adminGroup.GET("/chats/:chat_id", handler.GetChatDetails)
		adminGroup.POST("/chats/:chat_id/close", handler.CloseChatSession)
	}
}
