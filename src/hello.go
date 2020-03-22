// https://tour.golang.org/moretypes/1

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Counting:")

	for i := 0; i < 10; i++ {
		defer fmt.Println("", i)
	}
}
