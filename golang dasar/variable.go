package main

import "fmt"

func main() {
	// use var with data type declared
	var name string

	name = "Ikhda Muhammad"
	fmt.Println(name)

	name = "Ikhda Wildani"
	fmt.Println(name)

	// use var and no declared
	var oshi = "Fujiyoshi Karin"
	fmt.Println(oshi)

	// use := instead of var
	kawaii := "Asuka Saito"
	fmt.Println(kawaii)

	kawaii = "Saito Asuka"

	// define multiple variable
	var (
		firstName = "Ikhda"
		lastName  = "Wildani"
	)
	fmt.Println(firstName, lastName)
}
