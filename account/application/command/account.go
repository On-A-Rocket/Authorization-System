package command

type CreateAccountCommand struct {
	Id          string
	Password    string
	Name        string
	Email       string
	PhoneNumber string
	WorkCode    int
	HireDate    string
}
