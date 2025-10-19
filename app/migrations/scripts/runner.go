package scripts

import (
	"github.com/Uttamnath64/quixzap/app/appcontext"
	"github.com/Uttamnath64/quixzap/app/models"
	"gorm.io/gorm"
)

type SeedFunc func(appCtx *appcontext.AppContext) error

func RunOnce(scriptName string, db *gorm.DB, fn func(*gorm.DB) error) error {
	var mv models.MigrationVersion
	if err := db.First(&mv, "script = ?", scriptName).Error; err == nil {
		return nil
	}

	if err := fn(db); err != nil {
		return err
	}

	return db.Create(&models.MigrationVersion{Script: scriptName}).Error
}

func RunMigrations(appCtx *appcontext.AppContext) error {

	seeds := []struct {
		Name string
		Func SeedFunc
	}{
		{"Currencies", currencies},
		{"Avatars", avatars},
		{"Admins", admins},
	}
	for _, seed := range seeds {
		appCtx.Logger.Info("🔄 Running migration:", "name", seed.Name)
		if err := seed.Func(appCtx); err != nil {
			appCtx.Logger.Fatal("❌ Migration failed:", seed.Name, "->", err)
			return err
		}
		appCtx.Logger.Info("✅ Migration done:", "name", seed.Name)
	}

	appCtx.Logger.Info("🎉 All migrations completed successfully.")
	return nil
}
