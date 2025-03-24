// package main

// import (
// 	"fmt"
// 	"time"
// )

// type Car struct {
// 	Model string
// 	Speed int
// }

// // Функция для проверки четности числа
// func isEven(n int) bool {
// 	return n%2 == 0
// }

// // Функция для проверки четности числа
// func multiply(a int, b int) int {
// 	return a * b
// }

// func sayHello() {
// 	fmt.Println("Hello from goroutine")
// }

// func Sender(ch chan int) {
// 	time.Sleep(3 * time.Second) // Имитируем обработку полученного числа
// 	fmt.Println(time.Now().Format("15:04:05.000"), "Горутина отправителя запущена")
// 	for i := range 5 {
// 		ch <- i
// 		fmt.Println(time.Now().Format("15:04:05.000"), "Отправили число:", i)
// 		// time.Sleep(500 * time.Millisecond) // Имитируем задержку перед отправкой следующего числа
// 	}
// 	close(ch) // Закрытие канала после отправки всех чисел
// 	fmt.Println(time.Now().Format("15:04:05.000"), "Горутина отправителя завершена")
// }

// func Receiver(ch chan int) {
// 	time.Sleep(300 * time.Millisecond) // Имитируем обработку полученного числа
// 	fmt.Println(time.Now().Format("15:04:05.000"), "Горутина получателя запущена")
// 	for num := range ch {
// 		fmt.Println(time.Now().Format("15:04:05.000"), "Получено число:", num)
// 		time.Sleep(300 * time.Millisecond) // Имитируем обработку полученного числа
// 	}
// 	fmt.Println(time.Now().Format("15:04:05.000"), "Горутина получателя завершена")
// }

// func main() {
// 	fmt.Println("Hello, World!")

// 	// Проверка четности чисел от 0 до 4
// 	for i := range 5 {
// 		if isEven(i) {
// 			fmt.Println(i, ": четное число")
// 		} else {
// 			fmt.Printf("%d: нечетное число\n", i)
// 		}
// 	}

// 	// Функция для умножения двух чисел
// 	a, b := 3, 4
// 	result := multiply(a, b)
// 	fmt.Printf("Результат умножения %d * %d = %d\n", a, b, result)

// 	// Создание структуры и вывод информации о ней
// 	mycar := Car{"Tesla", 320}
// 	fmt.Printf("Марка машины: %s, скорость: %d км/ч\n", mycar.Model, mycar.Speed)

// 	// Создание слайса строк и добавление элемента
// 	names := []string{"John", "Mary", "Tom"}
// 	names = append(names, "Bob")
// 	fmt.Println(names)

// 	// Создание мапы городов и населения
// 	cities := map[string]int{
// 		"New York":    1000000,
// 		"Los Angeles": 800000,
// 		"Chicago":     500000,
// 		"Moscow":      12000000,
// 	}
// 	fmt.Println("Население Москвы", cities["Moscow"], "людей")

// 	// Использование горутины для выполнения функции в отдельной горутине
// 	go sayHello()
// 	time.Sleep(1 * time.Second)
// 	// fmt.Println("Main done")

// 	// Создание буфферизованного канала с размером буфера 2
// 	ch := make(chan int, 3)
// 	// Запуск горутин отправителя
// 	go Sender(ch)

// 	// Запуск как функции получателя
// 	Receiver(ch)

// 	fmt.Println("Программа завершена")
// }
