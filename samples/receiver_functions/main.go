package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// format the bill
func (b *bill) format() string {
	fs := "Bill Breakdown:\n"
	var total float64 = 0
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ..$%v \n", k+":", v)
		total += v
	}
	fs += fmt.Sprintf("%-25v ...$%0.2f \n", "tip:", b.tip)
	fs += fmt.Sprintf("%-25v ...$%0.2f \n", "total:", total)
	return fs
}

func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{}, //empty map
		tip:   0,                    // total is 0
	}
	return b
}

func (b *bill) updateTip(tip float64) {
	b.tip = tip
}

func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}

func main() {
	mybill := newBill("ara")
	mybill.addItem("burger", 4.99)
	mybill.addItem("pizza", 5.99)
	mybill.addItem("coke", 1.99)
	mybill.updateTip(5)
	fmt.Println(mybill.format()) //fmt.Println(mybill)
}
