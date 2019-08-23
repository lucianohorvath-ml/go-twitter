package domain_test

import (
	"github.com/lucianohorvath-ml/go-twitter/src/domain"
	"testing"
)

func TestTextTweetPrintsUserAndText(t *testing.T) {
	// Initialization
	user := domain.NewUser("grupoesfera", "ge", "ge", "1234")
	tweet := domain.NewTextTweet(user, "This is my tweet")

	// Operation
	text := tweet.PrintTweet()

	// Validation
	expectedText := "@grupoesfera: This is my tweet"
	if text != expectedText {
		t.Errorf("The expected text is %s but was %s", expectedText, text)
	}
}

func TestImageTweetPrintsUserTextAndImageURL(t *testing.T) {
	// Initialization
	user := domain.NewUser("esfera", "e", "e", "1")
	textTweet := domain.NewTextTweet(user, "This is my image")
	tweet := domain.NewImageTweet(textTweet,
		"http://www.grupoesfera.com.ar/common/img/grupoesfera.png")

	// Operation
	text := tweet.PrintTweet()

	// Validation
	expectedText := "@grupoesfera: This is my image " +
		"http://www.grupoesfera.com.ar/common/img/grupoesfera.png"
	if text != expectedText {
		t.Errorf("Expected text is: %s", expectedText)
	}
}

func TestQuoteTweetPrintsUserTextAndQuotedTweet(t *testing.T) {
	// Initialization
	user := domain.NewUser("nick", "e", "e", "1")
	user2 := domain.NewUser("lucas", "a", "e", "1")
	textTweet := domain.NewTextTweet(user, "estoy citando un tweet")
	quotedTweet := domain.NewTextTweet(user2, "hola chicos")
	tweet := domain.NewQuoteTweet(textTweet, quotedTweet)

	// Operation
	text := tweet.PrintTweet()

	// Validation
	expectedText := `@nick: estoy citando un tweet "@lucas: hola chicos"`
	if text != expectedText {
		t.Errorf("Expected text is: %s", expectedText)
	}
}
