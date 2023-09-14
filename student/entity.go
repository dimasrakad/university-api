package student

import "time"

type Student struct {
	ID        int
	Name      string
	Address   string
	Major     string
	Phone     string
	CreatedAt time.Time
	UpdatedAt time.Time
	// UserID    uint
}
