package main

import "fmt"

/* named return values, these can be treated as
variable defined at the top of the function
a return statement without arguments returns named values
 also called a naked return , use only in small functions

syntax : 
	func name(parameters) (return values)

 */
func split(blah int) (x, y, z, p int) {
	fmt.Println(blah)
	x = blah * 2 / 5
	y = blah - x
	z = x + y
	p = x-y
	
	return
}

func main() {
	fmt.Println(split(5))
}
