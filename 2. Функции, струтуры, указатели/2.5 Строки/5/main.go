/*
Дается строка. Нужно удалить все символы, которые встречаются более одного раза и вывести получившуюся строку

Sample Input:
zaabcbd
Sample Output:
zcd
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)
	res := ""
	for _, v := range s {
		if strings.Count(s, string(v)) > 1 {
			continue
		}
		res = res + string(v)
	}
	fmt.Println(res)
}
