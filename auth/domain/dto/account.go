package dto

type CreateAccount struct {
	Id          string `json:"id" example:"simson"`
	Password    string `json:"password" example:"1234"`
	Name        string `json:"name" example:"홍길동"`
	Email       string `json:"email" example:"abc@sample.com"`
	PhoneNumber string `json:"phone_number" example:"010-1234-1234"`
	WorkCode    int    `json:"work_code" example:"1"`
	HireDate    string `json:"hire_date" example:"2021-08-17"`
}

type Account struct {
	Name           string  `json:"name" example:"홍길동"`
	Email          string  `json:"email" example:"abc@sample.com"`
	PhoneNumber    string  `json:"phone_number" example:"010-1234-1234"`
	DepartmentCode int     `json:"department_code" example:"1"`
	PositionCode   int     `json:"position_code" example:"1"`
	AuthorityCode  int     `json:"authority_code" example:"1"`
	FirstPaymentId string  `json:"first_payment_id" example:"홍환원"`
	FinalPaymentId string  `json:"final_payment_id" example:"오승진"`
	WorkCode       int     `json:"work_code" example:"1"`
	TotalAnnual    float32 `json:"total_annual" example:"12.0"`
	UseAnnual      float32 `json:"use_annual" example:"5.5"`
	RemainAnnual   float32 `json:"remain_annual" example:"6.5"`
	HireDate       string  `json:"hire_date" example:"2021-08-17"`
}
