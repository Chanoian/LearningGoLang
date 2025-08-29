package main

import "fmt"

func updateName(n *string) {
	*n = "updated"
}

func main() {
	name := "Ara"
	m := &name
	fmt.Println("Memory Address: ", m)
	fmt.Println("Value: ", *m) //value at memory address
	updateName(m)
	fmt.Println("Memory Address: ", m)
	fmt.Println("Value: ", *m)
}
