// https://tour.golang.org/moretypes/15
// Read about modules https://github.com/golang/go/wiki/Modules,
package main

import (
	"fmt"
)

type Player struct {
	fartingPower int
	name         string
}

func main() {

	players := []Player{
		Player{fartingPower: 3, name: "Martin"},
	}

	fmt.Println("yo", players)
}
