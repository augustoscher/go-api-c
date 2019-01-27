package models

//User representa informações de um usuario da api
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
