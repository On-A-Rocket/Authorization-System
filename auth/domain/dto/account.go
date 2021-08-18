package dto

type CreateAccount struct {
	Id          string `json:"id"`
	Password    string `json:"password"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	WorkCode    int    `json:"work_code"`
	HireDate    string `json:"hire_date"`
}
