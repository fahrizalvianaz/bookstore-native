package model

type User struct {
	ID       int    `json:"id_user"`
	Username string `json:"username"`
	Password string `json:"password"`
}
