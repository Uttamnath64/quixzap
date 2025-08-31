package main

import (
	"context"
	"time"

	"github.com/Uttamnath64/logger"
	"github.com/Uttamnath64/quick-connect/internal/app/config"
	"github.com/Uttamnath64/quick-connect/internal/app/storage"
	"github.com/Uttamnath64/quick-connect/internal/app/utils/requests"
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
	storage.NewContainer(&con, log, redis, &env)

	// Add a long-running task to keep the container alive
	log.Info("Consumer started, waiting for tasks...")
	for {
		// Example: Periodically log or process tasks
		log.Info("Consumer is running...")
		time.Sleep(10 * time.Second) // Adjust as needed
	}
}
