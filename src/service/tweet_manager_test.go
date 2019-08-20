package service_test

import (
	"github.com/lucianohorvath-ml/go-twitter/src/domain"
	"github.com/lucianohorvath-ml/go-twitter/src/service"
	"testing"
)

//func TestPublishedTweetIsSaved(t *testing.T) {
//	var tweet string = "This is my first tweet"
//
//	service.PublishTweet(tweet)
//
//	if service.GetTweet() != tweet {
//		t.Error("Expected tweet is", tweet)
//	}
//}

// usando struct
func TestPublishedTweetIsSavedWithStruct(t *testing.T) {

	// Initialization
	var tweet *domain.Tweet
	user := "grupoesfera"
	text := "This is my first tweet"
	tweet = domain.NewTweet(user, text)

	// Operation
	service.PublishTweet(tweet)

	// Validation
	publishedTweet := service.GetTweet()
	if publishedTweet.User != user &&
		publishedTweet.Text != text {
		t.Errorf("Expected tweet is %s: %s \nbut is %s: %s",
			user, text, publishedTweet.User, publishedTweet.Text)
	}
	if publishedTweet.Date == nil {
		t.Error("Expected date can't be nil")
	}
}
