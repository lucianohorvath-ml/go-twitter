package service_test

import (
	"github.com/lucianohorvath-ml/go-twitter/src/domain"
	"github.com/lucianohorvath-ml/go-twitter/src/service"
	"testing"
)

func TestPublishedTweetIsSaved(t *testing.T) {
	// Initialization. todo: separar todas las inicializaciones en un setup
	writer := service.NewFileTweetWriter()
	manager := service.NewTweetManager(writer)
	var tweet domain.Tweet
	username := "grupoesfera"
	user := domain.NewUser(username, "aa", "esfera", "1234")
	manager.UserManager.RegisterUser(user)
	text := "This is my first tweet"
	tweet = domain.NewTextTweet(user, text)

	// Operation
	_, _ = manager.PublishTweet(tweet)

	// Validation
	publishedTweet := manager.GetTweet()
	if publishedTweet.GetUser() != user &&
		publishedTweet.GetText() != text {
		t.Errorf("Expected tweet is %s: %s \nbut is %s: %s",
			user, text, publishedTweet.GetUser(), publishedTweet.GetText())
	}
	if publishedTweet.GetDate() == nil {
		t.Error("Expected date can't be nil")
	}
}

// TestTweetWithoutUserIsNotPublished verifica que no se pueda twittear sin especificar usuario.
func TestTweetWithoutUserIsNotPublished(t *testing.T) {
	// Initialization
	userManager := service.NewUserManager()
	var tweet domain.Tweet

	user := domain.NewUser("", "pepe@asd.com", "a", "1234")
	userManager.RegisterUser(user)
	text := "This is my first tweet"
	tweet = domain.NewTextTweet(user, text)

	// Operation
	var err error
	writer := service.NewFileTweetWriter()
	manager := service.NewTweetManager(writer)
	_, err = manager.PublishTweet(tweet)

	// Validation
	if err == nil {
		t.Error("Expected error did not appear")
	}

	if err != nil && err.Error() != "user is required" {
		t.Error("Expected error is user is required")
	}
}

func TestTweetWithoutTextIsNotPublished(t *testing.T) {
	var tweet domain.Tweet

	var text string
	user := domain.NewUser("Luciano", "a", "a", "1234")
	tweet = domain.NewTextTweet(user, text)
	writer := service.NewMemoryTweetWriter()

	// Operation
	var err error
	manager := service.NewTweetManager(writer)
	_, err = manager.PublishTweet(tweet)

	// Validation
	if err == nil {
		t.Error("Expected error did not appear")
	}

	if err != nil && err.Error() != "text is required" {
		t.Error("Expected error is text is required")
	}
}

func TestTweetWhichExceeding140CharactersIsNotPublished(t *testing.T) {
	var tweet domain.Tweet

	user := domain.NewUser("Luciano", "a", "a", "1234")
	text := "Lorem ipsum dolor sit amet consectetur adipiscing elit donec, " +
		"risus natoque diam mauris felis maecenas placerat turpis luctus, " +
		"porttitor nam magna sa."
	tweet = domain.NewTextTweet(user, text)
	writer := service.NewMemoryTweetWriter()

	// Operation
	var err error
	manager := service.NewTweetManager(writer)
	_, err = manager.PublishTweet(tweet)

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
	writer := service.NewFileTweetWriter()
	manager := service.NewTweetManager(writer)
	var tweet, secondTweet domain.Tweet
	user := domain.NewUser("Luciano", "l", "lucho", "1234")
	user2 := domain.NewUser("Marcos", "m", "marquitos", "1234")
	manager.UserManager.RegisterUser(user)
	manager.UserManager.RegisterUser(user2)
	text := "Hola!"
	text2 := "Chau"
	tweet = domain.NewTextTweet(user, text)
	secondTweet = domain.NewTextTweet(user2, text2)
	var id, id2 int

	id, _ = manager.PublishTweet(tweet)
	id2, _ = manager.PublishTweet(secondTweet)

	publishedTweets := manager.GetTweets()
	if len(publishedTweets) != 2 {
		t.Errorf("Expected size is 2 but was %d", len(publishedTweets))
		return
	}
	firstPublishedTweet := publishedTweets[0]
	secondPublishedTweet := publishedTweets[1]
	if !isValidTweet(t, firstPublishedTweet, id, user.Nombre, text) {
		return
	}
	if !isValidTweet(t, secondPublishedTweet, id2, user2.Nombre, text2) {
		return
	}
}

