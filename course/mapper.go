package course

func ConvertToCourseResponse(course Course) CourseResponse {
	return CourseResponse{
		ID:   course.ID,
		Name: course.Name,
	}
}
