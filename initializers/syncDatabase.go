package initializers

import (
	"university-api/course"
	"university-api/student"
	"university-api/user"

	"gorm.io/gorm"
)

func SyncDatabase(db *gorm.DB) error {
	err := db.AutoMigrate(user.User{}, student.Student{}, course.Course{})
	return err
}
