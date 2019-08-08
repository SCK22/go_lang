package main

import (
	"fmt"
)

func netwonsSqrt(x float64) float64 {
	z_start := x / 4.0
	counter := 0
	for counter < 5 {
		fmt.Printf("Present value of z is %f\n", z_start)
		z_start -= (z_start*z_start - x) / (2 * z_start)
		z_loop := z_start
		if z_loop-z_start == 0.0 {
			counter += 1
		}
		fmt.Println(z_loop)
	}
	return z_start
}

func main() {
	fmt.Println(netwonsSqrt(25))
}
