package main

import (
	"context"
	"fmt"

	"github.com/Uttamnath64/logger"
	"github.com/Uttamnath64/quixzap/api-gateway/internal/routes"
	"github.com/Uttamnath64/quixzap/app/appcontext"
	"github.com/Uttamnath64/quixzap/app/config"
	"github.com/Uttamnath64/quixzap/app/utils/requests"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	var appCtx *appcontext.AppContext

	// initialize application
	ctx := context.Background()
	requests.NewResponse()

	// load env
	env, err := config.LoadEnv(".env")
	if err != nil {
		logger.New("none", nil).Error("api-application-env", err.Error())
		return
	}
	appCtx.Env = &env

	// load logger
	appCtx.Logger = logger.New(env.Server.Environment, nil)

	// load access and refresh keys
	err = config.LoadKeys(appCtx.Env)
	if err != nil {
		appCtx.Logger.Error("api-application-accessAndRefreshKeys", err.Error())
		return
	}

	// load redis
	appCtx.Redis, err = config.InitRedis(ctx, appCtx.Env.Redis.Addr, appCtx.Env.Redis.Addr, appCtx.Env.Redis.DB)
	if err != nil {
		appCtx.Logger.Error("api-application-redis", err.Error())
		return
	}

	// Setup Gin server
	server := gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{appCtx.Env.Server.ClientOrigin}
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
	corsConfig.AllowCredentials = true
	server.Use(cors.New(corsConfig))

	// Setup routes
	routes.New(appCtx, server).Handlers()

	// Run server
	if err := server.Run(fmt.Sprintf(":%d", appCtx.Env.Server.Port)); err != nil {
		appCtx.Logger.Error("api-application-server", err.Error())
		return
	}
}
