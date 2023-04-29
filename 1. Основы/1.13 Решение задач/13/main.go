/*
Номер числа Фибоначчи
Дано натуральное число A > 1. Определите, каким по счету числом Фибоначчи оно является, то есть выведите такое число n, что φn=A.
Если А не является числом Фибоначчи, выведите число -1.

Входные данные
Вводится натуральное число.

Выходные данные
Выведите ответ на задачу.

Sample Input:
8

Sample Output:
6

*/
package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	a := 0
	b := 1
	return func() int {
		a, b = b, a+b
		return b - a
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	resultCheck := false
	//n=9
	f := fibonacci()
	for i := 0; i <= n+5; i++ {
		res := f()
		if res == n {
			fmt.Print(i)
			resultCheck = true
			break
		}
		//fmt.Println(i, n, res)
	}

	if !resultCheck {
		fmt.Print(-1)
	}

}
