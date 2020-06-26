package message_test

import (
	"fmt"
	. "github.com/myugen/kata-boilerplate/internal/message"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"strings"
)

var _ = Describe("Tweet", func() {
	const (
		ShortTweet = "Be careful with 5G, if Bill Gates wants to control us then we need to assimiliate BSOD in our brains"
		LongTweet = "Be careful with 5G, if Bill Gates wants to control us then we need to assimiliate BSOD in our brains. So imagine, how we can press \"Enter\" key if we don't have any built-in keyboard... nah, forgetting press \"Enter\" never works, really anyone saw anything when pressed, it was a trap!!!"
	)

	var tweet *Tweet

	Context("writing a new Tweet", func() {
		BeforeEach(func() {
			tweet = NewTweet(ShortTweet)
		})

		It("should has the written text", func() {
			Expect(tweet.Text).To(Equal(ShortTweet))
		})
	})

	Context("having written one", func() {
		When("not exceeding limit character size", func() {
			BeforeEach(func() {
				tweet = NewTweet(ShortTweet)
			})

			It("should be sent", func() {
				result := tweet.Send("everyone")
				Expect(strings.Contains(result, fmt.Sprintf("Tweet: %s", tweet.Text))).To(BeTrue())
				Expect(strings.Contains(result, "To: everyone")).To(BeTrue())
				Expect(strings.Contains(result, "Status: SENT")).To(BeTrue())
			})
		})

		When("exceeding limit character size", func() {
			BeforeEach(func() {
				tweet = NewTweet(LongTweet)
			})

			It("should not be sent", func() {
				result := tweet.Send("everyone")
				Expect(strings.Contains(result, "Status: NOT SENT")).To(BeTrue())
			})
		})
	})
})
