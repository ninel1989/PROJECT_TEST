package player

import (
	cha "final_project3/channel"
	"fmt"
	"time"
)

//Player - Player in the game
type Player struct {
	username             string
	number               int
	ch                   cha.Channel
	otherPlayersChannels []cha.Channel
	allChan              []cha.Channel
}

//ChannelsListNotGoodErrMsg - Error message
const ChannelsListNotGoodErrMsg = "Channels propbability must be between (0,1]"

//New - Player constructor
func New(username string, number int, userChannel cha.Channel, channelsList []cha.Channel, allChan []cha.Channel) (Player, error) {
	if channelsList == nil || len(channelsList) == 0 {
		return Player{}, fmt.Errorf("%s", ChannelsListNotGoodErrMsg)
	}
	e := Player{username, number, userChannel, channelsList, allChan}
	return e, nil
}

//-----------Public functions-----------

//UserToString - Returns a representation of the player
func (e Player) UserToString() string {
	return fmt.Sprintf("Username: %d, User number: %d", e.GetNumber(), e.number)
}

//GetUsername - return the username
func (e Player) GetUsername() string {
	return e.username
}

//GetNumber - return the random number of the user
func (e Player) GetNumber() int {
	return e.number
}

// //GetChannel - return the channel of the user
// func (e Player) GetChannel() chan int {
// 	return e.ch.GetChannel()
// }

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
	var recTimer = 0
	var sendTimer = 0
	//var a = alfa
	var b = beta
	var d = delta
	var Leader = -1
	//e.GetChannel() <- round
	//x := <-e.GetChannel()
	fmt.Printf("Player %d start the algorithm\n", e.GetNumber())
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		fmt.Printf("Run number %d, player %s\n", i, e.username)
		for _, element := range e.otherPlayersChannels {
			starts := e.ch.GetStartMsg()
			alives := e.ch.GetAliveMsg()
			//s := strings.Split(msg, ",")
			//messageFrom, otherRound := s[0], s[1]
			if ((starts[element.GetID()]) > 0) || ((alives[element.GetID()]) > 0) {
				fmt.Printf("number %d in the condition, get message from:%d\n", e.GetNumber(), element.GetID())
				e.ch.InitialMessages(element.GetID())
				otherRound := element.GetRound()
				currentRound := e.ch.GetRound()
				fmt.Printf("number %d,my round is:%d, other round is:%d\n", e.GetNumber(), currentRound, otherRound)
				if currentRound > otherRound {
					fmt.Printf("number %d send start I am the leader to, %d \n", e.GetNumber(), element.GetID())
					e.sendMessage(element, "START")
				} else {

					if currentRound < otherRound {
						fmt.Printf("number %d is going to start new round, %d\n", e.GetNumber(), otherRound)
						e.startRound(otherRound)
						e.ch.SetRound(otherRound)
						fmt.Printf("number %d change is round to, %d\n", e.GetNumber(), e.ch.GetRound())
						sendTimer = int(d / b)
						currentRound = e.ch.GetRound()
					} else {
						fmt.Printf("number %d and number, %d are equals\n", e.GetNumber(), element.GetID())
					}
				}
				recTimer = 0
			}
		}
		recTimer = recTimer + 1
		// if recTimer > 8*int(math.Round(float64(d/a))) {
		// 	if e.GetNumber() != (currentRound % len(e.GetotherPlayersChannels())) {
		// 		e.startRound(currentRound + 1)
		// 		sendTimer = int(d / b)
		// 	}
		// 	recTimer = 0
		// }
		sendTimer = sendTimer + 1
		if e.GetNumber() == (e.ch.GetRound() % len(e.allChan)) {
			fmt.Printf("number %d send alive to all players, my current round is:%d\n", e.GetNumber(), e.ch.GetRound())
			e.SendMessagesToAllPlayers("ALIVE")
		}
		Leader = (e.ch.GetRound() % len(e.allChan))
		fmt.Printf("number %d,my leader in run %d, is :%d\n", e.GetNumber(), i, Leader)
		sendTimer = 0

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
	fmt.Printf("number %d start round number, %d\n", e.GetNumber(), s)
	if e.GetNumber() != (s % len(e.allChan)) {
		i := (s % len(e.allChan))
		fmt.Printf("number %d start round number, %d, the result of mod is,%d\n", e.GetNumber(), s, i)
		element := e.allChan[i]
		fmt.Printf("number %d send start i am with you to, %d\n", e.GetNumber(), element.GetID())
		e.sendMessage(element, "START")
	}

}
