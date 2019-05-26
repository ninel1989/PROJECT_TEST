package player

import (
	cha "final_project3/channel"
	"fmt"
	"math"
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

//GetNumber - return the random number of the user
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
func (e Player) SendMessagesToAllPlayers() {
	//countLostMessages := 0
	msg := <-e.GetChannel()
	for _, element := range e.otherPlayersChannels {
		if err := e.sendMessage(element, msg); err != nil {
			//countLostMessages++
		}
	}
	//return countLostMessages
}

//SendMessagesToAllPlayers - Sends the random number of the user to all the others players channels
// func (e Player) SendMessagesToAllPlayers() int {
// 	countLostMessages := 0
// 	for _, element := range e.otherPlayersChannels {
// 		if err := e.sendNumber(element); err != nil {
// 			countLostMessages++
// 		}
// 	}
// 	return countLostMessages
// }

//GetSum - get all the numbers from the channel, summerizes and prints it
// func (e Player) GetSum() int {
// 	sum := e.ch.GetSum()
// 	sum = sum + e.number
// 	return sum
// }

//LeaderAlgo - execute the second algorithm
func (e Player) LeaderAlgo(alfa int, beta int, delta int) int {
	var currentRound = 0
	var recTimer = 0
	var sendTimer = 0
	var a = alfa
	var b = beta
	var d = delta
	var Leader = -1

	for {
		for _, element := range e.otherPlayersChannels {
			msg := <-e.GetChannel()
			//s := strings.Split(msg, ",")
			//messageFrom, otherRound := s[0], s[1]
			if msg == "START" || msg == "ALIVE" {
				otherRound := element.GetRound()
				if currentRound > otherRound {
					e.sendMessage(element, "START")
				} else {

					if currentRound < otherRound {
						e.startRound(otherRound)
						sendTimer = int(d / b)
					}
				}
				recTimer = 0
			}
		}
		recTimer = recTimer + 1
		if recTimer > 8*int(math.Round(float64(d/a))) {
			if e.GetNumber() != (currentRound % 11) {
				e.startRound(currentRound + 1)
			}
			recTimer = 0
		}
		sendTimer = sendTimer + 1
		if sendTimer >= int(d/b) {
			if e.GetNumber() == (currentRound % 11) {
				msgToSend := "ALIVE"
				e.GetChannel() <- msgToSend
				e.SendMessagesToAllPlayers()
			}
			Leader = (currentRound % 11)
			sendTimer = 0
		}
	}
	return Leader
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

func (e Player) sendMessage(channel cha.Channel, msg string) error {
	if err := channel.InsertMessage(msg); err != nil {
		return err
	}
	return nil
}

func (e Player) startRound(s int) {
	if e.GetNumber() != (s % 11) {

	}
	e.ch.SetRound(s)
}
