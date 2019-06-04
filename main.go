package main

import (
	m "final_project3/manager"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	r := gin.Default()

	//Use cors for go server (Cross-Origin Resource Sharing)
	r.Use(cors.Default())

	projectAPI := r.Group("/project")
	projectAPI.GET("/", handleRequest)

	r.Run()

	// runScenario()
}

func handleRequest(c *gin.Context) {

	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.WriteString(algorithmResultsExample())
}

func algorithmResultsExample() string {
	return "3:2,3,Alive-4,2,Alive-5,1,Start-4,5,Alive-4,3,Start-4,1,Start-2,3,Alive-3,2,Start-1,3,Alive"
}

func runScenario() {
	//Create unique seed for this program - different random numbers every time
	rand.Seed(time.Now().UnixNano())

	//Create the manager and start the game
	manager := m.GetInstance()

	//Arguments: Number of players, Probability of loosing messages
	if err := manager.StartGame(5, 1); err != nil {
		fmt.Println("Something went wrong!")
	}
}
