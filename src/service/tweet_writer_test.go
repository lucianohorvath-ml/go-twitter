package service_test

import (
	"github.com/lucianohorvath-ml/go-twitter/src/domain"
	"github.com/lucianohorvath-ml/go-twitter/src/service"
	"strings"
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

func TestCanSearchForTweetContainingText(t *testing.T) {
	// Initialization
	var tweetWriter service.TweetWriter
	tweetWriter = service.NewMemoryTweetWriter()
	tweetManager := service.NewTweetManager(tweetWriter)
	user := domain.NewUser("a", "a", "a", "1234")
	tweet := domain.NewTextTweet(user, "first tweet")
	tweetManager.UserManager.RegisterUser(user)
	_, _ = tweetManager.PublishTweet(tweet)

	// Operation
	searchResult := make(chan domain.Tweet)
	query := "first"
	tweetManager.SearchTweetsContaining(query, searchResult)

	// Validation
	foundTweet := <-searchResult

	if foundTweet == nil {
		t.Errorf("No se encontró el tweet.")
	}
	if !strings.Contains(foundTweet.GetText(), query) {
		t.Errorf("Se buscaba %s y el tweet encontrado es %s", query, foundTweet.GetText())
	}
}

// TestCanSearchForTweetContainingTextAndFoundNothing prueba que la goroutine
// no muera si no encuentra ningún match.
func TestCanSearchForTweetContainingTextAndFoundNothing(t *testing.T) {
	// Initialization
	var tweetWriter service.TweetWriter
	tweetWriter = service.NewMemoryTweetWriter()
	tweetManager := service.NewTweetManager(tweetWriter)
	user := domain.NewUser("a", "a", "a", "1234")
	tweet := domain.NewTextTweet(user, "first tweet")
	tweetManager.UserManager.RegisterUser(user)
	_, _ = tweetManager.PublishTweet(tweet)

	// Operation
	searchResult := make(chan domain.Tweet)
	query := "second"
	tweetManager.SearchTweetsContaining(query, searchResult)

	// Esta sentencia se traba hasta que searchResult tenga algo O esté cerrado.
	// En este caso no se encontrará nada y se cerrará.
	foundTweet, opened := <-searchResult

	if foundTweet != nil {
		t.Errorf("El tweet encontrado debe ser nil.")
	}
	if opened {
		t.Errorf("El canal no está cerrado.")
	}
}
