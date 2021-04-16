---
Title: Bash Scripting Language
Subtitle: The Default Interactive Linux Shell
Query: true
---

> `/bin/bash` is premature optimization.
> [\@beginbot](https://twitch.tv/beginbot)

*Bourne Again Shell* is default interactive [user](/what/users/)
[shell](/what/shell/) of the Linux operating system. It is also a
powerful [scripting language](/lang/) known for creating useful scripts
complex single-line commands quickly. As such it has a rather unusual
[syntax](/what/syntax/) when compared to other languages. Bash usage and
programming are a [primary focus](/reviews/books/tlcl/) of the [Linux
Beginner Boost](/boosts/oldboost/schedule/).

## Bash Shell Survival {#survive}

Whether you are just starting out as a new technologist or have been one
for a while but not really learned the terminal command line shell then
you might find yourself stuck or not knowing what to do. Here's a
minimal survival guide to help you out or get you started. Remember as
you master these that there is a *lot* more to come. Don't worry about
understanding every, single option with every command at the moment or
even why they do what they do. That will come in time. 

:::co-stop
Before you get started make sure you understand all the [keyboard
symbols](/key/) and just know that you can usually
combine two options that use a [dash](/key/) like `ls
-al`.
:::

:::co-warn
If you happen to know (or have had someone else tell you about) the `ls`
aliases like `ll` just don't use them for now. It is far more important
for you to learn the actual commands, not the alias shortcuts. Later
after you have *really* mastered the *actual* command and options
consider adding your own aliases.
:::

### Cozy Shell Commands {#cozy}

First the *cozy* commands, those that make you feel at home and like you
know where you are, how to get around, what stuff is, and what's inside.

------------------ ------------------------------------------
 Command            Description
------------------ ------------------------------------------
 `ls`              List *visible* files and directories in 
                   the current directory.

 `ls -a`           List *all* files and directories in 
                   the current directory.

 `ls -l`           List the details of files and directories
                   including permissions.

 `pwd`             Print the working directory (where
                   you are).

 `clear`           Clear the screen. [*Do not use
                   `Control`+`L`!*](/advice/dont/controll)

 `cd`              Change into the home directory from
                   anywhere.

 `cd ..`           Change parent of current directory. The
                   two dots `..` mean parent.

 `cd -`            Change into the previous directory, which
                   is nice for changing back and forth a lot.


 `cat <file>`      Shows the content of the file.

 `less <file>`     Shows the content of a long file that
                   you can page through with `Space`. 
                   Remember `q` to quit.

 `type <this>`     Check if this command exists, what
                   *type* of command it is (alias, function,
                   binary executable), and where it lives.
------------------ ------------------------------------------

### Existential Shell Commands {#existential}

Now for the *existential* commands that create, change, and delete
stuff. You don't need to fear these commands, just be more careful when
using them. There is *no* undo on the Linux command line (until you get
Git setup).

-------------------------- ------------------------------------------
 Command                   Description
-------------------------- ------------------------------------------
`file <file>`              Show what kind of file it is.

`mkdir <dir>`              Make a new directory.

`touch <file>`             Create a new empty file. (Or update the
                           time stamp of the file.)

`mv <old> <new>`           Rename a file or directory.

`rm <file>`                Remove a file. 

`rmdir <dir>`              Remove an *empty* directory. (Use `rm`
                           first for the files until it is empty.)

`chmod u+x <file>`         Change a plain text file into an
                           *executable* script command.
                   
`./<script>`               Run a newly made command script in the
                           current directory.
-------------------------- ------------------------------------------

:::co-mad

There is a *reason* dangerous advanced commands like `rm -rf <dir>` and
brain-dead things like `Control`-`L` to clear your screen are not
included in this survival shell section. It is designed for beginners
who might seriously mess themselves up knowing these options exist at
this time. Eventually everyone should learn the more advanced options,
just not right now. 

By the way, if you do *not* understand why `Control`-`L` is so bad to
burn into your muscle memory you simply have not yet learned *basic*
command line Linux usage at all. Keep reading and researching. You'll
get there.

:::

:::co-pwz

Don't forget to know all the [extra key combinations and how to get
yourself unstuck](/key/).

:::

## Omissions

There are several valuable and important elements of Bash that the book leaves out. Make sure you understand them.

* `select`
* `until`
* `getopts`
* `complete`
* `CDPATH`
* ANSI Color and Terminal Escape Sequences
* Parameter expansion for default values

## See Also

* [Google Bash Style](https://google.github.io/styleguide/shellguide.html)
- <http://tiswww.case.edu/php/chet/bash/FAQ>
