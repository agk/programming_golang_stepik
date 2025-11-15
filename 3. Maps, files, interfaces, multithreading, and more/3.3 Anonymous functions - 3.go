// Возвращаем замыкание и используем его снаружи
package main

import "fmt"

func outer(name string) func() {
	// переменная
	text := "Modified " + name

	// замыкание. у функции есть доступ к переменной text даже
	// после выхода за пределы блока
	foo := func() {
		fmt.Println(text)
	}

	// возвращаем замыкание
	return foo
}

func main() {
	// foo это наше замыкание
	foo := outer("hello")

	// вызов замыкания
	foo()
}
