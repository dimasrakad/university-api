package student

import "time"

type Student struct {
	ID        int
	Name      string
	Address   string
	Major     string
	Phone     string
	Email     string
	Birthdate time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
