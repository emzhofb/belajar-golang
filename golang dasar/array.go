package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "ikhda"
	names[1] = "karin"
	names[2] = "asuka saito"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{
		30,
		90,
		100,
	}
	fmt.Println(values)

	fmt.Println(len(names), len(values))

	var empty [10]string
	fmt.Println(len(empty))
}
