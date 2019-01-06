package channel_test

import (
	cha "final_project2/channel"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestChannels(t *testing.T) {

	RegisterFailHandler(Fail)
	RunSpecs(t, "Channel Suite")
}

var _ = Describe("Channel", func() {
	Describe("Check channel's functionality", func() {
		Context("Create channel", func() {
			It("Create successed", func() {
				channel := make(chan int, 1)
				testChannel, err := cha.New(1, channel)
				Expect(err).To(BeNil())
				Expect(testChannel.GetChannel()).ToNot(BeNil())
			})
			It("Create with bad probability", func() {
				channel := make(chan int, 1)
				_, err := cha.New(1.5, channel)
				Expect(err).ToNot(BeNil())
				Expect(err.Error()).To(Equal(cha.ChannelsProbNotGoodErrMsg))
			})
		})
		Context("Insert number to channel", func() {
			It("Insert number", func() {
				channel := make(chan int, 1)
				testChannel, _ := cha.New(1, channel)
				testChannel.InsertNumber(5)
				Expect(<-testChannel.GetChannel()).To(Equal(5))
			})
		})
		Context("Get sum from channel", func() {
			It("Get sum", func() {
				channel := make(chan int, 3)
				testChannel, _ := cha.New(1, channel)
				testChannel.InsertNumber(1)
				testChannel.InsertNumber(2)
				testChannel.InsertNumber(3)
				Expect(testChannel.GetSum()).To(Equal(6))
			})
		})
	})
})
