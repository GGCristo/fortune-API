package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func randomCookie(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, fortunes[rand.Intn(len(fortunes))])
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
	// log.Fatal(http.ListenAndServe(":8080", nil))
}
