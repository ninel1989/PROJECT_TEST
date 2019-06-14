package player

import (
	cha "final_project3/channel"
	"fmt"
	"math"
	"time"
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
func (e Player) GetChannel() chan int {
	return e.ch.GetChannel()
}

//GetotherPlayersChannels - return other players channels
func (e Player) GetotherPlayersChannels() []cha.Channel {
	return e.otherPlayersChannels
}

//SendMessagesToAllPlayers - Sends the random number of the user to all the others players channels
func (e Player) SendMessagesToAllPlayers(msg string) {
	//countLostMessages := 0
	for _, element := range e.otherPlayersChannels {
		if err := e.sendMessage(element, msg); err != nil {
			//countLostMessages++
		}
	}
	//return countLostMessages
}

//LeaderAlgo - execute the second algorithm
func (e Player) LeaderAlgo(alfa int, beta int, delta int) int {
	var currentRound = 0
	var recTimer = 0
	var sendTimer = 0
	var a = alfa
	var b = beta
	var d = delta
	var Leader = -1
	e.GetChannel() <- 0
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		fmt.Printf("Run number %d, player %s\n", i, e.username)
		for _, element := range e.otherPlayersChannels {
			starts := e.ch.GetStartMsg()
			alives := e.ch.GetAliveMsg()
			//s := strings.Split(msg, ",")
			//messageFrom, otherRound := s[0], s[1]
			if ((starts[element.GetID()]) > 0) || ((alives[element.GetID()]) > 0) {
				starts[element.GetID()]--
				alives[element.GetID()]--
				otherRound := <-element.GetChannel()
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
			if e.GetNumber() != (currentRound % (len(e.GetotherPlayersChannels())+1)) {
				e.startRound(currentRound + 1)
				sendTimer = int(d / b)
			}
			recTimer = 0
		}
		sendTimer = sendTimer + 1
		if sendTimer >= int(d/b) {
			if e.GetNumber() == (currentRound % (len(e.GetotherPlayersChannels())+1)) {
				e.SendMessagesToAllPlayers("ALIVE")
			}
			Leader = (currentRound % (len(e.GetotherPlayersChannels())+1))
			sendTimer = 0
		}
	}
	return Leader
}

func (e Player) sendMessage(channel cha.Channel, msg string) error {
	if err := channel.InsertMessage(msg, e.ch.GetID()); err != nil {
		return err
	}
	return nil
}

func (e Player) startRound(s int) {
	if e.GetNumber() != (s % (len(e.GetotherPlayersChannels())+1)) {
		i := (s % (len(e.GetotherPlayersChannels())+1))
		element := e.otherPlayersChannels[i]
		e.sendMessage(element, "START")
	}
	e.ch.SetRound(s)
}
