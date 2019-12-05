package main

import (
	"log"
	"net/http"

	"github.com/adrianna157/trade-wars-client/pkg"
)

var ships []pkg.Ship

func addship(name string) {
	ships = append(ships, pkg.InitShip(name))
}

func updateCargoCookie(w http.ResponseWriter, r *http.Request, shipname string) {
	var cargoCookie, err = r.Cookie("cargo")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "invaild yPos cookie", 500)
		return
	}
	for _, ship := range ships {
		if shipname == ship.Name {
			cargoCookie.Value = ship.GetCargoString()
			http.SetCookie(w, cargoCookie)
		}
	}
}
