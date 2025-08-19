package main

import (
	"fmt"
	"sort"
)

func main() {
	ages := []int{44, 32, 21, 92}
	sort.Ints(ages)
	fmt.Println(ages) // replace the original slice
	index := sort.SearchInts(ages, 44)
	fmt.Println(index)
}
