package student

import (
	"time"
	"university-api/course"

	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	ID        int
	Name      string
	Address   string
	Major     string
	Phone     string
	Email     string
	Courses   []course.Course `gorm:"many2many:student_course;"`
	Birthdate time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
