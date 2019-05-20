package main

import (
	m "final_project2/manager"
	"fmt"
	"github.com/gorilla/mux"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]

		fmt.Fprintf(w, "You've requested the book ###: %s on page %s\n", title, page)
	})

	http.ListenAndServe(":8000", r)
	//runScenario()
}

func runScenario() {
	//Create uniq seed for this program - different random numbers every time
	rand.Seed(time.Now().UnixNano())

	//Create the manager and start the game
	manager := m.GetInstance()

	//Arguments: Number of players, Probability of loosing messages
	if err := manager.StartGame(5, 1); err != nil {
		fmt.Println("Something went wrong!")
	}
}
