---
Title: 'Head First Go: Chapter One'
Subtitle: "Let's Get Going: Syntax Basics (Annotations)"
Query: true
---

Hello world is hello world. Nothing fancy to see there.

Use the Go community convention of naming the file with the `func main()` in it `main.go` so you can find it easier.

Here's a version that introduces emojis as *runes* as well as the `fmt.Printf()` function:

```go
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
```

