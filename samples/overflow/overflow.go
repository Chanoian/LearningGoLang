pa eackage main

import "fmt"

func main() {
	var b byte = 255
	var smallI int32 = 2147483647
	var bigI uint64 = 18446744073709551615
	b = b + 1                    // This will overflow and wrap around to 0
	smallI = smallI + 1          // This will overflow and wrap around to -2147483648
	bigI = bigI + 1              // This will overflow and wrap around
	fmt.Println(b, smallI, bigI) // Output: 0 -2147483648 0
}
