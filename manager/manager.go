package manager

import (
	cha "final_project3/channel"
	p "final_project3/player"
	"fmt"
	"math/rand"
	"strconv"
	"sync"
)

//Manager - The manager of the game
//@Singelton
type Manager struct {
	players  []p.Player
	channels []cha.Channel
}

var instance *Manager
var once sync.Once

//GetInstance - Get instance of the manager - Manager constructor
func GetInstance() *Manager {
	//Syncronize creation and return
	once.Do(func() {
		instance = &Manager{}
	})
	return instance
}

//-----------Public functions-----------

//StartGame - Starts the game
func (m Manager) StartGame(numOfPlayers int, probability float64) error {

	m.printToConsole("-----------Starts the game...-----------")
	m.printToConsole("Adding players...")
	//Create channels
	for i := 0; i < numOfPlayers; i++ {
		round := make(chan int)
		round <- 0
		var alives []int
		var starts []int
		channel, _ := cha.New(probability, i, "none", starts, alives, round)
		addChannel(channel)
	}
	//Create players
	for i := 0; i < numOfPlayers; i++ {
		e, _ := p.New(fmt.Sprintf("%s%d", "player", i), i, instance.channels[i], getChannelsListWithoutIndex(i))
		addPlayer(e)
	}
	// m.printToConsole("The players in the currnt game:\n" + m.getPlayersList())
	// //Exchange messages between players
	// m.printToConsole("Exchange messages...")
	// for i := 0; i < numOfPlayers; i++ {
	// 	instance.players[i].SendMessagesToAllPlayers()
	// 	m.printToConsole(fmt.Sprintf("Username: %s", instance.players[i].GetUsername()))
	// }
	// //Print sums
	// m.printToConsole("Print sums...")
	// for i := 0; i < numOfPlayers; i++ {
	// 	sum := instance.players[i].GetSum()
	// 	m.printToConsole(fmt.Sprintf("Username: %s, Sum: %d", instance.players[i].GetUsername(), sum))
	// }
	// m.printToConsole("-----------Exiting game...-----------")
	return nil
}

//-----------Private functions-----------

//AddPlayer - Add a player to the players list
func addPlayer(player p.Player) {
	instance.players = append(instance.players, player)
}

//addChannel - Add a channel to the channels list
func addChannel(ch cha.Channel) {
	instance.channels = append(instance.channels, ch)
}

//randomNumber - Create and returns a random number
func randomNumber() int {
	return rand.Intn(100) + 1
}

//returns the channels list without a certain index
func getChannelsListWithoutIndex(index int) []cha.Channel {
	var newChannelsList []cha.Channel
	for i, element := range instance.channels {
		if i != index {
			newChannelsList = append(newChannelsList, element)
		}
	}
	return newChannelsList
}

//getPlayersList - Returns the players list
func (m Manager) getPlayersList() string {
	playersList := ""
	for i, player := range instance.players {
		playersList += strconv.Itoa(i)
		playersList += ": " + player.UserToString() + "\n"
	}
	return playersList
}

//printToConsole - Prints strings to the debug console
func (m Manager) printToConsole(toPrint string) {
	fmt.Println(toPrint)
}
