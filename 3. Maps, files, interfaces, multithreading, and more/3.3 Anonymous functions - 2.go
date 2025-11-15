// Замыкания
package main

import "fmt"

func outer(name string) {
	// переменная из внешней функции
	text := "Modified " + name

	// foo это внутренняя функция и из нее есть доступ к переменной `text`
	// у замыкания есть доступ к этим переменным даже после выхода из блока
	foo := func() {
		fmt.Println(text)
	}

	// вызываем замыкание
	foo()
}

func main() {
	outer("hello")
}
