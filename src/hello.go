// https://tour.golang.org/moretypes/1

package main

import (
	"fmt"
)

type Player struct {
	Name string
	Age  int
}

func main() {

	vi := []struct {
		Name string
		Age  int
	}{
		{Name: "Martin", Age: 34},
		{Name: "Chris", Age: 33},
	}

	fmt.Println("Hello everybody. This is Player:", vi[0].Name)

	fmt.Println("And he is together with", vi[1].Name)

}
