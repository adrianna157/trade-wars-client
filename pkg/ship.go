package pkg

type Ship struct {
	Name  string
	Cargo []Cargo
}

func (ship Ship) AddCargo(item Cargo) {
	ship.Cargo = append(ship.Cargo, item)
}

func (ship Ship) RemoveCargo(item string) {
	index := findCargoIndex(item, ship.Cargo)
	if index > -1 {
		ship.Cargo = ship.Cargo[:index+copy(ship.Cargo[index:], ship.Cargo[index+1:])]
	}
}

func findCargoIndex(item string, cargo []Cargo) int {
	for i, c := range cargo {
		if item == c.Name {
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
