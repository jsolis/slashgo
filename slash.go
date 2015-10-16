package main

import (
	"io"
	"net/http"
	"log"
)

func main() {
	http.HandleFunc("/slack", handler)
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "I'll be sure to add that to my backlog...\n")
/*
	command := r.FormValue("command")

	if command == "/jsolis" {
		fmt.Fprint(w,"I'll be sure to add that to my backlog")
	} else {
		fmt.Fprint(w,"I do not understand your command.")
	}
*/
}
