package routes

import (
	v1 "github.com/Uttamnath64/quixzap/api-gateway/internal/routes/v1"
	"github.com/Uttamnath64/quixzap/app/appcontext"
	"github.com/gin-gonic/gin"
)

type Routes struct {
	appCtx *appcontext.AppContext
	server *gin.Engine
	v1     *v1.RoutesV1
	// ws        *ws.WSRoutes
}

func New(appCtx *appcontext.AppContext, server *gin.Engine) *Routes {
	return &Routes{
		appCtx: appCtx,
		server: server,
		v1:     v1.New(appCtx, server),
		// ws:        ws.New(container, server),
	}
}

func (routes *Routes) Handlers() {
	// Serve static files for the main site
	routes.server.Static("/public/main", "./public/main")

	// Set up API routes under /api/v1
	routes.v1.Handlers()

	// Set up WebSocket routes under /ws
	// routes.ws.Handlers()
}
