package service

import "github.com/lucianohorvath-ml/go-twitter/src/domain"

//var Tweet string

// PublishTweet guarda un tweet que recibe por parámetro.
//func PublishTweet(tweet string) {
//	Tweet = tweet
//}
//
//func GetTweet() string {
//	return Tweet
//}

var Tweet *domain.Tweet

func PublishTweet(tweet *domain.Tweet) {
	Tweet = tweet
}

func GetTweet() *domain.Tweet {
	return Tweet
}
