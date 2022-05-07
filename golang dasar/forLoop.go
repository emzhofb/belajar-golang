package main

import (
	"fmt"
)

func main() {
	counter := 1
	for counter <= 10 {
		fmt.Println("Hello,", counter)
		counter++
	}

	for count := 1; count <= 10; count++ {
		fmt.Println("Count,", count)
	}

	slice := []string{"Ikhda", "Muhammad", "Wildani", "Fujiyoshi", "Karin"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(i+1, "=>", slice[i])
	}

	names := []string{"Asuka", "Saito", "Fujiyoshi", "Karin"}

	for index, name := range names {
		fmt.Println("index", index, "=", name)
	}

	// use _ for unused variable
	for _, name := range names {
		fmt.Println(name)
	}

	person := make(map[string]string)
	person["name"] = "Karin"
	person["title"] = "Kawaii"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}
