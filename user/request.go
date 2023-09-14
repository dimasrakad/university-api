package user

type SignupRequest struct {
	Email    string `binding:"required,email"`
	Password string `binding:"required,min=10,max=20"`
}

type LoginRequest struct {
	Email    string `binding:"required,email"`
	Password string `binding:"required,min=10,max=20"`
}
