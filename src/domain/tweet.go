package domain

import (
	"time"
)

type Tweet interface {
	PrintTweet() string
	GetUser() *User
	GetText() string
	GetId() int
	SetId(int)
	GetDate() *time.Time
}

func NewTextTweet(user *User, text string) *TextTweet {
	now := time.Now()
	tweet := TextTweet{user, text, &now, 0}
	return &tweet
}

func NewImageTweet(textTweet *TextTweet, URL string) *ImageTweet {
	imageTweet := ImageTweet{*textTweet, URL}
	return &imageTweet
}

func NewQuoteTweet(textTweet *TextTweet, tweet Tweet) *QuoteTweet {
	quoteTweet := QuoteTweet{
		TextTweet:   *textTweet,
		tweetCitado: tweet,
	}
	return &quoteTweet
}

// sin interfaz
//func (tweet TextTweet) PrintableTweet() string {
//	return fmt.Sprintf("@%s: %s", tweet.User.Nombre, tweet.Text)
//}

// No hace falta definir la interfaz acá, es idéntica a la que define fmt
//type Stringer interface {
//	String() string
//}

// con interfaz. quiera o no, por implementar el metodo ya se implementa la interfaz.
//func (tweet TextTweet) String() string {
//	return fmt.Sprintf("@%s: %s", tweet.User.Nombre, tweet.Text)
//}
