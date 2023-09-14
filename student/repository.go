package student

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Student, error)
	FindByID(ID int) (Student, error)
	Create(student Student) (Student, error)
	Update(student Student) (Student, error)
	Delete(student Student) (Student, error)
	// AddCourseToStudent(studentID int, courseID int) (Student, error)
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

	err := r.db.First(&student, ID).Error

	return student, err
}

func (r *repository) Create(student Student) (Student, error) {
	err := r.db.Create(&student).Error

	return student, err
}

func (r *repository) Update(student Student) (Student, error) {
	err := r.db.Save(&student).Error

	return student, err
}

func (r *repository) Delete(student Student) (Student, error) {
	err := r.db.Delete(&student).Error

	return student, err
}

// func (r *repository) AddCourseToStudent(studentID int, courseID int) (Student, error) {
// 	// Ambil data siswa dari database
//     student, err := GetStudentByID(studentID)
//     if err != nil {
//         return err
//     }

//     // Ambil data kursus dari database
//     course, err := GetCourseByID(courseID)
//     if err != nil {
//         return err
//     }

//     // Tambahkan ID kursus ke daftar kursus siswa
//     student.Courses = append(student.Courses, courseID)

//     // Simpan perubahan ke dalam database
//     if err := UpdateStudent(student); err != nil {
//         return err
//     }

//     return nil
// }
