package main

import (
	"context"
	"time"

	"github.com/Uttamnath64/logger"
	"github.com/Uttamnath64/quixzap/app/appcontext"
	"github.com/Uttamnath64/quixzap/app/config"
	"github.com/Uttamnath64/quixzap/app/utils/requests"
)

func main() {
	var appCtx appcontext.AppContext

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
	appCtx.Logger = logger.New(appCtx.Env.Server.Environment, nil)

	// load access and refresh keys
	err = config.LoadKeys(appCtx.Env)
	if err != nil {
		appCtx.Logger.Error("api-application-accessAndRefreshKeys", err.Error())
		return
	}

	// load mysql
	appCtx.MySQL, err = config.InitMySQL(appCtx.Env)
	if err != nil {
		appCtx.Logger.Error("api-application-mysql", err.Error())
		return
	}

	// load Redis
	appCtx.Redis, err = config.InitRedis(ctx, appCtx.Env.Redis.Addr, appCtx.Env.Redis.Password, appCtx.Env.Redis.DB)
	if err != nil {
		appCtx.Logger.Error("api-application-redis", err.Error())
		return
	}

	// add a long-running task to keep the container alive
	appCtx.Logger.Info("Consumer started, waiting for tasks...")
	for {
		// Example: Periodically log or process tasks
		appCtx.Logger.Info("Consumer is running...")
		time.Sleep(10 * time.Second) // Adjust as needed
	}
}
