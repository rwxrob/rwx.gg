---
Title: Don't `cat` Into a Shell Pipe
Subtitle: No Need for Extra Subprocess
Query: true
---

```sh
# WRONG!
cat foo | somecmd something`
```

This is just redundant, wasteful, and usually very stupid. Cat creates a subprocess unnecessarily. This is even worse when used to [feed a `while` loop](../pipe2while) because it hides variables.

Use redirection instead.

```sh
somecmd something < foo
```
