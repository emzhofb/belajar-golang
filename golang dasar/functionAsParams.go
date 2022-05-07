package main

import "fmt"

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Konnichiwa", nameFiltered)
}

func spamFilter(name string) string {
	if name == "anjing" {
		return "..."
	}

	return name
}

func main() {
	sayHelloWithFilter("karin", spamFilter)
	sayHelloWithFilter("anjing", spamFilter)
}
