package main

import (
	"context"
	"fmt"

	"github.com/Uttamnath64/logger"
	"github.com/Uttamnath64/quixzap/internal/app/config"
	"github.com/Uttamnath64/quixzap/internal/app/storage"
	"github.com/Uttamnath64/quixzap/internal/app/utils/requests"
	"github.com/Uttamnath64/quixzap/internal/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize application
	ctx := context.Background()
	requests.NewResponse()

	// Load env
	env, err := config.LoadEnv(".env")
	if err != nil {
		logger.New("none", nil).Error("api-application-env", err.Error())
		return
	}

	// Set logger
	log := logger.New(env.Server.Environment, nil)

	// Load access and refresh keys
	err = config.LoadAccessAndRefreshKeys(&env)
	if err != nil {
		log.Error("api-application-accessAndRefreshKeys", err.Error())
		return
	}

	// Load config
	var con config.Config
	err = config.LoadConfig(env, &con)
	if err != nil {
		log.Error("api-application-config", err.Error())
		return
	}

	// Load Redis
	redis, err := storage.NewRedisClient(ctx, env.Redis.Addr, env.Redis.Password, env.Redis.DB)
	if err != nil {
		log.Error("api-application-redis", err.Error())
		return
	}

	// DI container
	container := storage.NewContainer(&con, log, redis, &env)

	// Setup Gin server
	server := gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{container.Env.Server.ClientOrigin}
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
	corsConfig.AllowCredentials = true
	server.Use(cors.New(corsConfig))

	// Setup routes
	routes.New(container, server).Handlers()

	// Run server
	if err := server.Run(fmt.Sprintf(":%d", container.Env.Server.Port)); err != nil {
		container.Logger.Error("api-application-server", err.Error())
		return
	}
}
