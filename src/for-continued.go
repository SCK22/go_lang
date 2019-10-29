//  the init and post statements are optional
//  also for is go's while loop

package main

import "fmt"

func main() {
	sum := 1
	for sum < 1000 { // linter does not allow me to put ; before and after the condition statement
		sum += sum
	}
	fmt.Println(sum)
}

//  forever loop

// func main() {
// 	for {
// 	}
// }
