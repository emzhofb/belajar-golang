package main

import "fmt"

type Address struct {
	City     string
	Province string
	Country  string
}

func main() {
	// pass by value
	address1 := Address{"Jepara", "Jawa Tengah", "Indonesia"}
	address2 := address1

	address2.City = "Semarang"

	fmt.Println(address1) // {Jepara Jawa Tengah Indonesia}
	fmt.Println(address2) // {Semarang Jawa Tengah Indonesia}

	// pass by reference
	address3 := Address{"Saitama", "Tokyo", "Japan"}
	address4 := &address3

	address4.City = "Shinjuku"
	fmt.Println(address3)  // {Shinjuku Tokyo Japan}
	fmt.Println(address4)  // &{Shinjuku Tokyo Japan}
	fmt.Println(*address4) // {Shinjuku Tokyo Japan}

	address4 = &Address{"Kansai", "Kyoto", "Japan"}
	fmt.Println(address4) // &{Kansai Kyoto Japan}
}
