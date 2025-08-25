package routes

import (
	"github.com/Uttamnath64/quick-connect/api/internal/handlers"
	middleware "github.com/Uttamnath64/quick-connect/api/internal/middlewares"
)

func (routes *Routes) AuthRoutes() {
	handler := handlers.NewAuth(routes.container)
	middle := middleware.New(routes.container)
	userGroup := routes.server.Group("/auth").Use(middle.AuthMiddleware())
	{
		userGroup.POST("/login", handler.Login)
		userGroup.POST("/register", handler.Register)
		userGroup.POST("/token", handler.Token)
	}
}
