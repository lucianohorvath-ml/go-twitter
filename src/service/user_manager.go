package service

import "github.com/lucianohorvath-ml/go-twitter/src/domain"

var RegisteredUsers []*domain.User

func RegisterUser(user *domain.User) {
	RegisteredUsers = append(RegisteredUsers, user)
}

func IsRegistered(user *domain.User) bool {
	for _, r_user := range RegisteredUsers {
		if user.Nombre == r_user.Nombre {
			return true
		}
	}
	return false
}
