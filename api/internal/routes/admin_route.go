package routes

import (
	"github.com/Uttamnath64/quick-connect/api/internal/handlers"
	middleware "github.com/Uttamnath64/quick-connect/api/internal/middlewares"
)

func (routes *Routes) AdminRoutes() {
	handler := handlers.NewAdmin(routes.container)
	middle := middleware.New(routes.container)
	userGroup := routes.server.Group("admin").Use(middle.Middleware())
	{
		userGroup.GET("/:id", handler.Get)
		userGroup.GET("/", handler.GetAll)
		userGroup.PUT("/:id", handler.Update)
		userGroup.POST("/", handler.Create)
		userGroup.POST("/:id/:block", handler.Block)
	}
}
