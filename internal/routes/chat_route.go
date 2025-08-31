package routes

import (
	"github.com/Uttamnath64/quick-connect/internal/handlers"
	middleware "github.com/Uttamnath64/quick-connect/internal/middlewares"
)

func (routes *Routes) ChatRoutes() {
	handler := handlers.NewChat(routes.container)
	middle := middleware.New(routes.container)
	group := routes.server.Group("chat").Use(middle.AuthMiddleware())
	{
		group.POST("/", handler.Create)
	}
}
