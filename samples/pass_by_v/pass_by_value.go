package main

import "fmt"

func updateName(x string) {
	x = "updated"
	// x is a copy not the original variable
}

func updateMapCoffee(x map[string]float64) {
	x["coffee"] = 100
}

func main() {
	// group A (type) -> strings, ints. bools, floats , arrays and structs (are pass by value)

	name := "Ara"
	updateName(name)
	fmt.Println(name)

	menu := map[string]float64{
		"burger": 4.99,
		"fries":  3.99,
		"shake":  2.99,
		"coffe":  1.99,
	}

	updateMapCoffee(menu)
	fmt.Println(menu)
}
