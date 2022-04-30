package main

import "fmt"

func main() {
	type NoKTP string
	type married bool

	var ikhda NoKTP = "1234567890"
	var isMarried married = false
	fmt.Println(ikhda)
	fmt.Println(isMarried)
}
