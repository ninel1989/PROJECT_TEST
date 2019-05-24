package player

import (
	"fmt"
)

//Player - Player in the game
type Player struct {
	username             string
	number               int
	ch                   cha.Channel
	otherPlayersChannels []cha.Channel
}

//ChannelsListNotGoodErrMsg - Error message
const ChannelsListNotGoodErrMsg = "Channels propbability must be between (0,1]"

//New - Player constructor
func New(username string, number int, userChannel cha.Channel, channelsList []cha.Channel) (Player, error) {
	if channelsList == nil || len(channelsList) == 0 {
		return Player{}, fmt.Errorf("%s", ChannelsListNotGoodErrMsg)
	}
	e := Player{username, number, userChannel, channelsList}
	return e, nil
}

//-----------Public functions-----------

//UserToString - Returns a representation of the player
func (e Player) UserToString() string {
	return fmt.Sprintf("Username: %s, User number: %d", e.username, e.number)
}

//GetUsername - return the username
func (e Player) GetUsername() string {
	return e.username
}

//GetRandomNumber - return the random number of the user
func (e Player) GetNumber() int {
	return e.number
}

//GetChannel - return the channel of the user
func (e Player) GetChannel() chan string {
	return e.ch.GetChannel()
}

//GetotherPlayersChannels - return other players channels
func (e Player) GetotherPlayersChannels() []cha.Channel {
	return e.otherPlayersChannels
}

//SendMessagesToAllPlayers - Sends the random number of the user to all the others players channels
func (e Player) SendMessagesToAllPlayers() int {
	countLostMessages := 0
	for _, element := range e.otherPlayersChannels {
		if err := e.sendNumber(element); err != nil {
			countLostMessages++
		}
	}
	return countLostMessages
}

//GetSum - get all the numbers from the channel, summerizes and prints it
func (e Player) GetSum() int {
	sum := e.ch.GetSum()
	sum = sum + e.number
	return sum
}

func (e Player) LeaderAlgo(int alfa, int beta, int delta) {
	currentRound = 0
	recTimer = 0
	var sendTimer = 0
	a = alfa
	b = beta
	d = delta

	while(true)
	{
		for _, element := range e.otherPlayersChannels {
			if e.ch.message == "START" || e.ch.message == "ALIVE" {
				if currentRound > e.ch.GetID() {
					e.sendMessage("START")
				} else {
					if currentRound < e.ch.GetID() {
						startRound(e.ch.GetID)
					}
					recTimer = 0
				}
			}
		}
		recTimer = recTimer + 1
		if recTimer > 8*Round(d/a) {
			if e.GetNumber() != (currentRound % 11) {
				startRound(currentRound + 1)
			}
			recTimer = 0
		}
	}

}

//-----------Private functions-----------

//sendNumber - Sends the random number of the user to the channel (argument)
//func (e Player) sendNumber(channel cha.Channel) error {
//e.number is always the same!!!
//	if err := channel.InsertNumber(e.number); err != nil {
//		return err
//	}
//	return nil
//}

func (e Player) sendMessage(channel cha.Channel, string msg) error {
	if err := channel.InsertMessage(msg); err != nil {
		return err
	}
	return nil
}

func (e player) startRound(int s) {

}
