/*
Для решения данной задачи вам понадобится пакет strconv, возможно использовать пакеты strings или encoding/csv, или даже bufio
- вы не ограничены в выборе способа решения задачи.
В решениях мы поделимся своими способами решения этой задачи, предлагаем вам сделать то же самое.

В привычных нам редакторах электронных таблиц присутствует удобное представление числа с разделителем разрядов в виде пробела,
кроме того в России целая часть от дробной отделяется запятой.
Набор таких чисел был экспортирован в формат CSV, где в качестве разделителя используется символ ";".

На стандартный ввод вы получаете 2 таких вещественных числа, в качестве результата требуется вывести частное от деления первого числа на второе
с точностью до четырех знаков после "запятой" (на самом деле после точки, результат не требуется приводить к исходному формату).

P.S. небольшое отступление, связанное с чтением из стандартного ввода. Кто-то захочет использовать для этого пакет bufio.Reader.
Это вполне допустимый вариант, но если вы резонно обрабатываете ошибку метода ReadString('\n'), то получаете ошибку EOF, а точнее (io.EOF - End Of File).
На самом деле это не ошибка, а состояние, означающее, что файл (а os.Stdin является файлом) прочитан до конца.
Чтобы ошибка была обработана правильно, вы можете поступить так:

	if err != nil && err != io.EOF {
		...
	}

Sample Input:
1 149,6088607594936;1 179,0666666666666

Sample Output:
0.9750
*/
package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	text = strings.ReplaceAll(text, " ", "")

	r := csv.NewReader(strings.NewReader(text))
	r.Comma = ';'
	// r.Comment = '#'

	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	a := strings.ReplaceAll(records[0][0], ",", ".")
	a64, err := strconv.ParseFloat(a, 64)
	if err != nil {
		fmt.Println("Error converting string to float64:", err)
		return
	}
	b := strings.ReplaceAll(records[0][1], ",", ".")
	b64, err := strconv.ParseFloat(b, 64)
	if err != nil {
		fmt.Println("Error converting string to float64:", err)
		return
	}
	res := fmt.Sprintf("%.4f", a64/b64)
	fmt.Println(res)

}
