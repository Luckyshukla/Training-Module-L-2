package main

import (
	"fmt"
)

func handlepanic() {

	// if a := recover(); a != nil {
	
	// 	fmt.Println("RECOVER", a)
	// }
}


func entry(lang *string, aname *string) {

	// Deferred function
	defer handlepanic()

	if lang == nil {
	
		panic("Error: Language cannot be nil")
	}
	
	// is nil it will panic
	if aname == nil {
		panic("Error: Author name cannot be nil")
	}
	fmt.Printf("Author Language: %s \n Author Name: %s\n", *lang, *aname)
	fmt.Printf("Return successfully from the entry function")
}

// Main function
func main() {

	A_lang := "GO Language"
	entry(&A_lang, nil)
	fmt.Printf("Return successfully from the main function")
}
