package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	total float64
}

func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{}, //empty map
		total: 0,                    // total is 0
	}
	return b
}

func main() {
	mybill := newBill("ara")
	mybill.total = 100
	mybill.items["burger"] = 4.99
	mybill.items["fries"] = 3.99
	fmt.Println(mybill)
}
