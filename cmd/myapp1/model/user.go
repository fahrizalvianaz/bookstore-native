package model

type User struct {
	ID       int    `json:"id_user"`
	Nama     string `json:"nama"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}
