package main

import "fmt"

func getCompleteName() (firstName, middleName, lastName string) {
	return "Ikhda", "Muhammad", "Wildani"
}

func getCompleteName2() (firstName, middleName, lastName string) {
	firstName = "Monkey"
	middleName = "D."
	lastName = "Luffy"

	return
}

func main() {
	firstName, middleName, lastName := getCompleteName()
	fmt.Println(firstName, middleName, lastName)

	firstName2, middleName2, lastName2 := getCompleteName2()
	fmt.Println(firstName2, middleName2, lastName2)
}
