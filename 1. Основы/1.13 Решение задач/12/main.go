/*
По данному числу N распечатайте все целые значения степени двойки, не превосходящие N, в порядке возрастания.

Входные данные
Вводится натуральное число.

Выходные данные
Выведите ответ на задачу.

Sample Input:
50

Sample Output:
1 2 4 8 16 32

*/
package main

import "fmt"

func main() {
	var n, d int
	fmt.Scan(&n)
	//n=1
	d = 1
	for i := 0; i < n; i++ {

		if d <= n {
			fmt.Print(d, " ")
		} else {
			break
		}

		d *= 2
	}

}
