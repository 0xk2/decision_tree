package network

import (
	"log"
	"net/http"
	"os"

	. "github.com/0xk2/decision_tree/decision_tree/dtree"
	"github.com/joho/godotenv"
)

var Missions = make(map[string]*Mission)

func GoServeTree() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	http.HandleFunc("/create", CreateHandler)
	http.HandleFunc("/vote", VoteHandler)
	http.HandleFunc("/show", ShowHandler)
	http.HandleFunc("/check", healthCheck)
	log.Println("Server start at " + port)

	http.ListenAndServe(":"+port, nil)
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
