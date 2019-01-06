package main

import (
	m "final_project2/manager"
)

func main() {
	//Create the manager and start the game
	manager := m.GetInstance()
	//Arguments: Number of players, Probability of loosing messages
	manager.StartGame(5, 1)
}
