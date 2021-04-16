---
Title: Hello Bash! Challenge
Subtitle: Create a Hello World Script in Bash
Category: Challenge
Summary:
  Type a command into the command line to print out `Hello World!`. Now place that same command into a Bash file named `hello` and get it to run as if you had typed it in. Then get it to run using the `bash` interpreter. Now get it to run *without* the interpreter on the command line so that typing `./hello` from the command line makes it work. Replace `World` with a Bash variable declared as `name` assigned the string `World`. Play around with changing the value. Now make `./hello First` print `Hello First!` and `./hello 'First Last'` print `Hello First Last!` (but `./hello First Last` should only print `Hello First!`).

Prereqs:
- /tools/linux/openterm/
- /tools/linux/mkdir/
- Make a Bash Script
- Print Output with Bash
- Declare Variable in Bash
- Assign a Value to Variable in Bash
- Comment Out a Line in Bash
- Comment Out Multiple Lines in Bash
- Use Command Arguments in Bash
- Set a Default Variable Value with Parameter Expansion in Bash
- Check Condition with If Statement in Bash
- Lookup Bash `test` Conditions
- Check for Empty String in Bash
- Check for Full String in Bash
---

First you'll need to have a Linux Bash terminal open to work in and you should create a [codebook](/what/codebook/) Git project repo to contain you coding notes that you can refer to later.

Get on your terminal and change into your codebook directory and create a subdirectory for your Bash learning projects.

```sh
mkdir bash
```

Change into your `bash` directory.

```sh
cd bash
```

Now we need to make a file to contain the code.

```sh
touch hello
```

Let's add something to the file.

Open the file with the `vi` editor.

```sh
vi hello
```

The command to print something to the terminal screen is `echo`. You can place quotes around the stuff to print or not and it will guess pretty well. If you use punctuation you *do* need to have quotes, so why not.You can use single or double quotes. We'll go over the difference later.


```sh
echo 'Hello world!'
```

Now save the file with `:w` without closing it.

Open another terminal and navigate into the same directory. This second terminal will be for running the program while we edit it in the other.

:::co-tip
Now would be a great time to learn or practice your [TMUX](/tools/tmux/) skills instead of opening another terminal.
:::

How do we run it?

First try something. Type the same thing that you did into the file onto the command line in your new terminal and see what happens.

```sh
echo 'Hello world!'
```

```{.out}
Hello world!
```

Notice that it just runs.

*Now* let's run the commands in the file *as if they had been typed onto the command line*. For this we will use the `source` keyword.

```sh
source greet
```

:::co-btw
To save on typing you can use `.` instead of the word `source` but don't forget the space between the `.` and the file you are sourcing.
:::

```{.out}
Hello world!
```

Notice that it did exactly the same thing. The point is that the commands you type every day onto the command line *are* programming and just need to be placed into a file to become a program, more commonly known as a *script*.

There is another way to run the script. We can execute the `bash` interpreter command and give it the name of the file containing our `echo` command.

```sh
bash greet
```

```{.out}
Hello world!
```

Same thing! So what's the difference between using `source` and `bash`?

In order to understand the difference you first have to understand what is going on when you enter commands into the terminal.

When you open your terminal or login remotely and get on the command line you are actually running an interactive command interpreter, a program called a *shell* reads each command, evaluates and executes it, prints out any output and results, and then loops waiting for another command. This [human-computer interaction](/what/hci/) is sometimes called a [REPL](/what/repl/).

:::co-btw
This is where the name of popular sandbox coding web site <https://repl.it> comes from.
:::

The name of the shell interpreter program that is set by default for all interactive users (that's you) is `/bin/bash`. We'll go over this a lot but for now just know that program is *already* running and you are writing an ongoing script every time you enter a command.

When you use `source` (or `.`) you tell the *running* `/bin/bash` shell to essentially type each of the lines in the file as if you typed them yourself. They affect your immediate shell *environment*. Your environment is all the stuff related to you interactive running shell.

When you use `bash` you start a *new* `/bin/bash` shell program that is *not* interactive like the one you are interacting with. All programs are executed like this. The program is given its own environment and it *inherits* most things from your running environment. 

This separate execution of a program *of any kind* is called a *subprocess* sometimes also called a *subshell*. A *process* is just another name for any running program.

:::co-btw
You can play around with this difference by putting the word `exit` into your file and seeing what happens when you use `source greet` vs `bash greet`.
:::

*Need to complete these steps.*
