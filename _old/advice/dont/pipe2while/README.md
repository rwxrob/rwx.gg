---
Title: Don't Pipe to a Shell While Loop
Subtitle: You Lose All Your Variables
Query: true
---

This runs while in a subprocess masking the local variables (`count`) so they are not updated.

```sh
# WRONG!
myvar="a value"
declare -i count=0
grep something somefile | while read line; do
  echo "$myvar: $line"
  count+=1
done
echo "Count: $count"
```

```{.out}
a value: Here is something
a value: And something else
a value: One last something
Count: 0
```

Use redirection instead, either file descriptor or temporary file, (which both benchmark at about the same speed).

```sh
myvar="a value"
declare -i count=0
while read -u3 line; do
  echo "$myvar: $line"
  count+=1
# creates a tmp file descriptor and "pipes" in
done 333< <(grep something somefile)
echo "Count: $count"
```

The `-u3` and `3<` are to prevent the input from buffering and to
prevent anything within the body of the loop that needs to read from
standard input (file descriptor 0) from causing problems.

```{.out}
a value: Here is something
a value: And something else
a value: One last something
Count: 3
```


```sh
myvar="a value"
declare -i count=0
while read line; do
  echo "$myvar: $line"
  count+=1
# creates a tmp file with the grep output 
done <<< $(grep something somefile)
echo "Count: $count"
```

```{.out}
a value: Here is something
a value: And something else
a value: One last something
Count: 3
```
