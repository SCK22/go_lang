/* Variable declared without an explicit initial
value will be given 0 initial value */

/*
The zero value is:
0 for numeric types
 false for boolean type and
 "" for string type
*/

package main

import "fmt"

func main() {
	var (
		i int
		f float64
		b bool
		s string
	)
	fmt.Printf("%T %T %T %T\n", i, f, b, s)
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
