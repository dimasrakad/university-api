package student

import (
	"time"
	"university-api/course"
)

type StudentResponse struct {
	ID        int             `json:"id"`
	Name      string          `json:"name"`
	Address   string          `json:"address"`
	Major     string          `json:"major"`
	Phone     string          `json:"phone"`
	Email     string          `json:"email"`
	Birthdate time.Time       `json:"birthdate"`
	Courses   []course.Course `json:"courses"`
}
