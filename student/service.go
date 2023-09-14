package student

type Service interface {
	FindAll(UserID uint) ([]Student, error)
	FindByID(ID int) (Student, error)
	Create(studentRequest StudentRequest,UserID uint) (Student, error)
	Update(ID int, StudentRequest StudentRequest) (Student, error)
	Delete(ID int) (Student, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll(UserID uint) ([]Student, error) {
	students, err := s.repository.FindAll()
	return students, err
}

func (s *service) FindByID(ID int) (Student, error) {
	student, err := s.repository.FindByID(ID)
	return student, err
}

func (s *service) Create(studentRequest StudentRequest,UserID uint) (Student, error) {
	student := Student{
		Name:    studentRequest.Name,
		Address: studentRequest.Address,
		Major:   studentRequest.Major,
		Phone:   studentRequest.Phone,
		UserID:  UserID,
	}

	newStudent, err := s.repository.Create(student)
	return newStudent, err
}

func (s *service) Update(ID int, studentRequest StudentRequest) (Student, error) {
	student, err := s.repository.FindByID(ID)

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

	newStudent, err := s.repository.Update(student)
	return newStudent, err
}

func (s *service) Delete(ID int) (Student, error) {
	student, err := s.repository.FindByID(ID)
	_, err = s.repository.Delete(student)

	return student, err
}
