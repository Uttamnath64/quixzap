package v1

import (
	"github.com/Uttamnath64/quixzap/internal/handlers"
	"github.com/Uttamnath64/quixzap/internal/middlewares"
)

func (routes *RoutesV1) MainRoutes() {
	handler := handlers.NewAuth(routes.container)
	middle := middlewares.New(routes.container)

	// No-auth routes for onboarding
	routes.rGroup.Use(middle.NoAuthMiddleware())
	{
		routes.rGroup.POST("/auth/register", handler.Register)
		routes.rGroup.POST("/auth/login", handler.Login)
		routes.rGroup.GET("/customer/subscription/plans", handler.ListPlans)
	}
}
