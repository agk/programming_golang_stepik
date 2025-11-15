/*
Требуется написать программу, при выполнении которой с клавиатуры считываются два натуральных числа A и B (каждое не более 100, A < B). Вывести сумму всех чисел от A до B  включительно.
Sample Input:
1 5
Sample Output:
15
*/
package main

import "fmt"

func main() {
	var a, b, sum int
	//fmt.Print("Введите число: ")
	fmt.Scan(&a, &b)
	sum = 0
	for a > 100 && b < 100 && a < b {
		sum += a
		a++
	}
	fmt.Println(sum)
}
