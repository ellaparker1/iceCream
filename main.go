// @author Ella Parker

package main

import (
	"log"
	"net/http"
)

// main - set up a function handler to send some HTML and CSS to a browser
func main() {
	//use FileServer instead of ListenAndServe if running a file SYSTEM instead of one file
	//fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", http.FileServer(http.Dir("./static")))
	//HandleFunc registers the handler function
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {})

	//start the webserver and listen on port 8080
	log.Print("Listening on :8080... ")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("LisenAndServer:", err)
	}
}
