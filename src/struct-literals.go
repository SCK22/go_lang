/*
Struct Literals
A struct literal denotes a newly allocated struct value by listing the values of its fields.

You can list just a subset of fields by using the Name: syntax. (And the order of named fields is irrelevant.)

The special prefix & returns a pointer to the struct value
*/

package main

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2} // has type Vertex
	v2 = Vertex{X: 1} // Y is implicitly 0
	v3 = Vertex{}     // X , Y are 0
	p  = &Vertex{1, 2}
)

func main() {

}
