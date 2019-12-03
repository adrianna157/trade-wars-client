package pkg

import "fmt"

type Cargo struct {
	Name  string
	Price int
}

func (cargo Cargo) GetInfo() string {
	return fmt.Sprintf(cargo.Name+" $%d", cargo.Price)
}
