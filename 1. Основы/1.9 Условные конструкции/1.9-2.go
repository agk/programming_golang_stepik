/*
По данному трехзначному числу определите, все ли его цифры различны.
Формат входных данных
На вход подается одно натуральное трехзначное число.
Формат выходных данных
Выведите "YES", если все цифры числа различны, в противном случае - "NO".
Sample Input:
237
Sample Output:
YES
*/
package main

import "fmt"

func main() {
	var d, a, b, c int
	fmt.Scan(&d)

	a = d / 100
	b = d / 10 % 10
	c = d % 10
	if a != b && b != c && c != a {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
