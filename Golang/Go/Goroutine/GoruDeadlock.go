package main

func main() {
	messages := make(chan string)

	
	go func() {}()

	// But no groutine runnning or No sender
	
	// So it raises deadlock error
	messages <- "I wanna tell you." 
