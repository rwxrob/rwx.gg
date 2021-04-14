# Review of *GoBootcamp* E-Book

GoBootcamp is a nice attempt to quickly bring a person up to speed on
the Go programming language but suffers from several *major* flaws
related to how out of date it has become --- most significantly, it has
no information on `go mod init`, the now standard way to begin a
project.

## Chapter 2.8

> I personally recommend against named return parameters.

This is the author's opinion, but it is rather short-sighted. Named
parameters might be slightly more cognitive work for the coder or
maintainer but they are much more understandable to the person reading
the function signature to determine what it does and returns.

> Go has pointers, but no pointer arithmetic. 

That is categorically false. Pointer arithmetic is *strongly*
discouraged but very possible even through almost never needed.

```go
package main

import "fmt"
import "unsafe"

func main() {
    vals := []int{10, 20, 30, 40}
    start := unsafe.Pointer(&vals[0])
    size := unsafe.Sizeof(int(0))
    for i := 0; i < len(vals); i++ {
        item := *(*int)(unsafe.Pointer(uintptr(start) + size*uintptr(i)))
        fmt.Println(item)
    }
}
```

