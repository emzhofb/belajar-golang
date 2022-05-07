package main

import "fmt"

func main() {
	counter := 0
	name := "Karin"

	increment := func() {
		fmt.Println("increment")
		counter++
		name := "Saito"
		fmt.Println(name)
	}

	increment()
	fmt.Println(counter)
	fmt.Println(name)
}
