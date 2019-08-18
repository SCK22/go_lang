/*
A pointer holds the memory address of a value.package src

The type *T is a pointer to a T value. It's zero value is nil


The & operator generates a pointer to its operand.

i := 42
p = &i

The * operator denotes the pointer's underlying value.

fmt.Println(*p) // read i through the pointer p
*p = 21         // set i through the pointer p

This is known as "dereferencing" or "indirecting".

Unlike C, Go has no pointer arithmetic.

*/
package main

import "fmt"

func main() {
	i, j := 42, 2701
	fmt.Println("Original value of i:", i)

	p := &i // point to i
	fmt.Println("Pointer p value : ", *p)
	*p = 21                           // set i through the pointer
	fmt.Println("New value of i:", i) // set the new value of i

	fmt.Println("Original value of j:", j)
	p = &j       //point to j now
	*p = *p / 37 // divide j through the pointer
	fmt.Println("Pointer p value : ", *p)
	fmt.Println("New value of j:", j) // see the new value of j

}
