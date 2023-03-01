package main

import (
	"assignment-1/global"
	"assignment-1/handler"
	"log"
	"net/http"
	"os"
	"time"
)

/*
Main method of the application. Start the REST server.
*/
func main() {

	// Check for port enviromental variable
	port := os.Getenv("PORT")

	// If no environmental variable, set port to DEFAULT_PORT
	if port == "" {
		log.Println("$PORT has not been set. Default: " + global.DEFAULT_PORT)
		port = global.DEFAULT_PORT
	}

	// Endpoint handlers
	http.HandleFunc(global.DEFAULT_PATH, handler.DefaultPathHandler)
	http.HandleFunc(global.UNI_INFO_PATH, handler.UniversityHandler)
	http.HandleFunc(global.NEIGHBOUR_UNI_PATH, handler.NeighbouringUniHandler)
	http.HandleFunc(global.DIAG_PATH, handler.DiagnosticsHandler)

	// Start the server and capture the start time
	global.START_TIME = time.Now().Unix()
	log.Println("Listening on port: " + global.DEFAULT_PORT + " ...")
	log.Fatal(http.ListenAndServe(":"+global.DEFAULT_PORT, nil))
}
