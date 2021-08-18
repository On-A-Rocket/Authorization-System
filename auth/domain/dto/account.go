package dto

type CreateAccount struct {
	Id          string `json:"id" example:"simson"`
	Password    string `json:"password" example:"1234"`
	Name        string `json:"name" example:"홍길동"`
	Email       string `json:"email" example:"abc@sample.com"`
	PhoneNumber string `json:"phone_number" example:"010-1234-1234"`
	WorkCode    int    `json:"work_code" example:"1"`
	HireDate    string `json:"hire_date" example:"2021-08-17 09:00:00"`
}
