package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
		fmt.Println(i)
	}
	fmt.Println(sum)

	names := []string{"ara", "razmig", "gerard", "lori"}
	for i, v := range names {
		fmt.Printf("The Value at index is %v is %v \n", i, v)
	}
}
