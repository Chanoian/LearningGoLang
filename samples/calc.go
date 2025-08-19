package main

import "fmt"

func main() {
    var sum int
    fmt.Scanln(&sum)
    x := (sum * 4) / 9
    fmt.Printf("Result: %d\n", x)
}
