package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func changeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func changeCountryToJapan(address Address) {
	address.Country = "Japan"
}

func main() {
	var alamat = Address{
		City:     "Saitama",
		Province: "Tokyo",
		Country:  "",
	}

	changeCountryToIndonesia(&alamat)
	fmt.Println(alamat)

	var address = Address{
		City:     "Jepara",
		Province: "Jawa Tengah",
		Country:  "",
	}

	changeCountryToJapan(address)
	fmt.Println(address)

	var pointerAlamat *Address = &alamat
	changeCountryToIndonesia(pointerAlamat)
	fmt.Println(pointerAlamat)
}
