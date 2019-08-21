package domain

import "time"

type Tweet struct {
	User *User
	Text string
	Date *time.Time // es un puntero a la fecha
	Id   int
}

func NewTweet(user *User, text string) *Tweet {
	now := time.Now()

	tweet := Tweet{user, text, &now, 0}

	return &tweet
}
