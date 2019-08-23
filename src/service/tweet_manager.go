package service

import (
	"fmt"
	"github.com/lucianohorvath-ml/go-twitter/src/domain"
	"strings"
)

type TweetManager struct {
	tweets       []domain.Tweet
	tweetsByUser map[string][]domain.Tweet
	lastId       int
	UserManager  *UserManager
	writer       TweetWriter
}

func NewTweetManager(writer TweetWriter) *TweetManager {
	tweetManager := new(TweetManager)
	tweetManager.tweets = make([]domain.Tweet, 0)
	tweetManager.tweetsByUser = make(map[string][]domain.Tweet)
	tweetManager.lastId = 0
	tweetManager.UserManager = NewUserManager()
	tweetManager.writer = writer

	return tweetManager
}

func (tm *TweetManager) PublishTweet(tweet domain.Tweet) (int, error) {
	// En Go, se estila hacer el return al detectar el error, para cortar el flujo, en vez de
	// declarar error arriba y hacer el return abajo de todo.

	if tweet.GetUser().Nombre == "" {
		return 0, fmt.Errorf("user is required")
	} else if tweet.GetText() == "" {
		return 0, fmt.Errorf("text is required")
	} else if len(tweet.GetText()) > 140 {
		return 0, fmt.Errorf("text can not exceed 140 characters")
	} else if !tm.UserManager.IsRegistered(tweet.GetUser()) {
		return 0, fmt.Errorf("user must be registered")
	} else {
		tweet.SetId(tm.lastId + 1)
		tm.tweets = append(tm.tweets, tweet)
		tm.tweetsByUser[tweet.GetUser().Nombre] = append(tm.tweetsByUser[tweet.GetUser().Nombre], tweet)
		tm.lastId++
		tm.writer.Write(tweet)
	}
	return tweet.GetId(), nil
}

func (tm *TweetManager) GetTweets() []domain.Tweet {
	return tm.tweets
}

func (tm *TweetManager) GetTweet() domain.Tweet {
	return tm.tweets[0]
}

func (tm *TweetManager) GetTweetById(id int) domain.Tweet {
	for i := 0; i < len(tm.tweets); i++ {
		if tm.tweets[i].GetId() == id {
			return tm.tweets[i]
		}
	}
	return nil
}

func (tm *TweetManager) CountTweetsByUser(username string) int {
	return len(tm.tweetsByUser[username])
}

func (tm *TweetManager) GetTweetsByUser(username string) []domain.Tweet {
	return tm.tweetsByUser[username]
}

func (tm *TweetManager) SearchTweetsContaining(search string, tweetsChannel chan domain.Tweet) {
	// s es el texto de búsqueda
	// tweets es un canal cuyos mensajes son estructuras que implementan la interfaz Tweet
	go func() {
		for _, tweet := range tm.tweets {
			if strings.Contains(tweet.GetText(), search) {
				tweetsChannel <- tweet
			}
		}
		close(tweetsChannel)
	}()
}
