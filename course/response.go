package course

// import "university-api/student"

type CourseResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	// Students []student.Student `json:"name"`
}
