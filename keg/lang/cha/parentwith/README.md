---
Title: ParentWith Challenge 
Tags: moderate
---

Create a function that checks for the existence of a given file or
directory by name beginning in a given starting directory and if not
found checks all parent directories recursively. Ignore anything that
begins with dot (.) or underscore (\_).

## Requirements

* Must include a *safe* shebang line if a script.
* Write the algorithm first as comments using natural language.
* Use functional recursion.
* Name function `ParentWith()` or something similar.
* Accept a required name parameter.
* Allow *optional* starting directory parameter.
* Assume current directory if no starting dir.
* Return the absolute path to the item found.
* Return empty string if no item found.
* Examine *any* filesystem entity (files, dirs, fifos, etc.).
* Ignore anything beginning with dot (\.).
* Ignore anything beginning with underscore (\_).

## Bonus

* Validate using language test framework.

