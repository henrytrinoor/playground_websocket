package main

import (
	"log"
	"net/http"
	// "github.com/gorilla/websocket"
)

func setupAPI() {
	http.Handle(
		"/",
		http.FileServer(http.Dir("./frontend")),
	)
}

func main() {
	setupAPI()

	log.Fatal(http.ListenAndServe(":8080", nil))
}
