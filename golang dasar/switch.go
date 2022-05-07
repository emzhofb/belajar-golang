package main

import (
	"fmt"
)

func main() {
	name := "Karin"

	switch name {
	case "Karin":
		fmt.Println("Karin-san, daisuki!")
	case "Sekai":
		fmt.Println("Ohayou, sekai!")
	default:
		fmt.Println("Hello, World!")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Terlalu panjang")
	case false:
		fmt.Println("Sudah benar")
	}

	length2 := len(name)
	switch {
	case length2 > 10:
		fmt.Println("Sangat Panjang")
	case length2 > 5:
		fmt.Println("Lumayan Panjang")
	default:
		fmt.Println("Sudah Pas")
	}
}
