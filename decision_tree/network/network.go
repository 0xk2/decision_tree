package network

import (
	"log"
	"net/http"

	. "github.com/0xk2/decision_tree/decision_tree/dtree"
)

var Missions = make(map[string]*Mission)

func GoServeTree() {
	http.HandleFunc("/create", CreateHandler)
	http.HandleFunc("/vote", VoteHandler)
	http.HandleFunc("/show", ShowHandler)
	log.Println("Server start at 8080")

	http.ListenAndServe(":8080", nil)
}
