package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		fmt.Println("Hey " + os.Args[1])
		return
	}
	fmt.Println("Hey")
}
