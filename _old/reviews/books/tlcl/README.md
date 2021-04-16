---
Title: The Linux Command Line, 5th Edition
Subtitle: Best Book for Linux Beginners
Query: true
---

[*The Linux Command Line*](./tlcl.pdf) is the work of William Shotts and a number of other contributors from <http://linuxcommand.org>. Shotts' pleasant writing style and long experience provide the best possible book for any Linux command line beginner. In fact, it is currently the *only* Linux book available that covers Bash 4+ including such things as associative arrays. This puts TLCL well beyond established books such as *Learning Bash* (O'Reilly) which has grown far too out of date to be of any value.

TLCL does, however, contain several glaring flaws that can be summarized best by the fact that it covers `nroff`. This set of annotations is therefore intended to make up for these flaws without dismissing the book entirely, which unfortunately cannot be fixed because of another *major* flaw: it's license, which prevents fixes due to its "no derivative works" clause defeating the main reason for creating open content in the first place.

:::co-fyi
Eventually TLCL will be replaced with this [knowledge app](/what/knowledge/apps/) from the RWX community that will eventually cover everything in TLCL without qualifying as a "derivative work" and allow full forks without penalty. The annotations below are *not* a derivative work and are no different than anyone else's public notes taken and shared while reading the book.
:::

## Introduction

Make sure to read all of the Introduction. It provides an excellent answer to *why* you should learn the command line.

## Chapter 24: Writing Your First Script {#chapter-24}

Finally we get to write some shell code. 

:::co-warn
Don't confuse the term [shell code from pentesting](https://duck.com/lite?kae=t&q=shell code from pentesting) with *shell script*. 
:::

### Bash is *Usually* the Default Command Line

Bash has been the default shell for Linux for several decades but was replaced as the *system startup* scripting language by Dash on most Linux distributions some time ago. But Bash remains the default interactive shell assigned to new users.

#### What about Zsh?

Don't use it.

#### What about Dash?

`/bin/sh` is [symbolically linked](https://duck.com/lite?kae=t&q=symbolically linked) `/bin/dash`, a light-weight, [POSIX-compliant](https://duck.com/lite?q=POSIX-compliant) shell that runs much more quickly for use in the Linux startup process. 

Despite the claim on Dash's home page that it is *the* Linux shell. The [Arch distro](/tools/linux/distros/#arch) symlinks to `/bin/bash` instead.

### Bash History

Bash code is unlike most other coding languages because it evolved from Bourne Shell which was released in 1979. 

Bourne *Again* Shell has added stuff from lots of shells since that time including primarily Korn Shell, which came out in 1983. The first version of Bash was released in 1989. And the latest significant Bash release (version 5.0) came out on [January, 2019](https://lwn.net/Articles/776223/). More significantly, however, was the release of Bash 4.0, which came out in February, 2009 and included [Bash associative arrays](https://duck.com/lite?kae=t&q=Bash associative arrays), which are covered by the 5th edition of this book.

:::co-warning
This book is the *only* book available anywhere that covers Bash 4+ associative arrays, which is why it was picked even over Learning Bash from O'Reilly.
:::

### Bash is Interpreted

This chapter jumps right into how to use Bash but it is worth understanding that Bash is an [interpreted language](/lang/) and that the [shebang](https://duck.com/lite?kae=t&q=shebang) line tells the [operating system](https://duck.com/lite?q=operating system) to send the script file to the `/bin/bash` binary command which interprets it into [system calls](https://duck.com/lite?q=system calls) that `bash` itself executes. There is no intermediary [bytecode](https://duck.com/lite?q=bytecode) created.

### `PATH` Environment Variable

These days `~/.local/bin` is the preferred place to hide your local `bin` directory --- usually by symlinking to a directory in your personal [dotfiles config](https://duck.com/lite?kae=t&q=dotfiles config). 

You might consider adding the following script to your `~/.local/bin/` directory or adding the echo line as an alias in `~/.bashrc` file in order to simplify reading your PATH.

Script Version

```bash
#!/bin/bash
echo -e ${PATH//:/\\n}
```

Bash Alias

```bash
alias path="echo -e ${PATH//:/\\n}"
```
