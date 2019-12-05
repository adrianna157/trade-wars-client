package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
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
		displayTemplateFile(w, r, "./ui/html/welcome-screen.html")
	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		callSign := r.FormValue("callSign")
		addship(callSign)
		Callcookie := http.Cookie{
			Name:    "callSign",
			Value:   callSign,
			Expires: time.Now().AddDate(0, 0, 1),
			Path:    "/",
		}
		XposCookie := http.Cookie{
			Name:    "xPos",
			Value:   "0",
			Expires: time.Now().AddDate(0, 0, 1),
			Path:    "/map",
		}
		YposCookie := http.Cookie{
			Name:    "yPos",
			Value:   "0",
			Expires: time.Now().AddDate(0, 0, 1),
			Path:    "/map",
		}
		cargoCookie := http.Cookie{
			Name:    "cargo",
			Value:   ships[0].GetCargoString(),
			Expires: time.Now().AddDate(0, 0, 1),
			Path:    "/map",
		}
		//updateCargoCookie(w, r, callSign)
		http.SetCookie(w, &Callcookie)
		http.SetCookie(w, &XposCookie)
		http.SetCookie(w, &YposCookie)
		http.SetCookie(w, &cargoCookie)
		http.Redirect(w, r, "/map", http.StatusSeeOther)
		fmt.Fprintf(w, "Call sign = %s\n", callSign)

	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func showMapScreen(w http.ResponseWriter, r *http.Request) {
	var callSignCookie, err = r.Cookie("callSign")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "No call sign obtained from cookie", 500)
		return
	}
	switch r.Method {
	case "GET":
		displayTemplateFile(w, r, "./ui/html/navigation-screen.html")
	case "POST":
		callSign := callSignCookie.Value
		log.Println(callSign)
		w.Write([]byte(callSign))
	}

}

func showTradeScreen(w http.ResponseWriter, r *http.Request) {
	displayTemplateFile(w, r, "./ui/html/trade-screen.html")
}

func showChatScreen(w http.ResponseWriter, r *http.Request) {
	displayTemplateFile(w, r, "./ui/html/chat-screen.html")
}

func ws(w http.ResponseWriter, r *http.Request) {
	displayTemplateFile(w, r, "./ui/html/chat-room.html")
}

func redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/players", 301)
}

func moveLeft(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var xPosCookie, err = r.Cookie("xPos")
		value, err := strconv.Atoi(xPosCookie.Value)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "invaild xPos cookie", 500)
			return
		}
		if value < 1 {
			xPosCookie.Value = strconv.Itoa(4)
		} else {
			xPosCookie.Value = strconv.Itoa(value - 1)
		}
		http.SetCookie(w, xPosCookie)
		http.Redirect(w, r, "/map", http.StatusSeeOther)
		log.Println(value)
		w.Write([]byte(xPosCookie.Value))
	}
}
func moveRight(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var xPosCookie, err = r.Cookie("xPos")
		value, err := strconv.Atoi(xPosCookie.Value)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "invaild xPos cookie", 500)
			return
		}
		if value > 3 {
			xPosCookie.Value = strconv.Itoa(0)
		} else {
			xPosCookie.Value = strconv.Itoa(value + 1)
		}
		http.SetCookie(w, xPosCookie)
		http.Redirect(w, r, "/map", http.StatusSeeOther)
		log.Println(value)
		w.Write([]byte(xPosCookie.Value))
	}
}

func moveUp(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var yPosCookie, err = r.Cookie("yPos")
		value, err := strconv.Atoi(yPosCookie.Value)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "invaild yPos cookie", 500)
			return
		}
		if value < 1 {
			yPosCookie.Value = strconv.Itoa(4)
		} else {
			yPosCookie.Value = strconv.Itoa(value - 1)
		}
		http.SetCookie(w, yPosCookie)
		http.Redirect(w, r, "/map", http.StatusSeeOther)
		log.Println(value)
		w.Write([]byte(yPosCookie.Value))
	}
}
func moveDown(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var yPosCookie, err = r.Cookie("yPos")
		value, err := strconv.Atoi(yPosCookie.Value)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "invaild yPos cookie", 500)
			return
		}
		if value > 3 {
			yPosCookie.Value = strconv.Itoa(0)
		} else {
			yPosCookie.Value = strconv.Itoa(value + 1)
		}
		http.SetCookie(w, yPosCookie)
		http.Redirect(w, r, "/map", http.StatusSeeOther)
		log.Println(value)
		w.Write([]byte(yPosCookie.Value))
	}
}
