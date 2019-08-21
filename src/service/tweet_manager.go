package service

import (
	"fmt"
	"github.com/lucianohorvath-ml/go-twitter/src/domain"
)

//var Tweet string

// PublishTweet guarda un tweet que recibe por parámetro.
//func PublishTweet(tweet string) {
//	Tweet = tweet
//}
//
//func GetTweet() string {
//	return Tweet
//}
var Tweets []*domain.Tweet

func InitializeService() {

}

func PublishTweet(tweet *domain.Tweet) error {
	// En Go, se estila hacer el return al detectar el error, para cortar el flujo, en vez de
	// declarar error arriba y hacer el return abajo de tod o.
	if tweet.User == "" {
		return fmt.Errorf("user is required")
	} else if tweet.Text == "" {
		return fmt.Errorf("text is required")
	} else if len(tweet.Text) > 140 {
		return fmt.Errorf("text can not exceed 140 characters")
	} else {
		Tweets = append(Tweets, tweet)
	}
	return nil
}

func GetTweet() *domain.Tweet {
	return Tweet
}

func GetTweets() *domain.Tweet {
	return Tweet
}
