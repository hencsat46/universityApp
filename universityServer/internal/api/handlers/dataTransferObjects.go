package handlers

type GetUniversityDTO struct {
	Order int `json:"order"`
}

type SignUpDTO struct {
	StudentName    string `json:"StudentName"`
	StudentSurname string `json:"StudentSurname"`
	Username       string `json:"Username"`
	Password       string `json:"Password"`
}

type SignInDTO struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
}
