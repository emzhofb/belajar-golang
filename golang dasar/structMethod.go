package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello() {
	fmt.Println("Konnichiwa, my name is", customer.Name)
}

func (customer Customer) sayHi(name string) {
	fmt.Println("Konnichiwa", name, ",my name is", customer.Name)
}

func main() {
	karin := Customer{
		Name: "Fujiyoshi Karin",
	}

	karin.sayHello()
	karin.sayHi("Ikhda")
}
