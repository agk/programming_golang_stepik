/*
Дана строка, содержащая только арабские цифры. Найти и вывести наибольшую цифру.

# Входные данные

Вводится строка ненулевой длины. Известно также, что длина строки не превышает 1000 знаков и
строка содержит только арабские цифры.

# Выходные данные

Выведите максимальную цифру, которая встречается во введенной строке.

Sample Input:
1112221112
Sample Output:
2
*/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s string
	fmt.Scan(&s)
	/*
	   d, err := strconv.Atoi(s)
	   if err != nil {
	       panic(err)
	   }
	*/
	res := 0
	for i := 0; i < len(s); i++ {
		d, err := strconv.Atoi(string(s[i]))
		if err != nil {
			panic(err)
		}
		if res < d {
			res = d
		}

	}
	fmt.Print(res)

}
