/*
Количество нулей
По данным числам, определите количество чисел, которые равны нулю.

Входные данные
Вводится натуральное число N, а затем N чисел.

Выходные данные
Выведите количество чисел, которые равны нулю.

Sample Input:
5
1
8
100
0
12

Sample Output:
1

*/
package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	cnt := 0
	for i := 0; i < n; i++ {
		var v int
		fmt.Scan(&v)
		if v == 0 {
			cnt++
		}
	}
	fmt.Println(cnt)
}
