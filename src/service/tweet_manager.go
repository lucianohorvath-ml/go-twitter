package service

import (
	"fmt"
	"github.com/lucianohorvath-ml/go-twitter/src/domain"
)

var Tweets []*domain.Tweet
var lastId int = 0

// InitializeService resetea el slice de tweets
func InitializeService() {
	Tweets = nil
	lastId = 0
}

func PublishTweet(tweet *domain.Tweet) (int, error) {
	// En Go, se estila hacer el return al detectar el error, para cortar el flujo, en vez de
	// declarar error arriba y hacer el return abajo de tod o.
	if tweet.User == "" {
		return 0, fmt.Errorf("user is required")
	} else if tweet.Text == "" {
		return 0, fmt.Errorf("text is required")
	} else if len(tweet.Text) > 140 {
		return 0, fmt.Errorf("text can not exceed 140 characters")
	} else {
		tweet.Id = lastId + 1
		Tweets = append(Tweets, tweet)
		lastId++
	}
	return tweet.Id, nil
}

func GetTweets() []*domain.Tweet {
	return Tweets
}

func GetTweet() *domain.Tweet {
	return Tweets[0]
}

func GetTweetById(id int) *domain.Tweet {
	for i := 0; i < len(Tweets); i++ {
		if Tweets[i].Id == id {
			return Tweets[i]
		}
	}
	return nil
}
