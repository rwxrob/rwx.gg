---
Title: RGB Command Challenge
Tags: moderate
---

Create a command that takes three arguments for red, green, and blue
that are numbers between 0 and 255. Combine the numbers into a vt100
(ANSI) RGB color escape and print the escape sequence invisibly coloring
anything printed after. Test by printing different things in different
colors. Once that works go back and assign random default color values
for each if less than three arguments are passed.

## Requirements

* Must include a safe shebang line if a script.
* Use raw escape sequences.
* Do not use `tput` in any form. 
* When using `echo` do not use `-e`.
* When `./rgb` then print random color

## Hints

* Research [printing in color on the terminal](https://duck.com/lite?kae=t&q=printing in color on the terminal)
* You will need to escape escape (didn't stutter).
* Research how to generate a random number in your language.

## Bonus

* Validate using language test framework.
* Check that each value is between 0 and 255 or print usage.
* Write a `for` loop to cycle through gradient output to the terminal.
