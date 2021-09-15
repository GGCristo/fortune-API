package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func randomCookie(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(fortunes[rand.Intn(len(fortunes))])
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Take a random fortune\n")
}

func main() {
	port := os.Getenv("PORT")
	rand.Seed(time.Now().UnixNano())
	http.HandleFunc("/", homePage)
	http.HandleFunc("/fortune", randomCookie)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
