/*
Defer statement defers the execution of a funciton until the surrounding
function returns.
The deferred call's arguments are evaluated immediately but
the funtions call is not executed until the surrounding function returns
*/
package main

import "fmt"

func main() {
	defer fmt.Println("World")

	fmt.Println("hello")
}
