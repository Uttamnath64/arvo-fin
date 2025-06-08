package script

import (
	"github.com/Uttamnath64/arvo-fin/app/models"
	"github.com/Uttamnath64/arvo-fin/app/storage"
	"gorm.io/gorm"
)

func admins(container *storage.Container) error {
	return RunOnce("20250608_admins", container.Config.ReadWriteDB, func(db *gorm.DB) error {
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
