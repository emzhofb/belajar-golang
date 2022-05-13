package main

import "fmt"

type Woman struct {
	Name string
}

func (woman Woman) Married() {
	woman.Name = "Mrs. " + woman.Name
}

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	karin := Man{"Karin"}
	karin.Married()

	fmt.Println(karin.Name)

	fujiyoshi := Woman{"Fujiyoshi"}
	fujiyoshi.Married()

	fmt.Println(fujiyoshi.Name)
}
