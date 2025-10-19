package scripts

import (
	"github.com/Uttamnath64/quixzap/app/appcontext"
	"github.com/Uttamnath64/quixzap/app/models"
	"gorm.io/gorm"
)

func admins(appCtx *appcontext.AppContext) error {
	return RunOnce("20250608_admins", appCtx.MySQL.ReadWriteDB, func(db *gorm.DB) error {
		admins := []models.Admin{
			{AvatarId: 8, Name: "Admin", Email: "admin@arvofin.com", Username: "arvo.admin", Password: "$2a$10$N7RKD8VqYHY4kbGWmfElBOs/wPfdnGldKAoRGOPa7ERbxEzeEOl1u"},
		}
		for _, a := range admins {
			if err := db.FirstOrCreate(&a, models.Admin{Username: a.Username}).Error; err != nil {
				return err
			}
		}
		return nil
	})
}
