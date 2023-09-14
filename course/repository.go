package course

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Course, error)
	// FindByUserID(UserID uint) (Course, error)
	FindByID(ID int) (Course, error)
	Create(course Course) (Course, error)
	Update(course Course) (Course, error)
	Delete(course Course) (Course, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Course, error) {
	var course []Course

	err := r.db.Find(&course).Error

	return course, err
}

func (r *repository) FindByUserID(UserID uint) (Course, error) {
	var course Course

	err := r.db.Where("user_id = ?", UserID).First(&course).Error

	return course, err
}

func (r *repository) FindByID(ID int) (Course, error) {
	var course Course

	err := r.db.First(&course, ID).Error

	return course, err
}

func (r *repository) Create(course Course) (Course, error) {
	err := r.db.Create(&course).Error

	return course, err
}

func (r *repository) Update(course Course) (Course, error) {
	err := r.db.Save(&course).Error

	return course, err
}

func (r *repository) Delete(course Course) (Course, error) {
	err := r.db.Delete(&course).Error

	return course, err
}
