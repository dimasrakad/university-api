package initializers

import (
	"university-api/student"

	"gorm.io/gorm"
)

func SyncDatabase(db *gorm.DB) error {
	err := db.AutoMigrate(student.Student{})
	return err
}
