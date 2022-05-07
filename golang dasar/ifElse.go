package main

import (
	"fmt"
)

func main() {
	name := "Sekai"

	if name == "Sekai" {
		fmt.Println("Ohayou, Sekai!")
	} else if name == "Karin" {
		fmt.Println("Karin-san, daisuki!")
	} else {
		fmt.Println("Sekai ga zankoku da, sore demo kimi wo aisu yo!")
	}

	if length := len(name); length > 5 {
		fmt.Println("Terlalu panjang")
	} else {
		fmt.Println("Sudah benar")
	}
}
