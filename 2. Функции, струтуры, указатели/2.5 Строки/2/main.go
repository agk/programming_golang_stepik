/*
На вход подается строка. Нужно определить, является ли она палиндромом.
Если строка является палиндромом - вывести Палиндром иначе - вывести Нет. Все входные строчки в нижнем регистре.

Палиндром — буквосочетание, слово или текст, одинаково читающееся в обоих направлениях
(например, "топот", "око", "заказ").

Sample Input:
топот
Sample Output:
Палиндром
*/
package main

import (
	"fmt"
)

func main() {
	var word, reverseWord string
	fmt.Scan(&word)

	for _, v := range word {
		reverseWord = string(v) + reverseWord
	}

	if reverseWord == word {
		fmt.Println("Палиндром")
	} else {
		fmt.Println("Нет")
	}

}
