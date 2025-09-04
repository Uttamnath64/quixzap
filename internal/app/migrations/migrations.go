package main

import (
	"os"

	"github.com/Uttamnath64/logger"
	"github.com/Uttamnath64/quixzap/internal/app/config"
	"github.com/Uttamnath64/quixzap/internal/app/migrations/scripts"
	"github.com/Uttamnath64/quixzap/internal/app/models"
	"github.com/Uttamnath64/quixzap/internal/app/storage"
)

func getContainer() (*storage.Container, error) {
	var err error
	var con config.Config

	// load env
	env, err := config.LoadEnv(".env")
	if err != nil {
		return nil, err
	}

	log := logger.New(env.Server.Environment, nil)

	// 👇 Create required databases if not exists
	if err := config.EnsureDatabases(&env, log); err != nil {
		log.Error("migration-db-init-error", err.Error())
		os.Exit(1)
	}

	// load config DB
	err = config.LoadConfig(env, &con)
	if err != nil {
		return nil, err
	}

	return storage.NewContainer(&con, log, nil, &env), nil
}

func main() {
	container, err := getContainer()
	if err != nil {
		logger.New("none", nil).Error("api-application-env", err.Error())
	}

	// migration database
	err = container.Config.DB.AutoMigrate(
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
		container.Logger.Error("app-migrate-config-error", "Failed to migrate the database", err)
		return
	}

	if err := scripts.RunMigrations(container); err != nil {
		os.Exit(1)
	}

	container.Logger.Info("app-migrate-done", "message", "👍 Migration completed!")
}
