// A struct is a collection of fields.
package main

import "fmt"

type Vertex struct {
	X int
	Y int
	// Z int
}

func main() {
	fmt.Println(Vertex{1, 2})
}
