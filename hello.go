package main

import "fmt"

func isEven(n int) bool {
	return n%2 == 0
}
func multiply(a int, b int) int {
	return a * b
}

func main() {
	fmt.Println("Hello, World!")

	for i := range 5 {
		if isEven(i) {
			fmt.Println(i, ": четное число")
		} else {
			fmt.Printf("%d: нечетное число\n", i)
		}
	}
}
