package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"
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
		callSign := r.FormValue("callSign")
		cookie := http.Cookie{
			Name:    "callSign",
			Value:   callSign,
			Expires: time.Now().AddDate(0, 0, 1),
			Path:    "/",
		}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/map", http.StatusSeeOther)
		fmt.Fprintf(w, "Call sign = %s\n", callSign)

	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func showMapScreen(w http.ResponseWriter, r *http.Request) {
	var cookie, err = r.Cookie("callSign")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "No call sign obtained from cookie", 500)
		return
	}
	switch r.Method {
	case "GET":
		displayTemplateFile(w, r, "./ui/html/navigation-screen.tmpl")
	case "POST":
		callSign := cookie.Value
		log.Println(callSign)
		w.Write([]byte(callSign))
	}

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

func moveLeft(w http.ResponseWriter, r *http.Request) {
	if r.Method == "Get" {
		var cookie, err = r.Cookie("callSign")
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "No call sign obtained from cookie", 500)
			return
		}
		callSign := cookie.Value
		log.Println(callSign)
	}
}
func moveRight(w http.ResponseWriter, r *http.Request) {
	if r.Method == "Get" {
		var cookie, err = r.Cookie("callSign")
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "No call sign obtained from cookie", 500)
			return
		}
		callSign := cookie.Value
		log.Println(callSign)
	}
}
