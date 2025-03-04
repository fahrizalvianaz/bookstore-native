package dto

type RegisterRequest struct {
	Nama     string `json:"nama"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
