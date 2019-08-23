package service_test

import (
	"github.com/lucianohorvath-ml/go-twitter/src/domain"
	"github.com/lucianohorvath-ml/go-twitter/src/service"
	"testing"
)

func TestPublishedTweetIsSavedToExternalResource(t *testing.T) {
	// Initialization
	var tweetWriter service.TweetWriter
	tweetWriter = service.NewMemoryTweetWriter() // Mock implementation
	tweetManager := service.NewTweetManager(tweetWriter)

	var tweet domain.Tweet
	user := domain.NewUser("a", "a", "a", "1")
	tweetManager.UserManager.RegisterUser(user)
	tweet = domain.NewTextTweet(user, "holaaa")
	// Operation
	id, _ := tweetManager.PublishTweet(tweet)

	// Validation
	// se castea a la implementación xq tweetWriter es una interfaz
	// y no hay acceso a los métodos específicos de la impl.
	memoryWriter := (tweetWriter).(*service.MemoryTweetWriter)
	savedTweet := memoryWriter.GetLastSavedTweet()

	if savedTweet == nil {
		t.Errorf("No hay ningún tweet guardado.")
	}
	if savedTweet.GetId() != id {
		t.Errorf("Se obtuvo id %v y se esperaba: %v", savedTweet.GetId(), id)
	}
}
