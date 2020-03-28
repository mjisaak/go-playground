// https://tour.golang.org/moretypes/15
// Read about modules https://github.com/golang/go/wiki/Modules,
package main

import (
	"fmt"
)

func main() {

	var s []int
	var y []int

	s = append(s, 3)
	s = append(s, 5)

	y = append(s, 1)

	fmt.Println("s: ", s)
	fmt.Println("y: ", y)
}
