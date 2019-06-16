package channel

import (
	"fmt"
	"math/rand"
)

//Channel - Channel in the game
type Channel struct {
	probability float64
	id          int
	message     string
	startMsg    []int
	aliveMsg    []int
	round       int
}

//ChannelsProbNotGoodErrMsg - Error message
const ChannelsProbNotGoodErrMsg = "Channels propbability must be between (0,1]"

//MessageLostErrMsg - Error message
const MessageLostErrMsg = "The message got lost in the way"

//New - Channel constructor
func New(probability float64, id int, message string, startMsg []int, aliveMsg []int,
	round int) (Channel, error) {
	if probability > 1 || probability <= 0 {
		return Channel{}, fmt.Errorf("%s", ChannelsProbNotGoodErrMsg)
	}
	c := Channel{probability, id, message, startMsg, aliveMsg, round}
	return c, nil
}

//-----------Public functions-----------

//GetRound - return the channel of the user
func (c Channel) GetRound() int {
	return c.round
}

//GetID - return channel id
func (c Channel) GetID() int {
	return c.id
}

//GetMessage - return channel round
func (c Channel) GetMessage() string {
	return c.message
}

//SetRound - setting channel round
func (c *Channel) SetRound(i int) {
	c.round = i
}

//GetStartMsg - return the array of start messages
func (c Channel) GetStartMsg() []int {
	return c.startMsg
}

//GetAliveMsg - return the array of alive messages
func (c Channel) GetAliveMsg() []int {
	return c.aliveMsg
}

//SetStartMsg - return the array of start messages
// func (c Channel) SetStartMsg(i int) {
// 	if i > c.GetStartMsg() {
// 		c.startMsg = i
// 	}

// }

// //SetAliveMsg - return the array of alive messages
// func (c Channel) SetAliveMsg(i int) {
// 	if i > c.GetAliveMsg() {
// 		c.aliveMsg = i
// 	}
// }

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
func (c Channel) InsertMessage(msg string, fromCh int) error {
	randomNumber := rand.Float64()
	if randomNumber < c.probability {
		if msg == "ALIVE" {
			c.aliveMsg[fromCh] = 1
		} else {
			c.startMsg[fromCh] = 1
		}
		return nil
	}
	return fmt.Errorf("%s", MessageLostErrMsg)
}

//InitialMessages - initialize the array of messages
func (c Channel) InitialMessages(id int) {
	c.aliveMsg[id] = 0
	c.startMsg[id] = 0
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
