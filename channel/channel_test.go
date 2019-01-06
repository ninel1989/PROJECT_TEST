package channel_test

import (
	//cha "final_project2/channel"
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
			/*It("creae successed", func() {
				testPlayer, err := p.New(TestUsername, TestRandomNumber, channelPlayer, channels)
				Expect(err).To(BeNil())
				Expect(testPlayer.GetUsername()).To(Equal(TestUsername))
				Expect(testPlayer.GetRandomNumber()).To(Equal(TestRandomNumber))
				Expect(testPlayer.GetChannel()).ToNot(BeNil())
				channelsList := testPlayer.GetotherPlayersChannels()
				Expect(channelsList).ToNot(BeNil())
				Expect(len(channelsList)).To(Equal(2))
			})
			It("creae with bad probability", func() {
				_, err := p.New(TestUsername, TestRandomNumber, channelPlayer, nil)
				Expect(err).ToNot(BeNil())
				Expect(err.Error()).To(Equal(p.ChannelsListNotGoodErrMsg))
			})*/
		})
	})
})
