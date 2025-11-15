// Замыкание и состояние
package main

import "fmt"

func counter(start int) (func() int, func()) {
	// если значение мутирует, то мы получим изменение и в этом замыкании
	ctr := func() int {
		return start
	}

	incr := func() {
		start++
	}

	// и в ctr, и в incr сохраняется указатель на start
	// мы создали замыкания но еще не вызывали
	return ctr, incr
}

func main() {
	// ctr, incr и ctr1, incr1 различаются своим состоянием
	ctr, incr := counter(100)
	ctr1, incr1 := counter(100)
	fmt.Println("counter - ", ctr())
	fmt.Println("counter1 - ", ctr1())
	// увеличиваем на 1
	incr()
	fmt.Println("counter - ", ctr())
	fmt.Println("counter1- ", ctr1())
	// увеличиваем до 2
	incr1()
	incr1()
	fmt.Println("counter - ", ctr())
	fmt.Println("counter1- ", ctr1())
}
