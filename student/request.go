package student

type StudentRequest struct {
	Name      string `binding:"required,max=30"`
	Address   string `binding:"required,max=50"`
	Major     string `binding:"required"`
	Phone     string `binding:"required,numeric"`
	Email     string `binding:"required,email"`
	Birthdate string `binding:"required"`
}
