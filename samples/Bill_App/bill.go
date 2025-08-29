package main

type bill struct {
	tip   float64
	items map[string]float64
	name  string
}

func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}
	return b
}

func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}

func (b *bill) addTip(tip float64) {
	b.tip = tip
}
