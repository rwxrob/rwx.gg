package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println(math.Floor(2.75))
	fmt.Println(strings.Title("head first go"))
	a := '🙂'
	fmt.Printf("The letter '%c': %b\n", a, a)
}
