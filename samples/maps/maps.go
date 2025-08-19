package main

import "fmt"

func main() {
	// string is key and float64 is value
	menu := map[string]float64{
		"burger": 4.99,
		"fries":  3.99,
		"shake":  2.99,
	}

	menu["burger"] = 5.99
	delete(menu, "shake")
	fmt.Println(menu)

	for k, v := range menu {
		fmt.Println(k, "_", v)
	}
}
