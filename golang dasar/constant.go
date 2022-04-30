package main

import "fmt"

func main() {
	// variable constant cannot be replaced, and need to have initial value
	const firstName string = "Ikhda"
	const lastName = "Wildani"

	fmt.Println(firstName)

	const (
		kawaii string = "Asuka Saito"
		oshi          = "Fujiyoshi Karin"
	)
}
