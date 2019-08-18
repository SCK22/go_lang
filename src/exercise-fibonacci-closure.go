package main

import "fmt"

func fibonacci() func() int {
	x := 0
	y := 1
	z := 0
	return func() int {
		print(x, y, z)
		z, x, y = x, y, x+y
		return z
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
