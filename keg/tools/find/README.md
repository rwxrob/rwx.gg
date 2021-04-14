---
Title: Find Command
Subtitle: The Most Powerful Command Some Have Ever Heard Of
Query: true
---

The `find` command may just be the most powerful and useful command you can run from the command line (that isn't an entire language itself like `bash`, `perl`, or `awk`). It has two main uses (1) to find stuff (duh) and (2) to act on the things that it finds. The magic of `find` is in it's ability to automatically walk through an entire directory tree by descending down into it. Such traversal is not a trivial coding exercise. (You should write one someday to try it.) Although `find` isn't a language it definitely has a syntax. Learning all the switches, their order, and predicates is not unlike learning a small language itself.

:::co-rant
It is surprising how many veteran UNIX/Linux users --- even administrators --- don't even know about the ultra-powerful `find` command. Don't be them.
:::

## Advantages Over Shell Glob Expansion

Shell glob expansion is when you use one or two stars `*` to expand out to all the possible matching files and directories. This usually works but has one *major* disadvantage. The number of hits that it expands to are limited and unless activated might not even expand at all. Even though `find` costs a subprocess to use, it avoids these risks entirely and is fully POSIX compliant as well.

## Examples

The `man` page is rather large and hard to digest. Here are some examples of common tasks you'll likely need.

:::co-tip
Don't forget that some of these will need to be run as `sudo`. Otherwise you want to redirect all the permission to read errors by adding `2>/dev/null` to the end.
:::

### Find All Files 34 Bytes in Size

```sh
find / -size 34c
```

### Find All Symbolic Links

```sh
find / -type l
```

### Find Everything in `/home` or `/var` Owned by Root

```sh
find /home /var -owner root
```

### Find All Files on System with Setuid or Setgid and Long List

```sh
find / -perm /=s -ls
```

### Grep All Files in the Current Directory and Children with a Given Suffix

```sh
find . -name '*.md' -exec grep blah {} /dev/null \;
```

## See Also

* [Search About Find](https://duck.com/lite?kae=t&q=find%20linux%20command)
* `man find`
