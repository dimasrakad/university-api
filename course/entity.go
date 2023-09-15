package course

type Course struct {
	ID   int
	Name string `gorm:"type:varchar(255);not null;size:100"`
}
