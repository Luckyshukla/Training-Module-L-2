package main

import "fmt"

type Address struct {
	Name string
	city string
	Pincode int
}

func main() {

	a2 := Address{Name: "Ram", city: "Delhi", Pincode: 2400}
	a1 := Address{"Pam", "Dehradun", 2200}
	a3 := Address{Name: "Sam", city: "Lucknow", Pincode: 1070}

	var mp map[Address]int
	
	
	if mp == nil {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}

	sample := map[Address]int{a1: 1, a2: 2, a3: 3}
	fmt.Println(sample)
}
