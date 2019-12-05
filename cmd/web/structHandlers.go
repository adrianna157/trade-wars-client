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
		http.Error(w, "invaild cargo cookie", 500)
		return
	}
	cargoCookie.Value = getShipCargoString(shipname)
	http.SetCookie(w, cargoCookie)
}

func getShipCargoString(shipname string) string {
	for _, ship := range ships {
		if shipname == ship.Name {
			return ship.GetCargoString()
		}
	}
	return " "
}

func addCargo(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var callCookie, err = r.Cookie("callSign")
		var cargoCookie, er = r.Cookie("cargo")
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "invaild callSign cookie", 500)
			return
		}
		if er != nil {
			log.Println(er.Error())
			http.Error(w, "invaild  cargo cookie", 500)
			return
		}
		shipname := callCookie.Value
		log.Println(shipname)
		for _, ship := range ships {
			if shipname == ship.Name {
				ship.AddCargo(pkg.Cargo{Name: "foo", Price: 0})
				log.Panicln(ship.GetCargoString())
			}
		}
		cargoCookie.Value = getShipCargoString(shipname)
		http.SetCookie(w, cargoCookie)
		http.Redirect(w, r, "/map/trade", http.StatusSeeOther)
	}
}
