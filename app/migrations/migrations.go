package main

import (
	"os"

	"fmt"

	"github.com/Uttamnath64/logger"
	"github.com/Uttamnath64/quixzap/app/appcontext"
	"github.com/Uttamnath64/quixzap/app/config"
	"github.com/Uttamnath64/quixzap/app/migrations/scripts"
	"github.com/Uttamnath64/quixzap/app/models"
)

var appCtx appcontext.AppContext

func init() {
	// load env
	env, err := config.LoadEnv(".env")
	if err != nil {
		fmt.Println("app-migration-env-error", err.Error())
		os.Exit(1)
	}
	appCtx.Env = &env

	// load logger
	appCtx.Logger = logger.New(env.Server.Environment, nil)

	// check DB
	if err := config.EnsureDB(appCtx.Env, appCtx.Logger); err != nil {
		appCtx.Logger.Error("migration-db-init-error", err.Error())
		os.Exit(1)
	}

	// load DB
	appCtx.MySQL, err = config.InitMySQL(appCtx.Env)
	if err != nil {
		appCtx.Logger.Error("migration-db-connect-error", err.Error())
		os.Exit(1)
	}
}

func main() {

	// migration database
	err := appCtx.MySQL.ReadWriteDB.AutoMigrate(
		&models.Admin{},
		&models.Avatar{},
		&models.BlockedIP{},
		&models.Business{},
		&models.ChatMessage{},
		&models.ChatSession{},
		&models.Currency{},
		&models.Member{},
		&models.MigrationVersion{},
		&models.Plan{},
		&models.Session{},
		&models.Subscription{},
		&models.Widget{},
	)
	if err != nil {
		appCtx.Logger.Error("app-migrate-config-error", "Failed to migrate the database", err)
		return
	}

	if err := scripts.RunMigrations(&appCtx); err != nil {
		os.Exit(1)
	}

	appCtx.Logger.Info("app-migrate-done", "message", "👍 Migration completed!")
}
