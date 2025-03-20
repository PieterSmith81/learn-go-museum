package main

import (
	"fmt"
	"net/http"
	"text/template"
)

// Route handler functions
func handleHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from a Go server!"))
}

func handleTemplate(w http.ResponseWriter, r *http.Request) {
	// Read/parse the Go template file into a variable
	html, err := template.ParseFiles("templates/index.tmpl")

	// If an error occurred parsing the go template file, handle that error gracefully
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error. Please contact the website administrator."))
		return
	}

	// Execute/write the template to the client
	html.Execute(w, "We are now using Go templates!")
}

func main() {
	// Create (and assign) a new Go HTTP server variable, so that you don't need to use the net/http package name all the time
	server := http.NewServeMux()

	/* A basic implementation of a HTTP route handler in Go.
	Open your browser and navigate to localhost:3000/hello to see the response. */
	server.HandleFunc("/hello", handleHello)

	/* A Go template based HTTP route handler.
	Open your browser and navigate to localhost:3000/template to see the result. */
	server.HandleFunc("/template", handleTemplate)

	/* A more advanced implementation of a HTTP route handler in Go.
	This time serving public, static website files.
	Open your browser and navigate to localhost:3000 to see the result. */
	// Use the http.FileServer function to create a handler to serve our public, static website files
	fs := http.FileServer(http.Dir("./public"))

	// Serve our public, static website files using the FileServer handler
	server.Handle("/", fs)

	/* Start the web server */
	// Open a web server and listen for incoming connections on a TCP port
	err := http.ListenAndServe(":3000", server)

	// Check for errors after starting the server / opening the TCP port
	if err != nil {
		fmt.Println("Error while starting the web server or opening the TCP port.")
	}
}
