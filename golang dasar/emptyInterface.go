package main

import "fmt"

func ups(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "ups"
	}
}

func main() {
	var data interface{} = ups(1)
	// var data interface{} = ups(2)
	// var data interface{} = ups(3)
	fmt.Println(data)
}
