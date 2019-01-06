package player_test

import (
	p "final_project2/player"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const (
	TestUsername     = "testUser"
	TestRandomNumber = 15
)

var (
	channelPlayer chan int
	channels      []chan int
)

func TestPlayers(t *testing.T) {
	channelPlayer, channels = createPlayersChannels()

	RegisterFailHandler(Fail)
	RunSpecs(t, "Player Suite")
}

var _ = Describe("Player", func() {
	Describe("Check player's functionality", func() {
		Context("Create player", func() {
			It("creae successed", func() {
				testPlayer, err := p.New(TestUsername, TestRandomNumber, channelPlayer, channels)
				Expect(err).To(BeNil())
				Expect(testPlayer.GetUsername()).To(Equal(TestUsername))
				Expect(testPlayer.GetRandomNumber()).To(Equal(TestRandomNumber))
				Expect(testPlayer.GetChannel()).ToNot(BeNil())
				channelsList := testPlayer.GetotherPlayersChannels()
				Expect(channelsList).ToNot(BeNil())
				Expect(len(channelsList)).To(Equal(2))
			})
			It("creae with no channels list", func() {
				_, err := p.New(TestUsername, TestRandomNumber, channelPlayer, nil)
				Expect(err).ToNot(BeNil())
				Expect(err.Error()).To(Equal(p.ChannelsListNotGoodErrMsg))
			})
			It("creae with empty channels list", func() {
				var channels []chan int
				_, err := p.New(TestUsername, TestRandomNumber, channelPlayer, channels)
				Expect(err).ToNot(BeNil())
				Expect(err.Error()).To(Equal(p.ChannelsListNotGoodErrMsg))
			})
		})
	})
})

func createPlayersChannels() (chan int, []chan int) {
	//Create player channel
	channelPlayer := make(chan int)
	//Create othe players channels
	channel2 := make(chan int)
	channel3 := make(chan int)
	//Create a list of channels and add channels to the list
	var channels []chan int
	channels = append(channels, channel2)
	channels = append(channels, channel3)

	return channelPlayer, channels
}
