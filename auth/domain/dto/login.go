package dto

type Login struct {
	Id       string `json:"id" example:"simson"`
	Password string `json:"password" example:"1234"`
}

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
