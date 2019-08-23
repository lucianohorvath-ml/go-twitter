package domain

import (
	"fmt"
	"time"
)

type TextTweet struct {
	User *User
	Text string
	Date *time.Time // es un puntero a la fecha
	Id   int
}

type ImageTweet struct {
	TextTweet
	URL string
}

type QuoteTweet struct {
	TextTweet
	tweetCitado Tweet
}

func (tweet *TextTweet) PrintTweet() string {
	return fmt.Sprintf("@%s: %s", tweet.User.Nombre, tweet.Text)
}

// todo: revisar implementaciones (ocultas por la de texttweet)
func (tweet *ImageTweet) PrintTweet() string {
	return fmt.Sprintf("@%s: %s\nImagen: %s", tweet.User.Nombre, tweet.Text, tweet.URL)
}

func (tweet *QuoteTweet) PrintTweet() string {
	return fmt.Sprintf("%s. %s", tweet.tweetCitado.PrintTweet(), tweet.PrintTweet())
}

func (tweet *TextTweet) GetUser() *User {
	return tweet.User
}

//func (tweet *ImageTweet) GetUser() *User {
//	return tweet.User
//}
//
//func (tweet *QuoteTweet) GetUser() *User {
//	return tweet.User
//}

func (tweet *TextTweet) GetText() string {
	return tweet.Text
}

// Implementados en texttweet. Como imagetweet y quotetweet lo tienen embebido, tienen el metodo.
//func (tweet *ImageTweet) GetText() string {
//	return tweet.Text
//}
//
//func (tweet *QuoteTweet) GetText() string {
//	return tweet.Text
//}

func (tweet *TextTweet) GetId() int {
	return tweet.Id
}

func (tweet *TextTweet) SetId(id int) {
	tweet.Id = id
}

func (tweet *TextTweet) GetDate() *time.Time {
	return tweet.Date
}
