---
Title: Learn to Code Functions from Math Functions
---

You might not remember functions from Math but if you do there's a good
chance you will better understand [computer language functions](https://duck.com/lite?kae=t&q=computer language functions) and
[procedures](https://duck.com/lite?kae=t&q=procedures). 

Let's start with the basics but first you need to understand the
[computer programming language math operators](https://duck.com/lite?kae=t&q=computer programming language math operators). Here's a quick summary:

|Operator|Math|
|-|-|
|`+`|Add|
|`-`|Subtract|
|`*`|Multiply|
|`/`|Divide|
|`%`|Remainder|
|`^`|Exponent|

Not all languages use the same operators but these are pretty common.

You can test these out using `bc` from the command line.

## F of X FTW!

Remember that whole "f of x" `f(x)` stuff? Here's a simple example
(using programming operators):

    f(x) = 2*x + 3
    f(3) = 2*3 + 3
    f(3) = 6 + 3
    f(3) = 9 

    f(x,y) = 2*x + 3*y +3
    f(2,6) = 2*2 + 3*6 +3
    f(2,6) = 4 + 18 +3
    f(2,6) = 25

In the above the `x` and `y` are *parameters* which are like slots that
you put things into as if a machine. Parameters have names.

The numbers `3`, `2`, and `6` are *arguments* which to into the
parameter slots of the function machine. Arguments are just values.

## Moving from Math to Code

Using the same idea we can do things that are not Math related, for
example print a greeting.

    f(name) = print "Hi there, $name."
    f(Pete) = print "Hi there, Pete."

Now that's not *actually* how you would write that, but you get the
idea. Here's the actual code in four different languages. All of the
following output 'Hey there, Pete.'

Bash

```bash
hey () {
  name="$1"
  echo "Hey there, ${name}."
}
hey Pete
```

Python

```python
def hey(name): print(f'Hey there, {name}.')
hey("Pete")

```

JavaScript

```js
const hey = name => console.log(`Hey there, ${name}.`)
hey("Pete")
```

Golang

```go
import "fmt"
func hey(name string) {
  fmt.Printf("Hey there, %v.\n", name)
}
hey("Pete")
```

Even though we have been calling the above a *function* it is really
more of a *procedure* because it is a set of instructions and not really
a machine the spits out things when we put in things.

## How About Math Functions in the Same Languages

Remember that Match function we started out with? 

    f(x,y) = 2*x + 3*y +3

Let's code that in the same four languages. All of these print `The
result is 25`.

Bash

```bash
foo () {
  x="$1"
  y="$2"
  echo $(( 2*x + 3*y ))
}
echo "The result is $(foo 2 6)"
```


