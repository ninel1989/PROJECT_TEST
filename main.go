package main

import (
	m "final_project2/manager"
)

func main() {
	//Create the manager and start the game
	manager := m.GetInstance()
	manager.StartGame(5)
}
