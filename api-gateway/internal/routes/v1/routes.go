package v1

import (
	"github.com/Uttamnath64/quixzap/app/appcontext"
	"github.com/gin-gonic/gin"
)

type RoutesV1 struct {
	appCtx *appcontext.AppContext
	rGroup *gin.RouterGroup // Router group for /api/v1
}

func New(appCtx *appcontext.AppContext, server *gin.Engine) *RoutesV1 {
	return &RoutesV1{
		appCtx: appCtx,
		rGroup: server.Group("/api/v1"),
	}
}

func (routes *RoutesV1) Handlers() {
	// Register all v1 route groups
	routes.MainRoutes()
	// routes.CustomerRoutes()
	// routes.AdminRoutes()
	// routes.WidgetRoutes()
	// routes.SuperAdminRoutes()
}
