package pkg

type Ship struct {
	Name  string
	Cargo []Cargo
}

func (ship Ship) AddCargo(item Cargo) {
	ship.Cargo = append(ship.Cargo, item)
}

func (ship Ship) RemoveCargo(item Cargo) {
	index := findCargoIndex(item, ship.Cargo)
	ship.Cargo = append(ship.Cargo[:index], ship.Cargo[index+1:]...)
}

func findCargoIndex(item Cargo, cargo []Cargo) int {
	for i, c := range cargo {
		if item == c {
			return i
		}
	}
	return -1
}

func (ship Ship) GetInfo() (info string) {
	info = ship.Name + "\nCargo:\n"
	for _, cargo := range ship.Cargo {
		info += cargo.GetInfo() + "\n"
	}
	return
}
