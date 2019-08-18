//  struct fields can be accessed using a dot "."
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	fmt.Println("v before changing value :", v)
	v.X = 4
	fmt.Println("v after changing value : ", v)
}
