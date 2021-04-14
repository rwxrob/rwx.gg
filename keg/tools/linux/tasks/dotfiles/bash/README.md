---
Title: Set Up Bash Dotfiles
Subtitle: '`.bashrc`, `.bash_profile`, `.bash_aliases`'
Query: true
---

Check if you have a `~/.bashrc` file.

```sh
ls ~/.bashrc
```

Now you need to decide to use `~/.bash_aliases` or `~/.bashrc`. The
`.bash_aliases` is usually plenty and doesn't interfere with what the
system *thinks* you should use for `.bashrc`. In fact, sometimes you
don't even have permission to change it (as is the case with PicoCTF).
However, sometimes changing the `.bashrc` gives you more control.
Consider doing that later. You can always use the same `dotfiles/bashrc`
file and put it into `~/.bash_aliases` if you are on a system that does
not allow you to change it (again, like PicoCTF).

Peek at `~/.bashrc` to See if `~/.bash_aliases` Sourced
Change into your `dotfiles` project repo directory
Make a `bash` directory
Make a new `bashrc` file (no initial period)
Create a symlink from ~/.bashrc to the bashrc

```sh
ln -fs "$PWD/bashrc" "$HOME/.bashrc"
```

Confirm the creation of the symlink.

```sh
file ~/.bashrc
```

## Resetting Your Shell

Before we do much more it is important to know how to get your Bash
configuration to reset. You could certainly close your terminal and
reopen it, or log back in, but the fastest way is simply to use the
`exec` command as follows:

```sh
exec bash
```

:::co-pwz
Some people will wrongly tell you to just to `bash` without the `exec
bash`. This creates subshells after subshells rather than replacing your
existing running shell process with a brand new one.
:::

This works even better if one of the first things in your `bashrc` file is a command to clear everything that was set before. Otherwise the new shell inherits all the environment from the previous shell and only changes those things that are new. This means that anything you have removed will still be there.

To clear *everything* from the previous settings consider unsetting all your aliases.

```sh
unalias -a
```

:::co-pwz
You cannot unset the previous environment variables and functions this way. Just remember you will need to `unset <thing>` when you make those changes otherwise they will be inherited by the `exec bash` shell.
:::

Now you can start building the rest of your custom `bashrc` file.

## Adding Aliases

First off let's add an alias.

How about making an alias for `clear`.

Open the `bashrc` file for editing.

```sh
alias c=clear
```

:::co-btw
No quotes are needed because there are no spaces.
:::

Now test it by resetting your shell.

```sh
exec bash
```

How about we add an alias with spaces requiring quotes. We'll use single quotes because they are easier to type.

```sh
alias ll='ls -la'
```

Now reset and test that one. 

Okay, looks like everything is working. 

You might want to take a moment and make a comment for the section about aliases so that you can find it fast later is it grows longer.

Now for some more important stuff.

```sh
set -o vi
```

Reset and test that your arrow keys still work. Vi mode means that you are effectively in `--INSERT--` mode when working with your history as if it were a big file, except you can only see one line. This is why your arrow keys work but not your `hjkl` navigation keys. 

You activate command mode by pressing `Esc` or the `Control-[` equivalent. 

Now can do anything to your history as if it were a large file you are working with in Vi command mode. This includes searching through your entire history with `/` and then finding the next occurrence with `n` or previous with `N`. Notice that `n` continuously wraps around with the search. This is *really* fast way to pull up complex lines from your history that you did weeks ago.

:::co-tip

Vi mode is one of the most powerful, yet unknown shell configurations of all. Not understanding its existence or power is why uninformed people suggest stupid things like [using `Control-L` to clear the screen](/advice/dont/controll/).

You don't need to learn all the shell `history` related shell commands and keys if you learn Vi mode.

:::

## Configure Your Shell Command History

Now is probably a good time to add more capacity to your history file, or remove it if you are on an ultra-secure system.

Type `history` to get a sample of your history. You will probably have a hundreds of lines.  Do it again and pipe the output to `less` if you want to see them all.

:::co-tip
Now might be a good time alias `more` to `less-R`.
:::

Your history is actually saved in a `~/.bash_history` file that you can reset at any time. Some people even preserve it in their `dotfiles` but that is probably not a good idea considering how much it changes.

Let's up the size of the history from the default of about 500.

Commonly the following environment variables are set related to history:

```bash
HISTCONTROL=ignoreboth  # don't save dups and starting with spaces
HISTFILESIZE=10000      # how many lines to save
```

You can learn more about them by searching for them in the [Bash manual page](https://duck.com/lite?kae=t&q=Bash manual page). There are many other settings related to history you can set.

