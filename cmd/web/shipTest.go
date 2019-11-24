package main

import "github.com/adrianna157/trade-wars-client/pkg"

func mockCargo() []pkg.Cargo {
	return []pkg.Cargo{
		pkg.Cargo{Name: "cheese", Price: 4},
		pkg.Cargo{Name: "foo", Price: 4},
		pkg.Cargo{Name: "water", Price: 4},
		pkg.Cargo{Name: "tofu", Price: 4},
	}
}
func testFakeShip() {
	shipFake := pkg.Ship{Name: "Enterprise", Cargo: mockCargo()}
	println(shipFake.GetInfo())
	shipFake.AddCargo(pkg.Cargo{Name: "CoCo", Price: 8})
	println(shipFake.GetInfo())
	shipFake.RemoveCargo("cheese")
	shipFake.RemoveCargo("foo")
	println(shipFake.GetInfo())
}
