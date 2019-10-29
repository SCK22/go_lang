package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x)) // this function is calling itself with -x if x < 0
}

func main() {
	fmt.Println(sqrt(2))
	fmt.Println(sqrt(-4))
}
