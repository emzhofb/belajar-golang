package main

import (
	"errors"
	"fmt"
)

func divide(n int, v int) (int, error) {
	if v == 0 {
		return 0, errors.New("divider is zero")
	} else {
		return n / v, nil
	}
}

func main() {
	example, err := divide(10, 5)
	example2, err2 := divide(10, 0)

	if err == nil {
		fmt.Println(example)
	} else {
		fmt.Println(err)
	}

	if err2 == nil {
		fmt.Println(example2)
	} else {
		fmt.Println(err2)
	}
}
