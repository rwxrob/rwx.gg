---
Title:  Eloquent JavaScript
Subtitle: Free, Comprehensive, But Heavy on Math for Beginners
Query: true
---

*Eloquent JavaScript* is available for [free as a PDF](./eloqjs.pdf) or you can purchase it online. It was crowd funded by a group of people concerned there was a good modern resource covering topics for beginners. Here's how to get it from the command line with `curl`:

```sh
curl -LO https://rwx.gg/reviews/books/eloqjs/eloqjs.pdf 
```

*These annotations are incomplete at this time.*

## Introduction

* What what your favorite part?

## Chapter 1

## Chapter 2

## Stuff to Find a Home For (in the Above)

### Dangers of `Number()`

In the following `Number()` is *not* a *cast* (if you know what that is). It doesn't force the string into being a number. Instead, it returns `NaN` if its argument is not a number. This then has to be checked to be sure that the string entered in response to the `prompt()` is actual a number and not something else.

### For Loops with `in` Versus `of`

The old `in` operator iterates over the *properties* (usually the names) and not the *values*. This makes `in` do things that seem unexpected and is the reason `of` was added to ES6 (and that you should generally use it instead).

```js
let list = [4, 5, 6]

for (let i in list) {
   console.log(i) # NOT 4, 5, 6
}

for (let i of list) {
   console.log(i)
}
```

```{.out}
0
1
2
4
5
6
```

Here's one that is even more unexpected:

```js
let devices = new Set(["Linux", "Mac", "Windows"]);
devices["category"] = "desktops"

for (let device in devices) {
   console.log(device)
}

for (let device of devices) {
    console.log(device)
}

```

```{.out}
category
Linux
Mac
Windows
```

