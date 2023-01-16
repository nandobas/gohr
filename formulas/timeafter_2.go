package formulas

import (
	"fmt"
	"time"
)

func sendMessageInChannel(c chan string) {
	time.Sleep(6 * time.Second)
	c <- "Its timeout 6 sec.."
}

func TimeAfterOutput2() {
	// Creating a channel
	// Using make keyword
	channel := make(chan string, 2)

	go sendMessageInChannel(channel)
	// Select statement
	select {

	// Using case statement to receive
	// or send operation on channel
	case output := <-channel:
		fmt.Printf("saida exibida quando recebe info no channel:%s\n", output)

	// parameter
	case <-time.After(5 * time.Second):

		// Printed after 5 seconds
		fmt.Println("Its timeout..")
	}
	time.Sleep(10 * time.Second)
}
