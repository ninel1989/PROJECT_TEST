package channel

import (
	"fmt"
)

//Channel - Channel in the game
type Channel struct {
	probability int
	ch          chan int
}

//ChannelsProbNotGoodErrMsg - Error message
const ChannelsProbNotGoodErrMsg = "Channels propbability must be between (0,1]"

//New - Channel constructor
func New(probability int, channel chan int) (Channel, error) {
	if probability > 1 || probability <= 0 {
		return Channel{}, fmt.Errorf("%s", ChannelsProbNotGoodErrMsg)
	}
	c := Channel{probability, channel}
	return c, nil
}

//-----------Public functions-----------

//InsertNumber - Insert number from the player to the channel
func (c Channel) InsertNumber(number int) {
	c.ch <- number
}

//GetChannel - return the channel of the user
func (c Channel) GetChannel() chan int {
	return c.ch
}

//GetSum - get all the numbers from the channel, summerizes and prints it
func (c Channel) GetSum(numOfPlayers int, playerNumber int) int {
	close(c.ch)
	sum := playerNumber
	for elem := range c.ch {
		sum += elem
	}
	return sum
}
