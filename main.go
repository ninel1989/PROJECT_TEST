package main

import (
	m "final_project2/manager"
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

	//runScenario()
}

func handleRequest(c *gin.Context) {

	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.WriteString("Response from the GO Server\n")
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
