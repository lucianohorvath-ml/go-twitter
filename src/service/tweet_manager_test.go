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

// TestTweetWithoutUserIsNotPublished verifica que no se pueda twittear sin especificar usuario.
func TestTweetWithoutUserIsNotPublished(t *testing.T) {
	// Initialization
	var tweet *domain.Tweet

	var user string
	text := "This is my first tweet"
	tweet = domain.NewTweet(user, text)

	// Operation
	var err error
	err = service.PublishTweet(tweet)

	// Validation
	if err == nil {
		t.Error("Expected error did not appear")
	}

	if err != nil && err.Error() != "user is required" {
		t.Error("Expected error is user is required")
	}
}

func TestTweetWithoutTextIsNotPublished(t *testing.T) {
	var tweet *domain.Tweet

	var text string
	user := "Luciano"
	tweet = domain.NewTweet(user, text)

	// Operation
	var err error
	err = service.PublishTweet(tweet)

	// Validation
	if err == nil {
		t.Error("Expected error did not appear")
	}

	if err != nil && err.Error() != "text is required" {
		t.Error("Expected error is text is required")
	}
}

func TestTweetWhichExceeding140CharactersIsNotPublished(t *testing.T) {
	var tweet *domain.Tweet

	user := "Luciano"
	text := "Lorem ipsum dolor sit amet consectetur adipiscing elit donec, " +
		"risus natoque diam mauris felis maecenas placerat turpis luctus, " +
		"porttitor nam magna sa."
		tweet = domain.NewTweet(user, text)

	// Operation
	var err error
	err = service.PublishTweet(tweet)

	// Validation
	if err == nil {
		t.Error("Expected error did not appear")
	}

	if err != nil && err.Error() != "text can not exceed 140 characters" {
		t.Error("Expected error is text can not exceed 140 characters")
	}
}

func TestCanPublishAndRetrieveMoreThanOneTweet(t *testing.T) {
	// Initialization
	service.InitializeService()
	var tweet, secondTweet *domain.Tweet // Fill the tweets with data

	// Operation
	service.PublishTweet(tweet)
	service.PublishTweet(secondTweet)

	// Validation
	publishedTweets := service.GetTweets()
	if len(publishedTweets) != 2 {
		t.Errorf("Expected size is 2 but was %d", len(publishedTweets))
		return
	}
	firstPublishedTweet := publishedTweets[0]
	secondPublishedTweet := publishedTweets[1]
	if !isValidTweet(t, firstPublishedTweet, user, text) {
		return
	}
	// Same for secondPublishedTweet
}