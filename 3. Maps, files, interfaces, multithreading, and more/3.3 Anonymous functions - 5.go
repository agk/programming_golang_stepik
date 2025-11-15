/*
Используем анонимные функции на практике.

Внутри функции main (объявлять ее не нужно) вы должны объявить функцию вида func(uint) uint,
которую внутри функции main в дальнейшем можно будет вызвать по имени fn.

Функция на вход получает целое положительное число (uint), возвращает число того же типа
в котором отсутствуют нечетные цифры и цифра 0.
Если же получившееся число равно 0, то возвращается число 100.
Цифры в возвращаемом числе имеют тот же порядок, что и в исходном числе.

Пакет main объявлять не нужно!
Вводить и выводить что-либо не нужно!
Уже импортированы - "fmt" и "strconv", другие пакеты импортировать нельзя.

Sample Input:
727178

Sample Output:
28
*/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num uint
	fmt.Scan(&num)

	fn := func(x uint) uint {
		if x == uint(0) {
			return uint(100)
		}
		newX := []rune(strconv.FormatUint(uint64(x), 10))
		// fmt.Printf("Type - %T, Num - %v\n", newX, newX)

		var strX string

		for _, value := range newX {
			if value%2 == 0 && value != '0' && value != ' ' {
				strX = strX + string(value)
			}
		}

		// fmt.Printf("Type - %T, Num - %c\n", strX, strX)

		if len(strX) == 0 {
			// fmt.Printf("000 Type - %T, Num - %c\n", strX, strX)
			return uint(100)
		}

		u64, err := strconv.ParseInt(strX, 10, 64)
		if err != nil {
			panic(err)
		}
		// fmt.Printf("Type - %T, Num - %v\n", u64, u64)

		myUint := uint(u64)
		// fmt.Printf("Type - %T, Num - %v\n", myUint, myUint)

		return myUint

	}

	// num1, _ := strconv.Atoi(num)
	// fmt.Print(a == 0) // true

	// x := uint(727178)
	x := uint(num)
	fmt.Println(fn(x))
}
