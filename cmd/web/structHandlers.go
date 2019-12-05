package main

import "github.com/adrianna157/trade-wars-client/pkg"

var ships []pkg.Ship

func addship(name string) {
	ships = append(ships, pkg.InitShip(name))
}
