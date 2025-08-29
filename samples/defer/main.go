package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	client := http.Client{
		Timeout: 5 * time.Second,
	}
	// this will run in the end of any return in the function ,
	defer fmt.Println("Closing connection")

	fmt.Println("Connecting to google.com ...")
	resp, err := client.Get("https://swww.google.com")
	if err != nil {
		fmt.Println("Error while connecting to Google")
		return
	}
	if resp.StatusCode == http.StatusOK {
		fmt.Println("Connection Succesful !!")
	}
}
