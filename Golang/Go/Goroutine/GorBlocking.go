
package main

import "fmt"
import "time"

func main() {
	messages := make(chan string)
		time.Sleep(1000 * time.Millisecond)
		messages <- "Hello"
	}()

	fmt.Println(<-messages) // Hello
}