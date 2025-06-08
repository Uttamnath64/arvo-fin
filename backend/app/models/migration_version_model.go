package models

import "time"

type MigrationVersion struct {
	ID        uint   `gorm:"primaryKey"`
	Script    string `gorm:"uniqueIndex;size:100"`
	CreatedAt time.Time
}

func (m MigrationVersion) GetName() string {
	return "migration_versions"
}
