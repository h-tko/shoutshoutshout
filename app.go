package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"shoutshoutshout/libraries"
	//for extracting service credentials from VCAP_SERVICES
	//"github.com/cloudfoundry-community/go-cfenv"
)

const (
	DEFAULT_PORT = "8080"
)

type Page struct {
	Shout      string
	BotMessage string
}

var index = template.Must(template.ParseFiles(
	"templates/_base.html",
	"templates/index.html",
))

func get(w http.ResponseWriter, req *http.Request) {

	if req.Method == "GET" {
		index.Execute(w, nil)
	} else {
		req.ParseForm()
		shout := req.Form["Shout"][0]

		watson := libraries.NewWatson()
		data := watson.ConversationApi(shout)

		page := Page{shout, data}

		index.Execute(w, page)
	}

}

func main() {
	var port string
	if port = os.Getenv("PORT"); len(port) == 0 {
		port = DEFAULT_PORT
	}

	http.HandleFunc("/", get)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Printf("Starting app on port %+v\n", port)
	http.ListenAndServe(":"+port, nil)
}
