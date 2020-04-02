// https://tour.golang.org/methods/4
// Read about modules https://github.com/golang/go/wiki/Modules,
package main

import (
	"fmt"
)

type Player struct {
	Name string
	Age  int
}

func main() {

	// var m map[string]Player
	m := make(map[string]Player)

	m["Martin"] = Player{"Martin JI Brandl", 34}
	_, ok := m["Martin2"]

	fmt.Println("present:", ok)
	fmt.Println(m["Martin"])

}
