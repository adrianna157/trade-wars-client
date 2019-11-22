package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func displayTemplateFile(w http.ResponseWriter, r *http.Request, pathToFile string) {
	files := []string{
		pathToFile,
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func players(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/players" {
		http.NotFound(w, r)
		return
	}

	switch r.Method {
	case "GET":
		displayTemplateFile(w, r, "./ui/html/welcome-screen.tmpl")
	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		handler := r.FormValue("handler")
		fmt.Fprintf(w, "Handler = %s\n", handler)
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func showNavigationScreen(w http.ResponseWriter, r *http.Request) {
	displayTemplateFile(w, r, "./ui/html/navigation-screen.tmpl")
}

func showTradeScreen(w http.ResponseWriter, r *http.Request) {
	displayTemplateFile(w, r, "./ui/html/trade-screen.tmpl")
}

func showChatScreen(w http.ResponseWriter, r *http.Request) {
	displayTemplateFile(w, r, "./ui/html/chat-screen.tmpl")
}

func ws(w http.ResponseWriter, r *http.Request) {
	displayTemplateFile(w, r, "./ui/html/chat-room.tmpl")
}

func redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/players", 301)
}
