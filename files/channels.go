package files

import "fmt"

func Task23() {
	fmt.Println()
	messages := make(chan string)

	go func() { messages <- "ping" }()

	msg := <-messages
	fmt.Println(msg)

}
