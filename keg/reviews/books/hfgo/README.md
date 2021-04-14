---
Title: Head First Go
Subtitle: Annotations
Query: true
---

Even though *Head First Go* is *very* out of date it is currently the best Go programming language book for beginners. It is available for [purchase](https://www.amazon.com/Head-First-Go-Jay-McGavren/dp/1491969555) and online as PDF for free (although it is difficult to find and illegal in most countries).

* [Install Golang](/tools/editors/vim/plugins/vimgo/)
* [Get Vim Ready for Go Programming](/tools/editors/vim/plugins/vimgo/)
* [Chapter One: Let's Get Going: Syntax Basics](#ch01)

:::co-faq

## Where can I get the PDF of the book?

Online. But using a PDF without having paid for the book is illegal in most countries (and some would argue even *with* having purchased the book). Please do not ask anyone from the rwx.gg team to provide you with a PDF. If you want one you'll need to find one on your own.

## Can I use Goland?

Sure. Goland is an incredible graphic development environment for Go that can be particularly useful for very large Go projects. In general graphic IDEs can make development tasks very efficient (such as renaming functions throughout the entire code base). However, these annotations assume you are using Vim have mastered Bash enough to do these same important tasks from the command line instead.

## Can I use Visual Studio Code (VSCode)?

Using anything but `vim` from the `bash` command line is strongly discouraged, but plenty of people do.  You can certainly use VSCode (or any editor for that matter) if you really want to but no time is spent explaining how to set it up so please *do not ask how to configure VSCode during live sessions*. 

:::

## Chapter One: Let's Get Going: Syntax Basics {#ch01}

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

