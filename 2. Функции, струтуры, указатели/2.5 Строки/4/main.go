/*
На вход дается строка, из нее нужно сделать другую строку, оставив только нечетные символы (считая с нуля)

Sample Input:
ihgewlqlkot
Sample Output:
hello
*/
package main

import "fmt"

func main() {
	var s, res string
	fmt.Scan(&s)
	res = ""
	for i, v := range s {
		if i%2 == 0 {
			continue
		}
		res = res + string(v)
	}
	fmt.Print(res)

}
