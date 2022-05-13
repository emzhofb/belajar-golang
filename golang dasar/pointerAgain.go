package main

import "fmt"

type Address struct {
	City     string
	Province string
	Country  string
}

func main() {
	var address1 Address = Address{"Saitama", "Tokyo", "Japan"}
	var address2 *Address = &address1
	var address3 *Address = &address1

	address2.City = "Shinjuku"
	*address3 = Address{"Fujiyama", "Kyoto", "Japan"}
	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	var address4 *Address = new(Address)
	address4.City = "Fukushiyama"
	address4.Province = "Osaka"
	address4.Country = "Japan"
	fmt.Println(address4)
}
