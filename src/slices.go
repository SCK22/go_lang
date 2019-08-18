// This selects a half-open range which includes the first element, but excludes the last one.
package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4] // create emply array s, put a slice of primes into s
	fmt.Println(s)
}
