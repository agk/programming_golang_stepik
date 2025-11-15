/*
Количество минимумов
Найдите количество минимальных элементов в последовательности.

Входные данные
Вводится натуральное число N, а затем N целых чисел последовательности.

Выходные данные
Выведите количество минимальных элементов последовательности.

Sample Input:
3
21
11
4

Sample Output:
1

*/
package main

import "fmt"

func main() {
	var n, m, min int
	fmt.Scan(&n)
	num_min := 1
	for i := 0; i < n; i++ {
		fmt.Scan(&m)
		if i == 0 {
			min = m
		}
		//fmt.Println(m, min, num_min)
		if m < min {
			min = m
			num_min = 1
		} else if m == min && i != 0 {
			num_min++
		}
	}
	fmt.Println(num_min)
}
