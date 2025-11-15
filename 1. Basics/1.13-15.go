/*
Из натурального числа удалить заданную цифру.

Входные данные
Вводятся натуральное число и цифра, которую нужно удалить.

Выходные данные
Вывести число без заданных цифр.

Sample Input:
38012732
3

Sample Output:
801272

*/
package main

import "fmt"

func IntToSliceExceptNum(n int64, num int64, sequence []int64) []int64 {
	if n != 0 {
		i := n % 10
		if i != num {
			// sequence = append(sequence, i) // reverse order output
			sequence = append([]int64{i}, sequence...)

		}
		return IntToSliceExceptNum(n/10, num, sequence)
	}
	return sequence
}

func main() {
	var n, d int64
	fmt.Scan(&n)
	fmt.Scan(&d)
	//fmt.Println(d)

	sequence := make([]int64, 0, 0)
	sequence = IntToSliceExceptNum(n, d, sequence)
	//fmt.Print(sequence)
	for _, v := range sequence {
		fmt.Print(v, "")

	}
}
