package main

import "fmt"

func factorialLoop(n int) int {
	result := 1
	for i := n; i > 0; i-- {
		result *= i
	}

	return result
}

func factorialRecursive(n int) int {
	if n == 1 {
		return 1
	}

	return n * factorialRecursive(n-1)
}

func main() {
	resultLoop := factorialLoop(7)
	fmt.Println(resultLoop)
	fmt.Println(7 * 6 * 5 * 4 * 3 * 2 * 1)

	resultRecursive := factorialRecursive(5)
	fmt.Println(resultRecursive)
	fmt.Println(5 * 4 * 3 * 2 * 1)
}
