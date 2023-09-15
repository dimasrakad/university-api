package student

import (
	"time"
	"university-api/course"

	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	ID        int
	Name      string          `gorm:"type:varchar(255);not null;size:100"`
	Address   string          `gorm:"type:varchar(255);not null;size:255"`
	Major     string          `gorm:"type:varchar(255);not null;size:50"`
	Phone     string          `gorm:"type:varchar(20);not null;size:20"`
	Email     string          `gorm:"type:varchar(255);not null;unique;size:30"`
	Courses   []course.Course `gorm:"many2many:student_course;constraint:OnDelete:CASCADE"`
	Birthdate time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
