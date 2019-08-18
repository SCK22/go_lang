//The zero value of a map is nil. A nil map has no keys, nor can keys be added.

package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

func main() {
	m := make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m)
	fmt.Println(m["Bell Labs"])
}
