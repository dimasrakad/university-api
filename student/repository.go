package student

import (
	"university-api/course"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]Student, error)
	FindByID(ID int) (Student, error)
	Create(student Student) (Student, error)
	Update(student Student) (Student, error)
	Delete(student Student) (Student, error)
	AddCourseToStudent(studentID int, courseID int) (Student, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Student, error) {
	var students []Student

	err := r.db.Preload("Courses").Find(&students).Error

	return students, err
}

func (r *repository) FindByID(ID int) (Student, error) {
	var student Student

	err := r.db.Preload("Courses").First(&student, ID).Error

	return student, err
}

func (r *repository) Create(student Student) (Student, error) {
	err := r.db.Preload("Courses").Create(&student).Error

	return student, err
}

func (r *repository) Update(student Student) (Student, error) {
	err := r.db.Preload("Courses").Save(&student).Error

	return student, err
}

func (r *repository) Delete(student Student) (Student, error) {
	err := r.db.Preload("Courses").Delete(&student).Error

	return student, err
}

func (r *repository) AddCourseToStudent(studentID int, courseID int) (Student, error) {
	// Ambil data siswa dari database
	student, err := r.FindByID(studentID)
	if err != nil {
		return student, err
	}

	// Ambil data mata kuliah dari database
	newCourse, err := course.NewRepository(r.db).FindByID(courseID)
	if err != nil {
		return student, err
	}

	// Tambahkan ID mata kuliah ke daftar mata kuliah siswa
	student.Courses = append(student.Courses, newCourse)

	// Simpan perubahan ke dalam database
	if student, err := r.Update(student); err != nil {
		return student, err
	}

	return student, nil
}
