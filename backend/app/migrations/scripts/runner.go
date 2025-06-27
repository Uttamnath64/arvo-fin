package script

import (
	"github.com/Uttamnath64/arvo-fin/app/models"
	"github.com/Uttamnath64/arvo-fin/app/storage"
	"gorm.io/gorm"
)

type SeedFunc func(container *storage.Container) error

func RunOnce(scriptName string, db *gorm.DB, fn func(*gorm.DB) error) error {
	var mv models.MigrationVersion
	if err := db.First(&mv, "script = ?", scriptName).Error; err == nil {
		return nil // already run
	}

	if err := fn(db); err != nil {
		return err
	}

	return db.Create(&models.MigrationVersion{Script: scriptName}).Error
}

func RunMigrations(container *storage.Container) error {

	seeds := []struct {
		Name string
		Func SeedFunc
	}{
		{"Currencies", currencies},
		{"Avatars", avatars},
		{"Admins", admins},
		{"Categories", categories},
	}
	for _, seed := range seeds {
		container.Logger.Info("🔄 Running migration:", "name", seed.Name)
		if err := seed.Func(container); err != nil {
			container.Logger.Fatal("❌ Migration failed:", seed.Name, "->", err)
			return err // Exit early on failure
		}
		container.Logger.Info("✅ Migration done:", "name", seed.Name)
	}

	container.Logger.Info("🎉 All migrations completed successfully.")
	return nil
}
