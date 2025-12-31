package main

import (
	"fmt"
)

// This is example of embedded struct
type car struct {
	brand   string
	mileage int
}

type truck struct {
	car
	bedSize int
}

// This is example of nested struct
type train struct {
	brand      string
	frontWheel wheel
}

type wheel struct {
	radius int
}

func main() {
	myTruck := truck{
		bedSize: 10,
		car: car{
			brand:   "Toyota",
			mileage: 100000,
		},
	}
	fmt.Println(myTruck.mileage) // in embedded struct you can use the mileage from the parent struct without using myTruck.car.mileage
	fmt.Println(myTruck.bedSize)

	myTrain := train{
		brand: "Volvo",
		frontWheel: wheel{
			radius: 10,
		},
	}
	fmt.Println(myTrain.brand)
	fmt.Println(myTrain.frontWheel.radius) //you can access it within the frontWheel struct because it's nested

}
