package main

import (
	"fmt"
	"math"
	"strings"
)

func sayGreeting(n string) {
	fmt.Printf("Hello, %v \n", n)
}

func sayBye(n string) {
	fmt.Printf("Bye, %v \n", n)
}

func cycleNames(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}

func circleAre(r float64) float64 {
	return math.Pi * r * r
}

func getInitials(n string) (string, string) {
	names := strings.Split(strings.ToUpper(n), " ")
	fmt.Println(names)
	var initials []string
	for _, v := range names {
		initials = append(initials, v[0:1])
		fmt.Println(v)
	}
	if len(initials) > 1 {
		return initials[0], initials[1]
	}
	return initials[0], "_"
}

func main() {
	cycleNames([]string{"Ara", "Razmig", "Lori"}, sayGreeting)
	cycleNames([]string{"Ara", "Razmig", "Lori"}, sayBye)
	circleArea := circleAre(10)
	fmt.Printf("Circle area is %0.3f\n", circleArea)
	fmt.Println(getInitials("Ara Chanoian"))
	fmt.Println(getInitials("Lori"))
	first, last := getInitials("Ara Chanoian")
	fmt.Println(first, last)
}
