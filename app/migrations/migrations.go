package main

import (
	"os"

	"github.com/Uttamnath64/logger"
	"github.com/Uttamnath64/quick-connect/app/config"
	"github.com/Uttamnath64/quick-connect/app/models"
	"github.com/Uttamnath64/quick-connect/app/storage"
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
		&models.Chat{},
		&models.Session{},
		&models.Admin{},
		&models.Message{},
	)
	if err != nil {
		container.Logger.Error("app-migrate-config-error", "Failed to migrate the database", err)
		return
	}

	container.Logger.Info("app-migrate-done", "message", "👍 Migration completed!")
}
