package service

import "github.com/lucianohorvath-ml/go-twitter/src/domain"

type UserManager struct {
	registeredUsers []*domain.User
}

func NewUserManager() *UserManager {
	manager := new(UserManager)
	manager.registeredUsers = make([]*domain.User, 0)

	return manager
}

func (userManager *UserManager) RegisterUser(user *domain.User) {
	userManager.registeredUsers = append(userManager.registeredUsers, user)
}

func (userManager *UserManager) IsRegistered(user *domain.User) bool {
	for _, r_user := range userManager.registeredUsers {
		if user.Nombre == r_user.Nombre {
			return true
		}
	}
	return false
}
