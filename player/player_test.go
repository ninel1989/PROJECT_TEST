package player_test

import (
	cha "final_project2/channel"
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
	channelPlayer cha.Channel
	channels      []cha.Channel
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
				var channels []cha.Channel
				_, err := p.New(TestUsername, TestRandomNumber, channelPlayer, channels)
				Expect(err).ToNot(BeNil())
				Expect(err.Error()).To(Equal(p.ChannelsListNotGoodErrMsg))
			})
		})
	})
})

func createPlayersChannels() (cha.Channel, []cha.Channel) {
	//Create player channel
	channel1, _ := cha.New(1, make(chan int, 2))
	channelPlayer := channel1
	//Create othe players channels
	channel2, _ := cha.New(1, make(chan int, 2))
	channel2player := channel2
	channel3, _ := cha.New(1, make(chan int, 2))
	channel3player := channel3
	//Create a list of channels and add channels to the list
	var channels []cha.Channel
	channels = append(channels, channel2player)
	channels = append(channels, channel3player)

	return channelPlayer, channels
}
