package main

import "fmt"

func formatter(message string) string {
	formatted := fmt.Sprintf("%v.", message)
	return formatted
}

func reformat(message string, formatter func(string) string) string {
	firstResult := formatter(message)
	secondResult := formatter(firstResult)
	thirdResult := formatter(secondResult)
	finalResult := fmt.Sprintf("TEXTIO: %v", thirdResult)
	return finalResult
}

func main() {
	reformattedString := reformat("Hello", formatter)
	fmt.Println(reformattedString)
}
