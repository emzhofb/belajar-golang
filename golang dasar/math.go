package main

import "fmt"

func main() {
	var a = 10
	var b = 5
	var c = a + b
	fmt.Println(c)
	c *= 4
	c = c / 2
	fmt.Println(c)
	c++
	fmt.Println(c)
}
