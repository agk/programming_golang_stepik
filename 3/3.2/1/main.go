/*
Напишите функцию с именем convert, которая конвертирует входное число типа int64 в uint16.
Считывать и выводить ничего не нужно!
Считайте что функция main и пакет main уже объявлены!
Sample Input:
Sample Output:
*/
package main

import "fmt"

func convert(d int64) uint16 {
	return uint16(d)
}

func main() {
	var n int64
	fmt.Scan(&n)
	fmt.Println(convert(n))

}
