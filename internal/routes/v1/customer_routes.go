package v1

import (
	"github.com/Uttamnath64/quixzap/internal/handlers"
	"github.com/Uttamnath64/quixzap/internal/middlewares"
)

func (routes *RoutesV1) CustomerRoutes() {
	handler := handlers.NewCustomer(routes.container)
	middle := middlewares.New(routes.container)
	customerGroup := routes.rGroup.Group("/customer").Use(middle.Middleware())
	{
		customerGroup.GET("/profile", handler.GetCustomerProfile)
		customerGroup.PUT("/profile", handler.UpdateCustomerProfile)
		customerGroup.POST("/subscription", handler.SubscribeToPlan)
		customerGroup.POST("/domain", handler.AddCustomerDomain)
		customerGroup.PUT("/widget", handler.ConfigureWidget)
		customerGroup.GET("/widget/code", handler.GenerateWidgetCode)
		customerGroup.GET("/admins", handler.ListAdmins)
		customerGroup.POST("/admins", handler.AddAdmin)
		customerGroup.PUT("/admins/:admin_id", handler.UpdateAdmin)
		customerGroup.DELETE("/admins/:admin_id", handler.RemoveAdmin)
		customerGroup.GET("/chats", handler.ListChatSessions)
		customerGroup.GET("/chats/:chat_id", handler.GetChatDetails)
		customerGroup.POST("/chats/:chat_id/close", handler.CloseChatSession)
		customerGroup.POST("/chats/:chat_id/block", handler.BlockChatOrIP)
		customerGroup.GET("/chats/export", handler.ExportChatHistory)
		customerGroup.POST("/autoclose", handler.SetAutoCloseRules)
	}
}
