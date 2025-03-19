package main

import (
	"fmt"
	"net/http"
)

func main() {
	/* A basic implementation of a HTTP route handler in Go.
	Open your browser and navigate to localhost:3000 to see the response. */
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from a Go server!"))
	})

	// Open a web server and listen for incoming connections on a TCP port
	err := http.ListenAndServe(":3000", nil)

	// Check for errors after starting the server / opening the TCP port
	if err != nil {
		fmt.Println("Error while starting the web server or opening the TCP port.")
	}
}
