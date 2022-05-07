package main

import "fmt"

func getFullName() (string, string) {
	return "Fujiyoshi", "Karin"
}

func main() {
	firstName, lastName := getFullName()
	fmt.Println("Konnichiwa", firstName, lastName)

	familyName, _ := getFullName()
	fmt.Println(familyName, "san daisuki!")
}
