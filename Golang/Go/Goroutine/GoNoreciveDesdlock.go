package main

func main() {
	messages := make(chan string)

	// Do nothing spawned goroutine
	go func() {}()

	// A groutinetrying to receive message from channel

	// So channel will be deadlocking
	<-messages // fatal error: