/*
Ваша задача сделать проверку подходит ли пароль вводимый пользователем под заданные требования.
Длина пароля должна быть не менее 5 символов, он должен содержать только арабские цифры и
буквы латинского алфавита. На вход подается строка-пароль.

Если пароль соответствует требованиям - вывести "Ok", иначе вывести "Wrong password"

Sample Input:
fdsghdfgjsdDD1
Sample Output:
Ok
*/
package main

import (
	"fmt"
	"unicode"
)

func checkPass(password string) bool {
	if len(password) < 5 {
		return false
	}

	for _, v := range []rune(password) {
		if unicode.IsDigit(v) || unicode.Is(unicode.Latin, v) {
			continue
		}
		return false
	}

	return true

}

func main() {
	var password string
	fmt.Scan(&password)

	if !checkPass(password) {
		fmt.Println("Wrong password")
	} else {
		fmt.Println("Ok")
	}
}
