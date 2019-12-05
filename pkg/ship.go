package pkg

type Ship struct {
	Name       string
	xPos, yPos int
	Cargo      []Cargo
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

func InitShip(name string) Ship {
	return Ship{
		Name:  name,
		xPos:  0,
		yPos:  0,
		Cargo: initCargo(),
	}
}

func (ship Ship) GetInfo() (info string) {
	info = ship.Name + "\nCargo:\n"
	for _, cargo := range ship.Cargo {
		info += cargo.GetInfo() + "\n"
	}
	return
}
func (ship Ship) GetCargoString() (info string) {
	info = "Cargo:\n"
	for _, cargo := range ship.Cargo {
		info += cargo.GetInfo() + "\n"
	}
	return
}

func initCargo() []Cargo {
	return []Cargo{
		Cargo{Name: "cheese", Price: 4},
		Cargo{Name: "foo", Price: 4},
		Cargo{Name: "water", Price: 4},
		Cargo{Name: "tofu", Price: 4},
	}
}
