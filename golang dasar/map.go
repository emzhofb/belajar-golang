package main

import (
	"fmt"
)

func main() {
	person := map[string]string{
		"name":    "karin",
		"address": "tokyo",
	}

	person["title"] = "idol"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	fmt.Println(len(person))

	var book map[string]string = make(map[string]string)
	book["title"] = "Ngidol Jepang"
	book["author"] = "Ikhda"
	book["oshi"] = "Karin"

	fmt.Println(book)
	fmt.Println(len(book))
	delete(book, "oshi")
	fmt.Println(book)
	fmt.Println(len(book))
}
