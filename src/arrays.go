//  a [10]int means array a consists of 10 values of the type int
//  Arrays length is a part of its definitions, so the size cannot be changed
package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])

	primes := [6]int{2, 3, 5, 7, 11, 13} // initialize and assign values
	fmt.Println(primes)
}
