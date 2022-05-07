package main

import "fmt"

func getHello(name string) string {
	if name == "" {
		return "Dare da teme?"
	}

	return "Konnichiwa " + name + "-san"
}

func main() {
	result := getHello("Karin")
	fmt.Println(result)
	fmt.Println(getHello(""))
}
