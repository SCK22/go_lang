package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value: ", v, "Present?", ok)

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok = m["Answer"]
	fmt.Println("The value: ", v, "Present?", ok)
	// delete does not throw an error if the key is not present
	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])
}
