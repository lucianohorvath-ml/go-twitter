package service

import (
	"github.com/lucianohorvath-ml/go-twitter/src/domain"
	"os"
)

type TweetWriter interface {
	Write(tweet domain.Tweet)
}

type MemoryTweetWriter struct {
	lastTweet domain.Tweet
}

type FileTweetWriter struct {
	file *os.File
}

func NewMemoryTweetWriter() *MemoryTweetWriter {
	writer := new(MemoryTweetWriter)

	return writer
}

func NewFileTweetWriter() *FileTweetWriter {
	file, _ := os.OpenFile(
		"tweets.txt",
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0666,
	)

	writer := new(FileTweetWriter)
	writer.file = file

	return writer
}

func (writer *MemoryTweetWriter) Write(tweet domain.Tweet) {
	writer.lastTweet = tweet
}

func (writer *FileTweetWriter) Write(tweet domain.Tweet) {
	go func() {
		if writer.file != nil {
			byteSlice := []byte(tweet.PrintTweet() + "\n")
			writer.file.Write(byteSlice)
		}
	}()
}

func (writer *MemoryTweetWriter) GetLastSavedTweet() domain.Tweet {
	return writer.lastTweet
}
