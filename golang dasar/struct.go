package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var karin Customer
	karin.Name = "Fujiyoshi Karin"
	karin.Age = 22
	karin.Address = "Japan"

	fmt.Println(karin)
	fmt.Println(karin.Name)

	ikhda := Customer{
		Name:    "Ikhda M Wildani",
		Address: "Indonesia",
		Age:     23,
	}

	fmt.Println(ikhda)
}
