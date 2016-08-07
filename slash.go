package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

func main() {
	http.HandleFunc("/slack", handler)
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

type Message struct {
	Channel    string `json:"channel"`
	Text       string `json:"text"`
	Username   string `json:"username"`
	Icon_emoji string `json:"icon_emoji"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	//io.WriteString(w, "I'll be sure to add that to my backlog...\n")
/*
	command := r.FormValue("command")

	if command == "/jsolis" {
		fmt.Fprint(w,"I'll be sure to add that to my backlog")
	} else {
		fmt.Fprint(w,"I do not understand your command.")
	}
*/

	text := r.FormValue("text")

	m := Message{"#scotty-testing", "Request: " + text, "Jay-Slack-Bot", ":thinking_face:"}
	b, err := json.Marshal(m)
	if err != nil {
		return
	}

	v := url.Values{}
	v.Set("payload", string(b))

	incoming_webhook_url := os.Getenv("SLASHGO_INCOMING_WEBHOOK_URL")
	resp, err := http.PostForm(incoming_webhook_url, v)

	if err != nil {
		return
	}
	defer resp.Body.Close()

	io.WriteString(w, "Your request has been forwarded\n")

}
