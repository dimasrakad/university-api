package user

import (
	"university-api/student"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"unique"`
	Password string
	Students    []student.Student
}
