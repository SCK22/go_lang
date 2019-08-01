package main

import "fmt"

/* named return values, these can be treated as
variable defined at the top of the function
a return statement without arguments returns names values
 also called a naked return , use only in small functions
*/
func split(sum int) (x, y, z int) {
	x = sum * 2 / 5
	y = sum - x
	z = x + y
	return
}

func main() {
	fmt.Println(split(5))
}
