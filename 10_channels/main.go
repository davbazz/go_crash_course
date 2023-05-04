package main

import "fmt"

func main() {
	messages := make(chan string)

	go func() { messages <- "ping" }()

	msg := <-messages

	fmt.Println(msg)

	bufferedChan := make(chan string, 2)

	bufferedChan <- "new entry"
	bufferedChan <- "another entry"

	fmt.Println(<-bufferedChan)
	fmt.Println(<-bufferedChan)

}
