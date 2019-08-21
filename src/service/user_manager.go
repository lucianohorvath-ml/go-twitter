package service

import "github.com/lucianohorvath-ml/go-twitter/src/domain"

var RegisteredUsers []*domain.User

func RegisterUser(user *domain.User){
	RegisteredUsers = append(RegisteredUsers, user)
}
