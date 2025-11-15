package main

import (
	"fmt"
	"strconv"
)

func main() {
	// res := []rune("stepik")
	// res := string([]byte("stepik"))
	// a := 10.123	res := int64(a)
	// res := 101.0 / 10
	// res := strconv.Itoa(int(float64(100/10)) + 1)
	// res := strconv.FormatBool(10 == int16(float64(100/10)))
	// res := (strconv.FormatBool(true)) == (10 == int(float64(100/10)))
	res := strconv.FormatBool(10.1 == float32(float64(100/10.1)))

	fmt.Println(res)

}
