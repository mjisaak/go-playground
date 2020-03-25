// https://tour.golang.org/moretypes/13
// Read about modules https://github.com/golang/go/wiki/Modules,
package main

import (
	"fmt"
)

func main() {

	s := make([]int, 3)

	fmt.Println(s)
	ar := []int{1, 2, 3, 4, 5}
	printSize(ar[:3])
}

func printSize(array []int) {
	fmt.Printf("Capacitiy=%d; Length=%d", cap(array), len(array))
}
