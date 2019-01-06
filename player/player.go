package player

import (
	cha "final_project2/channel"
	"fmt"
)

//Player - Player in the game
type Player struct {
	username             string
	number               int
	ch                   cha.Channel
	otherPlayersChannels []chan int
}

//ChannelsListNotGoodErrMsg - Error message
const ChannelsListNotGoodErrMsg = "Channels propbability must be between (0,1]"

//New - Player constructor
func New(username string, number int, userChannel chan int, channelsList []chan int) (Player, error) {
	if channelsList == nil {
		return Player{}, fmt.Errorf("%s", ChannelsListNotGoodErrMsg)
	}
	c, _ := cha.New(1, userChannel)
	e := Player{username, number, c, channelsList}
	return e, nil
}

//-----------Public functions-----------

//UserToString - Returns a representation of the player
func (e Player) UserToString() string {
	return fmt.Sprintf("username: %s, user number: %d", e.username, e.number)
}

//GetUsername - return the username
func (e Player) GetUsername() string {
	return e.username
}

//GetRandomNumber - return the random number of the user
func (e Player) GetRandomNumber() int {
	return e.number
}

//GetChannel - return the channel of the user
func (e Player) GetChannel() chan int {
	return e.ch.GetChannel()
}

//GetotherPlayersChannels - return other players channels
func (e Player) GetotherPlayersChannels() []chan int {
	return e.otherPlayersChannels
}

//SendMessagesToAllPlayers - Sends the random number of the user to all the others players channels
func (e Player) SendMessagesToAllPlayers() {
	for _, element := range e.otherPlayersChannels {
		e.sendNumber(element)
	}
}

//GetSum - get all the numbers from the channel, summerizes and prints it
func (e Player) GetSum(numOfPlayers int) string {
	sum := e.ch.GetSum(numOfPlayers, e.number)
	return fmt.Sprintf("Username: %s, Sum: %d", e.username, sum)
}

//-----------Private functions-----------

//sendNumber - Sends the random number of the user to the channel (argument)
func (e Player) sendNumber(channel chan int) {
	//e.number is always the same!!!
	channel <- e.number
}
