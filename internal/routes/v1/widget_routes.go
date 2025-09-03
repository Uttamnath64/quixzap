package v1

import (
	"github.com/Uttamnath64/quixzap/internal/handlers"
	"github.com/Uttamnath64/quixzap/internal/middlewares"
)

func (routes *RoutesV1) WidgetRoutes() {
	handler := handlers.NewWidget(routes.container)
	middle := middlewares.New(routes.container)
	widgetGroup := routes.rGroup.Group("/widget/:customer_id").Use(middle.WidgetTokenMiddleware())
	{
		widgetGroup.GET("/config", handler.GetWidgetConfig)
		widgetGroup.POST("/message", handler.SendFirstMessage)
	}
}
