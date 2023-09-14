package student

func ConvertToStudentResponse(student Student) StudentResponse {
	return StudentResponse{
		ID:      student.ID,
		Name:    student.Name,
		Address: student.Address,
		Major:   student.Major,
		Phone:   student.Phone,
	}
}
