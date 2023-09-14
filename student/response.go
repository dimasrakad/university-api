package student

type StudentResponse struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Major   string `json:"major"`
	Phone   string `json:"phone"`
}
