package scripts

import (
	"github.com/Uttamnath64/quixzap/internal/app/common/types"
	"github.com/Uttamnath64/quixzap/internal/app/models"
	"github.com/Uttamnath64/quixzap/internal/app/storage"
	"gorm.io/gorm"
)

func avatars(container *storage.Container) error {
	return RunOnce("20250608_avatars", container.Config.DB, func(db *gorm.DB) error {
		avatars := []models.Avatar{

			// Default Icons (Generic)
			{Name: "Star", Icon: "⭐", Type: types.AvatarTypeDefault},
			{Name: "Sparkles", Icon: "✨", Type: types.AvatarTypeDefault},
			{Name: "Fire", Icon: "🔥", Type: types.AvatarTypeDefault},
			{Name: "Heart", Icon: "❤️", Type: types.AvatarTypeDefault},
			{Name: "Globe", Icon: "🌍", Type: types.AvatarTypeDefault},
			{Name: "Rocket", Icon: "🚀", Type: types.AvatarTypeDefault},

			// Business Avatars (For Super User / Businesses)
			{Name: "Business", Icon: "🏢", Type: types.AvatarTypeBusiness},
			{Name: "Store", Icon: "🏬", Type: types.AvatarTypeBusiness},
			{Name: "Shop", Icon: "🛍️", Type: types.AvatarTypeBusiness},
			{Name: "Online Service", Icon: "💻", Type: types.AvatarTypeBusiness},
			{Name: "E-Commerce", Icon: "🛒", Type: types.AvatarTypeBusiness},
			{Name: "Support HQ", Icon: "🎯", Type: types.AvatarTypeBusiness},
			{Name: "Verified Business", Icon: "✅", Type: types.AvatarTypeBusiness},

			// Member / Agent Avatars (Support Team)
			{Name: "Agent", Icon: "🧑‍💻", Type: types.AvatarTypeUser},
			{Name: "Support", Icon: "🎧", Type: types.AvatarTypeUser},
			{Name: "Moderator", Icon: "🛡️", Type: types.AvatarTypeUser},
			{Name: "Manager", Icon: "👨‍💼", Type: types.AvatarTypeUser},
			{Name: "Team Lead", Icon: "👩‍💼", Type: types.AvatarTypeUser},
			{Name: "Assistant", Icon: "🤝", Type: types.AvatarTypeUser},
			{Name: "Intern", Icon: "📚", Type: types.AvatarTypeUser},

			// Chat Status Avatars
			{Name: "Online", Icon: "🟢", Type: types.AvatarTypeStatus},
			{Name: "Offline", Icon: "🔴", Type: types.AvatarTypeStatus},
			{Name: "Busy", Icon: "🟠", Type: types.AvatarTypeStatus},
			{Name: "Away", Icon: "🟡", Type: types.AvatarTypeStatus},
			{Name: "In Call", Icon: "📞", Type: types.AvatarTypeStatus},
			{Name: "Typing", Icon: "⌨️", Type: types.AvatarTypeStatus},

			// Dashboard / Panel Avatars
			{Name: "Main Dashboard", Icon: "📊", Type: types.AvatarTypePanel},
			{Name: "Support Dashboard", Icon: "📋", Type: types.AvatarTypePanel},
			{Name: "Team Panel", Icon: "👥", Type: types.AvatarTypePanel},
			{Name: "Analytics", Icon: "📈", Type: types.AvatarTypePanel},
			{Name: "Settings", Icon: "⚙️", Type: types.AvatarTypePanel},
			{Name: "Notifications", Icon: "🔔", Type: types.AvatarTypePanel},
			{Name: "Chats", Icon: "💬", Type: types.AvatarTypePanel},
			{Name: "Tickets", Icon: "🎟️", Type: types.AvatarTypePanel},
		}
		for _, a := range avatars {
			if err := db.FirstOrCreate(&a, models.Avatar{Name: a.Name, Type: a.Type}).Error; err != nil {
				return err
			}
		}
		return nil
	})
}
