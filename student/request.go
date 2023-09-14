package student

type StudentRequest struct {
	Name    string `binding:"required,max=30"`
	Address string `binding:"required,max=50"`
	Major   string `binding:"required"`
	Phone   string `binding:"required,numeric"`
}
