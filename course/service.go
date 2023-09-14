package course

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

type Service interface {
	FindAll() ([]Course, error)
	FindByID(ID int) (Course, error)
	Create(courseRequest CourseRequest) (Course, error)
	Update(ID int, CourseRequest CourseRequest) (Course, error)
	Delete(ID int) (Course, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Course, error) {
	courses, err := s.repository.FindAll()
	return courses, err
}

func (s *service) FindByID(ID int) (Course, error) {
	course, err := s.repository.FindByID(ID)
	return course, err
}

func (s *service) Create(courseRequest CourseRequest) (Course, error) {
	course := Course{
		Name: courseRequest.Name,
		// UserID:  UserID,
	}

	newCourse, err := s.repository.Create(course)
	return newCourse, err
}

func (s *service) Update(ID int, courseRequest CourseRequest) (Course, error) {

	// Validasi ID
	if ID <= 0 {
		return Course{}, errors.New("ID mahasiswa tidak valid")
	}

	// Validasi data mahasiswa
	v := validator.New()
	if err := v.Struct(courseRequest); err != nil {
		return Course{}, err
	}

	// Temukan mahasiswa berdasarkan ID
	course, err := s.repository.FindByID(ID)
	if err != nil {
		return Course{}, err
	}

	if courseRequest.Name != "" {
		course.Name = courseRequest.Name
	}

	newStudent, err := s.repository.Update(course)
	return newStudent, err
}

func (s *service) Delete(ID int) (Course, error) {
	student, err := s.repository.FindByID(ID)
	_, err = s.repository.Delete(student)

	return student, err
}
