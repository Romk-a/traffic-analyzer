package main

import "fmt"

type Car struct {
	Model string
	Speed int
}

// Функция для проверки четности числа
func isEven(n int) bool {
	return n%2 == 0
}

// Функция для проверки четности числа
func multiply(a int, b int) int {
	return a * b
}

func main() {
	fmt.Println("Hello, World!")

	// Проверка четности чисел от 0 до 4
	for i := range 5 {
		if isEven(i) {
			fmt.Println(i, ": четное число")
		} else {
			fmt.Printf("%d: нечетное число\n", i)
		}
	}

	a, b := 3, 4
	result := multiply(a, b)
	fmt.Printf("Результат умножения %d * %d = %d\n", a, b, result)

	mycar := Car{"Tesla", 320}
	fmt.Printf("Марка машины: %s, скорость: %d км/ч\n", mycar.Model, mycar.Speed)
}
