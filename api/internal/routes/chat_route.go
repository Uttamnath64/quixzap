package routes

import (
	"github.com/Uttamnath64/quick-connect/api/internal/handlers"
	middleware "github.com/Uttamnath64/quick-connect/api/internal/middlewares"
)

func (routes *Routes) ChatRoutes() {
	handler := handlers.NewChat(routes.container)
	middle := middleware.New(routes.container)
	userGroup := routes.server.Group("user").Use(middle.AuthMiddleware())
	{
		userGroup.GET("/start", handler.Start)
		userGroup.GET("/:uuid", handler.GetAll)
		userGroup.GET("/", handler.GetAll)
	}
}
