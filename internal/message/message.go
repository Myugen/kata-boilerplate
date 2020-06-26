package message

import "fmt"

type Message interface {
	Send(to string, props ...interface{}) string
}

type Tweet struct {
	Text    string
	success string
	failed  string
}

func NewTweet(text string) *Tweet {
	tweet := new(Tweet)
	tweet.Text = text
	tweet.success = "Tweet: %s\nTo: %s\nStatus: SENT"
	tweet.failed = "Status: NOT SENT\nReason: Error occurs trying to send a tweet.\nCause: %s"
	return tweet
}

func (t *Tweet) Send(to string) string {
	if len(t.Text) > 280 {
		return fmt.Sprintf(t.failed, "Tweet exceeds maximum character size, 280. ")
	}
	return fmt.Sprintf(t.success, t.Text, to)
}
