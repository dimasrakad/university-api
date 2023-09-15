package student

import (
	"errors"
	"time"

	"github.com/go-playground/validator/v10"
)

type Service interface {
	FindAll() ([]Student, error)
	FindByID(ID int) (Student, error)
	Create(studentRequest StudentRequest) (Student, error)
	Update(ID int, StudentRequest StudentRequest) (Student, error)
	Delete(ID int) (Student, error)
	AddCourseToStudent(studentID int, courseID int) (Student, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Student, error) {
	students, err := s.repository.FindAll()
	return students, err
}

func (s *service) FindByID(ID int) (Student, error) {
	student, err := s.repository.FindByID(ID)
	return student, err
}

func (s *service) Create(studentRequest StudentRequest) (Student, error) {
	// Mengubah format data dari YYYY-MM-DD enjaf
	t, _ := time.Parse("2006-01-02", studentRequest.Birthdate)
	rfc3339Date, _ := time.Parse(time.RFC3339, t.Format(time.RFC3339))

	student := Student{
		Name:      studentRequest.Name,
		Address:   studentRequest.Address,
		Major:     studentRequest.Major,
		Phone:     studentRequest.Phone,
		Email:     studentRequest.Email,
		Birthdate: rfc3339Date,
	}

	newStudent, err := s.repository.Create(student)
	return newStudent, err
}

func (s *service) Update(ID int, studentRequest StudentRequest) (Student, error) {

	// Validasi ID
	if ID <= 0 {
		return Student{}, errors.New("ID mahasiswa tidak valid")
	}

	// Validasi data mahasiswa
	v := validator.New()
	if e := v.Struct(studentRequest); e != nil {
		return Student{}, e
	}

	// Temukan mahasiswa berdasarkan ID
	student, e := s.repository.FindByID(ID)
	if e != nil {
		return Student{}, e
	}

	if studentRequest.Name != "" {
		student.Name = studentRequest.Name
	}
	if studentRequest.Address != "" {
		student.Address = studentRequest.Address
	}
	if studentRequest.Major != "" {
		student.Major = studentRequest.Major
	}
	if studentRequest.Phone != "" {
		student.Phone = studentRequest.Phone
	}
	if studentRequest.Email != "" {
		student.Email = studentRequest.Email
	}
	if studentRequest.Birthdate != "" {
		// Mengubah format data dari YYYY-MM-DD enjaf
		t, _ := time.Parse("2006-01-02", studentRequest.Birthdate)
		rfc3339Date, _ := time.Parse(time.RFC3339, t.Format(time.RFC3339))
		student.Birthdate = rfc3339Date
	}

	newStudent, err := s.repository.Update(student)
	return newStudent, err
}

func (s *service) Delete(ID int) (Student, error) {
	student, err := s.repository.FindByID(ID)
	_, err = s.repository.Delete(student)

	return student, err
}

func (s *service) AddCourseToStudent(studentID int, courseID int) (Student, error) {
	student, err := s.repository.AddCourseToStudent(studentID, courseID)

	return student, err
}
