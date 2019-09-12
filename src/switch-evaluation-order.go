package main

import (
	"fmt"
	"time"
)

func main() {

	today := time.Now().Weekday()
	// fmt.Print(reflect.TypeOf(today))
	fmt.Printf("Today : %s \n", today)
	fmt.Println("When's Saturday?")
	fmt.Println(time.Saturday)
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("More than two days away.")
	}
}
