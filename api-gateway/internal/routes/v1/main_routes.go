package v1

import (
	"github.com/Uttamnath64/quixzap/internal/handlers"
	"github.com/Uttamnath64/quixzap/internal/middlewares"
)

func (routes *RoutesV1) MainRoutes() {
	handler := handlers.NewMain(routes.container)
	middle := middlewares.New(routes.container)

	// No-auth routes for onboarding
	routes.rGroup.Use(middle.NoAuthMiddleware())
	{
		routes.rGroup.POST("/auth/register", handler.Register)
		routes.rGroup.POST("/auth/login", handler.Login)
		routes.rGroup.POST("/auth/send-otp", handler.SendOTP)
		routes.rGroup.POST("/auth/reset-password", handler.ResetPassword)
		routes.rGroup.POST("/auth/token", handler.Token)
		// routes.rGroup.GET("/customer/subscription/plans", handler.ListPlans)
	}
}
