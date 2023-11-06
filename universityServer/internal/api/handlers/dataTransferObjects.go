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
