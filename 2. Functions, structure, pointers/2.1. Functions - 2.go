/*
Напишите функцию, находящую наименьшее из четырех введённых в этой же функции чисел.

Входные данные
Вводится четыре числа.

Выходные данные
Необходимо вернуть из функции наименьшее из 4-х данных чисел

Sample Input:
4 5 6 7

Sample Output:
4
*/
package main

import "fmt"

func minimumFromFour() int {
	//print your code here
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)
	return min(min(a, b), min(c, d))
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	fmt.Println(minimumFromFour())
}
