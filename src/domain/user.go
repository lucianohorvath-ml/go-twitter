package domain

type User struct {
	Nombre     string `json:"name"`
	Mail       string `json:"mail"`
	Nick       string `json:"nick"`
	Contraseña string `json:"password"`
}

func NewUser(nombre, mail, nick, contraseña string) *User {
	user := User{nombre, mail, nick, contraseña}
	return &user
}
