package domain

type User struct {
	Nombre string
	Mail string
	Nick string
	Contraseña string
}

func NewUser(nombre, mail, nick, contraseña string) *User {
	user := User{nombre, mail, nick, contraseña}
	return &user
}
