package main

import "fmt"

func main() {
	messages := make(chan string)

	go func() {
		fmt.Println("Try to receive message") // Printing
		<-messages                            // Blocking
		fmt.Println("Receive message")        // Never reached
	}()

	fmt.Println("Try to receive message") // Printing
	<-messages                            // Blocking
	fmt.Println("Receive message")        // Never reached

}