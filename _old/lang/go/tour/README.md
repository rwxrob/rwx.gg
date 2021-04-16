---
Title: Tour of Go, A Workthrough
Subtitle: Helping Beginners Get Through It
---

This is a [workthrough](/what/workthrough/) of the [Tour of
Go](https://tour.golang.org), which, for better or worse, is what
[most](https://go.dev) recommend people do to get started in the
language. The Tour is clearly *not* for beginner programmers and assumes
coding in at least one other language (preferably C) before attempting
it. It's pedantic tone can be off-putting for many just coming to
programming for the first time. This is an attempt to help assuage their
their distress so they can come to realize Go *is* fun for beginners as
well.

Video Playlist: **[Tour of Go, A
Workthrough](https://www.youtube.com/playlist?list=PLrK9UeDMcQLpSguoPyoroznF-rf7ewMbX)**

## Incomplete

The Tour is actually not nearly comprehensive. It is missing a lot of
things that remain on the project [TODO
list](https://github.com/golang/tour/blob/master/TODO). Given the fact
that the Tour *requires* a JavaScript enabled browser, and that it is not
really designed for beginners, this will probably remain unfinished for
the foreseeable future. Check the [Go resources](/lang/go/) to find
other learning material to consider, but be careful, there is a *lot* of
outdated stuff that looks like it will be there for a very long time.

:::co-rant
It is *really* time for the community to come together and create a
*true* beginner overview full of examples that is maintained with the
latest practices. This is a particularly important goal given the
soon-to-be release of generics which will require even more explanation.
:::

## Exercise: Loops and Functions

At first you just 

```go
func Sqrt(x float64) float64 {
	z := 1.0
	for i := 1; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
	}
	return z
}
```

Then you can keep track of the difference:

```go
func Sqrt(x float64) float64 {
	last, z := 1.0, 1.0
	for i := 1; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		delta := z - last
		fmt.Printf("z=%v delta=%v\n", z, delta)
		last = z
	}
	return z
}
```

And finally remove the limit of 10 and checking for a specific delta
(difference) from the last value.

```go
package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	last, delta, z := 1.0, 1.0, 1.0
	for delta >1e-6{
		z -= (z*z - x) / (2 * z)
		delta = math.Abs(last -z)
		//fmt.Printf("z=%v delta=%v\n", z, delta)
		last = z
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}
```

It looks like Go must use the value `1e-6` internally within the
`math.Sqrt()` function.
