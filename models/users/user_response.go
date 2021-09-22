package users

type Response struct {
	Id       uint   `json:"id"`
	SalaryId int    `json:"salary_id"`
	Name     string `json:"name"`
	IsAdmin  int    `json:"is_admin" validate:"numeric"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Gender   string `json:"gender"`
	Birthday string `json:"birthday"`
	Address  string `json:"address"`
	Company  string `json:"company"`
	IsValid  int    `json:"is_valid"`
}
