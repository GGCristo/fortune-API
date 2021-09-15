package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func randomCookie(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(fortunes[rand.Intn(len(fortunes))])
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Take a random fortune\n")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/fortune", randomCookie)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	rand.Seed(time.Now().UnixMicro())
	handleRequests()
}
