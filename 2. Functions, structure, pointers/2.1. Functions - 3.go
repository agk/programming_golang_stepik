/*
Напишите "функцию голосования", возвращающую то значение (0 или 1),
которое среди значений ее аргументов x, y, z встречается чаще.

Входные данные
Вводится 3 числа - x, y и z (x, y и z равны 0 или 1).

Выходные данные
Необходимо вернуть значение функции от x, y и z.

Sample Input:
0 0 1
Sample Output:
0
*/
package main

import "fmt"

func vote(x int, y int, z int) int {
	//print your code here
	cntZero := 0
	cntOne := 0
	if x != 0 {
		cntOne++
	} else {
		cntZero++
	}
	if y != 0 {
		cntOne++
	} else {
		cntZero++
	}

	if z != 0 {
		cntOne++
	} else {
		cntZero++
	}

	if cntOne > cntZero {
		return 1
	}
	return 0
}

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	fmt.Println(vote(a, b, c))
}