func TestCanRetrieveTweetById(t *testing.T) {
	writer := service.NewMemoryTweetWriter()
	manager := service.NewTweetManager(writer)

	var tweet domain.Tweet
	var id int

	user := domain.NewUser("grupoesfera", "a", "a", "1234")
	manager.UserManager.RegisterUser(user)
	text := "This is my first tweet"
	tweet = domain.NewTextTweet(user, text)
	id, _ = manager.PublishTweet(tweet)

	publishedTweet := manager.GetTweetById(id)

	isValidTweet(t, publishedTweet, id, user.Nombre, text)
}

func isValidTweet(t *testing.T, tweet domain.Tweet, id int, user string, text string) bool {
	if !(tweet.GetId() == id && tweet.GetUser().Nombre == user && tweet.GetText() == text) {
		t.Error("El tweet no es válido.")
		return false
	}
	return true
}

func TestCanCountTheTweetsSentByAnUser(t *testing.T) {
	// Initialization
	writer := service.NewMemoryTweetWriter()
	manager := service.NewTweetManager(writer)

	var tweet, secondTweet, thirdTweet domain.Tweet
	user := domain.NewUser("grupoesfera", "a", "a", "1234")
	anotherUser := domain.NewUser("nico", "n", "n", "1234")
	manager.UserManager.RegisterUser(user)
	manager.UserManager.RegisterUser(anotherUser)
	text := "This is my first tweet"
	secondText := "This is my second tweet"
	tweet = domain.NewTextTweet(user, text)
	secondTweet = domain.NewTextTweet(user, secondText)
	thirdTweet = domain.NewTextTweet(anotherUser, text)
	_, _ = manager.PublishTweet(tweet)
	_, _ = manager.PublishTweet(secondTweet)
	_, _ = manager.PublishTweet(thirdTweet)
	// Operation
	count := manager.CountTweetsByUser(user.Nombre)
	// Validation
	if count != 2 {
		t.Errorf("Expected count is 2 but was %d", count)
	}
}

func TestCanRetrieveTheTweetsSentByAnUser(t *testing.T) {
	writer := service.NewMemoryTweetWriter()
	manager := service.NewTweetManager(writer)

	var tweet, secondTweet, thirdTweet domain.Tweet
	var id1, id2 int

	user := domain.NewUser("grupoesfera", "a", "a", "1234")
	anotherUser := domain.NewUser("nico", "n", "n", "1234")
	manager.UserManager.RegisterUser(user)
	manager.UserManager.RegisterUser(anotherUser)
	text := "This is my first tweet"
	secondText := "This is my second tweet"
	tweet = domain.NewTextTweet(user, text)
	secondTweet = domain.NewTextTweet(user, secondText)
	thirdTweet = domain.NewTextTweet(anotherUser, text)
	id1, _ = manager.PublishTweet(tweet)
	id2, _ = manager.PublishTweet(secondTweet)
	_, _ = manager.PublishTweet(thirdTweet)

	// Operation
	tweets := manager.GetTweetsByUser(user.Nombre)

	// Validation
	if count := len(tweets); count != 2 {
		t.Errorf("Expected count is 2 but was %d", count)
	}
	firstPublishedTweet := tweets[0]
	secondPublishedTweet := tweets[1]
	isValidTweet(t, firstPublishedTweet, id1, user.Nombre, text)
	isValidTweet(t, secondPublishedTweet, id2, user.Nombre, secondText)
}

// Ejercicios extra
// TestUserCanRegister un usuario puede registrarse en el sistema.
func TestUserCanRegister(t *testing.T) {
	userManager := service.NewUserManager()
	user := domain.NewUser("pepe",
		"pepe@hotmail.com",
		"pepito",
		"1234")
	userManager.RegisterUser(user)
	// todo: es correcto probar esta funcionalidad usando funcion del servicio?
	if !userManager.IsRegistered(user) {
		t.Error("Expected registered user!")
	}
}

// TestUnregisteredUserCanNotTweet un usuario no registrado no puede twittear.
func TestUnregisteredUserCanNotTweet(t *testing.T) {
	user := domain.NewUser("pepe",
		"pepe@hotmail.com",
		"pepito",
		"1234")
	tweet := domain.NewTextTweet(user, "hola mundo!")
	writer := service.NewMemoryTweetWriter()
	manager := service.NewTweetManager(writer)
	_, err := manager.PublishTweet(tweet)

	// Validation
	if err == nil {
		t.Error("Expected error did not appear")
	}

	if err != nil && err.Error() != "user must be registered" {
		t.Error("Expected error is: user must be registered")
	}
}
