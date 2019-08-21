//The zero value of a map is nil. A nil map has no keys, nor can keys be added.
//  for key value pairs
package main

import "fmt"

type Vertex struct {
	Lat     float64
	Long    float64
	Address string
}

func main() {
	m := make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		Lat: 40.68433, Long: -74.39967, Address: "Cupertino",
	}
	m["Koinearth"] = Vertex{0.000, 0.000, "HSR"}
	fmt.Println(m)
	// fmt.Println(m["Bell Labs"])
}
