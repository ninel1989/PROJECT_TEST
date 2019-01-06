package main

import (
	m "final_project2/manager"
	"fmt"
)

func main() {
	//Create the manager and start the game
	manager := m.GetInstance()
	//Arguments: Number of players, Probability of loosing messages
	if err := manager.StartGame(5, 1); err != nil {
		fmt.Println("Something went wrong!")
	}
}
