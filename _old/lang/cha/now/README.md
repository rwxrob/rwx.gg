---
Title: Now Command Challenge
Subtitle: Create a Human-Friendly Time Stamp
Tags: easy
---

Create a program called `now` that prints a specific human-friendly time
stamp suitable for including in an article or a blog (Monday, June 29,
2020, 8:28:26PM). Create *another* program called `hnow` that calls the
first program and adds a default Markdown header prefix `## `. Have
`hnow` accept one argument and use that instead of the default if set.
Demonstrate using `hnow` from `vi` by combining it with `!!` correctly. 

## Requirements

* Must include a *safe* shebang line if a script.
* When `./now` then prints `Monday, June 29, 2020, 8:28:26PM`.
* When `./hnow ` then prints `## Monday, June 29, 2020, 8:28:26PM`.
* When `./hnow //` then prints `// Monday, June 29, 2020, 8:28:26PM`.
* When `./hnow '### '` then prints `### Monday, June 29, 2020, 8:28:26PM`.
* Demonstrate how to use within Vi/m combined with `!!`.

## Bonus

* Validate using language test framework.
