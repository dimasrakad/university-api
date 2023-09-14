package student

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Student, error)
	// FindAllByUser(UserID uint) ([]Student, error)
	FindByID(ID int) (Student, error)
	Create(student Student) (Student, error)
	Update(student Student) (Student, error)
	Delete(student Student) (Student, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Student, error) {
	var students []Student

	err := r.db.Find(&students).Error

	return students, err
}

// func (r *repository) FindAllByUser(UserID uint) ([]Student, error) {
// 	var students []Student

// 	err := r.db.Where("user_id = ?", UserID).Find(&students).Error

// 	return students, err
// }

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
