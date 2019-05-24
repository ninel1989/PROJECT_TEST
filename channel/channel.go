package channel

import (
	"fmt"
	"math/rand"
)

//Channel - Channel in the game
type Channel struct {
	probability float64
	id          int
	message     chan string
}

//ChannelsProbNotGoodErrMsg - Error message
const ChannelsProbNotGoodErrMsg = "Channels propbability must be between (0,1]"

//MessageLostErrMsg - Error message
const MessageLostErrMsg = "The message got lost in the way"

//New - Channel constructor
func New(probability float64, id int, message chan string) (Channel, error) {
	if probability > 1 || probability <= 0 {
		return Channel{}, fmt.Errorf("%s", ChannelsProbNotGoodErrMsg)
	}
	c := Channel{probability, id, message}
	return c, nil
}

//-----------Public functions-----------

//GetChannel - return the channel of the user
func (c Channel) GetChannel() chan string {
	return c.message
}

//GetID - return channel id
func (c Channel) GetID() int {
	return c.id
}

// InsertNumber - Insert number from the player to the channel
// func (c Channel) InsertNumber(number int) error {
// 	randomNumber := rand.Float64()
// 	if randomNumber < c.probability {
// 		c.ch <- number
// 		return nil
// 	}
// 	return fmt.Errorf("%s", MessageLostErrMsg)
// }

//InsertMessage - Insert number from the player to the channel
func (c Channel) InsertMessage(msg string) error {
	randomNumber := rand.Float64()
	if randomNumber < c.probability {
		c.message <- msg
		return nil
	}
	return fmt.Errorf("%s", MessageLostErrMsg)
}

// GetSum - get all the numbers from the channel, summerizes and prints it
// func (c Channel) GetSum() int {
// 	close(c.message)
// 	sum := 0
// 	for elem := range c.message {
// 		sum += elem
// 	}
// 	return sum
// }
