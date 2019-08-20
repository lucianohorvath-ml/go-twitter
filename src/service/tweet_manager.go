package service

var Tweet string

// PublishTweet guarda un tweet que recibe por parámetro.
func PublishTweet(tweet string) {
	Tweet = tweet
}

func GetTweet() string {
	return Tweet
}
