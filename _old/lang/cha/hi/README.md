---
Title: '"Hi You" Challenge'
Tags: moderate
---

Create a program named `hi` that prints `Hi there.` by default but print
`Hi <arg>.` if there is a command line argument. If a one of two special
names (like `Linux` and `Rob`) are passed then instead of the default
the script should print something special like `Woah <arg>, you rock!`
If a specific insult name (like `Dork`) is used the script should print
`Um, no need to be rude.`

## Requirements

* Must include a *safe* shebang line if a script.
* When `./hi` then prints `Hi there.`
* When an argument is passed `./hi <name>` then prints `Hi <name>.`
* When `./hi Rob` then prints `Woah Rob, you rock!` only.
* When `./hi Linux` then prints `Woah Linux, you rock!` only.
* When `./hi Dork` then prints `Um, no need to be rude.` only.
* Do not read anything from standard input.
* Use safest, most efficient design, and best practices.
* Be ready to defend your design under peer review.

## Bonus

* Write the algorithm first as comments using natural language.
* Validate using language test framework.
* Code it again using `switch/case`.
* Code it again using `else if / elif` variations.
* Code it again *without any* `else` statements.
* Code it again using a function map design.
* Code it again using short-circuit logical operators.
