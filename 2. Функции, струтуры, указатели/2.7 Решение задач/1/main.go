/*
На вход подаются a и b - катеты прямоугольного треугольника. Нужно найти длину гипотенузы

Sample Input:
6 8
Sample Output:
10
*/
package main

import (
	"fmt"
	"math"
)

func hypo(k1 float64, k2 float64) float64 {
	return math.Sqrt(k1*k1 + k2*k2)
}

func main() {
	var a, b float64
	fmt.Scan(&a, &b)
	res := hypo(a, b)
	fmt.Println(res)
}
