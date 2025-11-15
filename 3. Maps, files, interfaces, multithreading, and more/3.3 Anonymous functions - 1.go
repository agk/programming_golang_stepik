// Анонимные функции
package main

import "fmt"

func printMessage(message string) {
	fmt.Println(message)
}

func getPrintMessage() func(string) {
	// Возвращаем анонимную функцию
	return func(message string) {
		fmt.Println(message)
	}
}

func main() {
	// Именованная функция
	printMessage("Hello function!")

	// Анонимная фукция объявляется и вызывается
	func(message string) {
		fmt.Println(message)
	}("Hello anonymous function!")

	// Получаем анонимную функцию и вызываем ее
	printfunc := getPrintMessage()
	printfunc("Hello anonymous function using caller!")

}
