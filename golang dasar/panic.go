package main

import "fmt"

func endApp() {
	message := recover()
	if message != nil {
		fmt.Println("something went wrong", message)
	}

	fmt.Println("aplikasi selesai")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("aplikasi error")
	}
	fmt.Println("aplikasi berjalan")
}

func main() {
	// runApp(false)
	runApp(true)
}
