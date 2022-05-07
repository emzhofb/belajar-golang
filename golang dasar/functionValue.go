package main

import "fmt"

func getGoodBye(name string) string {
	return "Sayonara " + name + "-san"
}

func main() {
	goodbye := getGoodBye
	fmt.Println(goodbye("Karin"))
	fmt.Println(getGoodBye("Saito"))
}
