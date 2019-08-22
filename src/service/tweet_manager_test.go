package service_test

import (
	"github.com/lucianohorvath-ml/go-twitter/src/domain"
	"github.com/lucianohorvath-ml/go-twitter/src/service"
	"testing"
)

// usando struct
func TestPublishedTweetIsSaved(t *testing.T) {
	// Initialization
	service.InitializeService()
	var tweet *domain.Tweet
	username := "grupoesfera"
	user := domain.NewUser(username, "aa", "esfera", "1234")
	service.RegisterUser(user)
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

	user := domain.NewUser("", "pepe@asd.com", "a", "1234")
	service.RegisterUser(user)
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
	user := domain.NewUser("Luciano", "a", "a", "1234")
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

	user := domain.NewUser("Luciano", "a", "a", "1234")
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
	user := domain.NewUser("Luciano", "l", "lucho", "1234")
	user2 := domain.NewUser("Marcos", "m", "marquitos", "1234")
	service.RegisterUser(user)
	service.RegisterUser(user2)
	text := "Hola!"
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
	if !isValidTweet(t, firstPublishedTweet, id, user.Nombre, text) {
		return
	}
	if !isValidTweet(t, secondPublishedTweet, id2, user2.Nombre, text2) {
		return
	}
}

func TestCanRetrieveTweetById(t *testing.T) {
	service.InitializeService()

	var tweet *domain.Tweet
	var id int

	user := domain.NewUser("grupoesfera", "a", "a", "1234")
	text := "This is my first tweet"
	tweet = domain.NewTweet(user, text)
	id, _ = service.PublishTweet(tweet)

	publishedTweet := service.GetTweetById(id)

	isValidTweet(t, publishedTweet, id, user.Nombre, text)
}

func isValidTweet(t *testing.T, tweet *domain.Tweet, id int, user string, text string) bool {
	if !(tweet.Id == id && tweet.User.Nombre == user && tweet.Text == text) {
		t.Error("El tweet no es válido.")
		return false
	}
	return true
}

func TestCanCountTheTweetsSentByAnUser(t *testing.T) {
	// Initialization
	service.InitializeService()
	var tweet, secondTweet, thirdTweet *domain.Tweet
	user := domain.NewUser("grupoesfera", "a", "a", "1234")
	anotherUser := domain.NewUser("nico", "n", "n", "1234")
	text := "This is my first tweet"
	secondText := "This is my second tweet"
	tweet = domain.NewTweet(user, text)
	secondTweet = domain.NewTweet(user, secondText)
	thirdTweet = domain.NewTweet(anotherUser, text)
	_, _ = service.PublishTweet(tweet)
	_, _ = service.PublishTweet(secondTweet)
	_, _ = service.PublishTweet(thirdTweet)
	// Operation
	count := service.CountTweetsByUser(user.Nombre)
	// Validation
	if count != 2 {
		t.Errorf("Expected count is 2 but was %d", count)
	}
}

func TestCanRetrieveTheTweetsSentByAnUser(t *testing.T) {
	service.InitializeService()
	var tweet, secondTweet, thirdTweet *domain.Tweet
	var id1, id2 int

	user := domain.NewUser("grupoesfera", "a", "a", "1234")
	anotherUser := domain.NewUser("nico", "n", "n", "1234")
	text := "This is my first tweet"
	secondText := "This is my second tweet"
	tweet = domain.NewTweet(user, text)
	secondTweet = domain.NewTweet(user, secondText)
	thirdTweet = domain.NewTweet(anotherUser, text)
	id1, _ = service.PublishTweet(tweet)
	id2, _ = service.PublishTweet(secondTweet)
	_, _ = service.PublishTweet(thirdTweet)

	// Operation
	tweets := service.GetTweetsByUser(user.Nombre)

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
	user := domain.NewUser("pepe",
		"pepe@hotmail.com",
		"pepito",
		"1234")
	service.RegisterUser(user)
	// todo: es correcto probar esta funcionalidad usando funcion del servicio?
	if !service.IsRegistered(user) {
		t.Error("Expected registered user!")
	}
}

// TestUnregisteredUserCanNotTweet un usuario no registrado no puede twittear.
func TestUnregisteredUserCanNotTweet(t *testing.T) {
	user := domain.NewUser("pepe",
		"pepe@hotmail.com",
		"pepito",
		"1234")
	tweet := domain.NewTweet(user, "hola mundo!")
	_, err := service.PublishTweet(tweet)

	// Validation
	if err == nil {
		t.Error("Expected error did not appear")
	}

	if err != nil && err.Error() != "user must be registered" {
		t.Error("Expected error is: user must be registered")
	}
}
