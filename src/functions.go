package main

import "fmt"

func add(x float32, y float32) float32 {
	return float32(x + y)
}

func main() {
	fmt.Println(add(4214.3, 123))
}
