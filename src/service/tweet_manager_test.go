package service_test

import (
	"github.com/lucianohorvath-ml/go-twitter/src/domain"
	"github.com/lucianohorvath-ml/go-twitter/src/service"
	"testing"
)

// usando struct
func TestPublishedTweetIsSavedWithStruct(t *testing.T) {
	// Initialization
	service.InitializeService()
	var tweet *domain.Tweet
	user := "grupoesfera"
	text := "This is my first tweet"
	tweet = domain.NewTweet(user, text)

	// Operation
	_, _ = service.PublishTweet(tweet)

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
	_, err = service.PublishTweet(tweet)

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
	_, err = service.PublishTweet(tweet)

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
	_, err = service.PublishTweet(tweet)

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
	var tweet, secondTweet *domain.Tweet
	user := "Luciano"
	text := "Hola!"
	user2 := "Marcos"
	text2 := "Chau"
	tweet = domain.NewTweet(user, text)
	secondTweet = domain.NewTweet(user2, text2)
	var id, id2 int

	id, _ = service.PublishTweet(tweet)
	id2, _ = service.PublishTweet(secondTweet)

	publishedTweets := service.GetTweets()
	if len(publishedTweets) != 2 {
		t.Errorf("Expected size is 2 but was %d", len(publishedTweets))
		return
	}
	firstPublishedTweet := publishedTweets[0]
	secondPublishedTweet := publishedTweets[1]
	if !isValidTweet(t, firstPublishedTweet, id, user, text) {
		return
	}
	if !isValidTweet(t, secondPublishedTweet, id2, user2, text2) {
		return
	}
}

func TestCanRetrieveTweetById(t *testing.T) {
	service.InitializeService()

	var tweet *domain.Tweet
	var id int

	user := "grupoesfera"
	text := "This is my first tweet"
	tweet = domain.NewTweet(user, text)
	id, _ = service.PublishTweet(tweet)

	publishedTweet := service.GetTweetById(id)

	isValidTweet(t, publishedTweet, id, user, text)
}

func isValidTweet(t *testing.T, tweet *domain.Tweet, id int, user string, text string) bool {
	if !(tweet.Id == id && tweet.User == user && tweet.Text == text) {
		t.Error("El tweet no es válido.")
		return false
	}
	return true
}

func TestCanCountTheTweetsSentByAnUser(t *testing.T) {
	// Initialization
	service.InitializeService()
	var tweet, secondTweet, thirdTweet *domain.Tweet
	user := "grupoesfera"
	anotherUser := "nick"
	text := "This is my first tweet"
	secondText := "This is my second tweet"
	tweet = domain.NewTweet(user, text)
	secondTweet = domain.NewTweet(user, secondText)
	thirdTweet = domain.NewTweet(anotherUser, text)
	_, _ = service.PublishTweet(tweet)
	_, _ = service.PublishTweet(secondTweet)
	_, _ = service.PublishTweet(thirdTweet)
	// Operation
	count := service.CountTweetsByUser(user)
	// Validation
	if count != 2 {
		t.Errorf("Expected count is 2 but was %d", count)
	}
}

func TestCanRetrieveTheTweetsSentByAnUser(t *testing.T) {
	service.InitializeService()
	var tweet, secondTweet, thirdTweet *domain.Tweet
	var id1, id2 int
	user := "grupoesfera"
	anotherUser := "nick"
	text := "This is my first tweet"
	secondText := "This is my second tweet"
	tweet = domain.NewTweet(user, text)
	secondTweet = domain.NewTweet(user, secondText)
	thirdTweet = domain.NewTweet(anotherUser, text)
	id1, _ = service.PublishTweet(tweet)
	id2, _ = service.PublishTweet(secondTweet)
	_, _ = service.PublishTweet(thirdTweet)

	// Operation
	tweets := service.GetTweetsByUser(user)

	// Validation
	if count := len(tweets); count != 2 {
		t.Errorf("Expected count is 2 but was %d", count)
	}
	firstPublishedTweet := tweets[0]
	secondPublishedTweet := tweets[1]
	isValidTweet(t, firstPublishedTweet, id1, user, text)
	isValidTweet(t, secondPublishedTweet, id2, user, secondText)
}
