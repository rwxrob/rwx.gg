---
Title: Linux Terminal Boost
Subtitle: Learn Linux Right
Query: true
---

Learning to use the Linux terminal is essential to your further growth and mastery later for a number of important tech [occupations](/jobs/). Unfortunately it is rather common for technologists to *never* learn the terminal, learn it *eventually* as a secondary priority, or to learn it *wrong* and create horrible habits that are difficult or impossible to break --- much like learning to hunt-n-peck to type rather than using home row.

:::co-tip
You *do* touch type don't you? If not, stop everything and learn type correctly *right now* before you do anything else! You will *never* master the terminal until you do. True terminal mastery fundamentally depends on home row typing. Might was wall just keep using your pretty graphic user interface and your mouse if you don't.
:::

Here's the thing. There are a *lot* of books, videos, apps, and sites that *claim* to help you learn the terminal but *most* of them are made by three types of people that you really want to avoid:

* Ancient pedantic Unix boomers
* Trend-chasing script kiddies and Mac cultists
* Linux noobs who are just so excited they have to share 
* Inexperienced click-baiting charlatans 

You can usually spot these types a mile away. But to help you, you might want to check the [list of stupid advice](/advice/dont/) they will say that will immediately let you know to quietly leave and never come back.

:::co-dum
The worst are usually the loudest and most flagrant who are obviously care more about their views and subscribes than you. They will usually have videos with gawd-awful thumbnails. Run away!! Hyperbole for hits is one thing. Hyperbole as a rhetorical device supported by objective facts is another.

And remember, you want to learn *Linux* essentials, *not* Mac essentials. Apple is so stupid they actually *removed* the `Escape` key from their keyboard. 
:::

Use the following table to review the different specific skills and commands you should eventually be ready and able to use and explain at any time. Note that this table contains *only* the modern essentials. 

:::co-btw
You will *never* be done learning Linux. That's what makes it so amazing. Accept it.
:::

  Category              Specifics
----------------------- ------------------------------------------------



 Manage Processes       `&` • `fg` • `pgrep` • `pkill` • `top` • `htop`
                        `pstree` • `ps` • `kill`

 Use Variables          `some=thing` • `other='some long'` •
                        `another="$other $some"` • `GLOBAL=thing` •
                        `$PATH` • `$CDPATH` • `$HOME` • `$PWD` •
                        `$RANDOM` • `${some}thing`


 Use Commands as        `$(date %s)`
 Variables

 Redirect and Pipe      `cmd > foo` • `cmd >> foo` • `cmd >| foo` • 
 Input and Output       `cmd < foo` • `cmd << EOM` • `cmd <<< str` •
                        `cmd | cmd` • `cmd 2>&1` • `cmd &> foo`


 Edit Text Files        `vi` • `vim` • `ed` • `pie`

 Use Aliases            `alias pie='perl -p -i -e '` • `unalias`

 Use Regular            `perl`
 Expressions

 Transform Text         `tr`

 Command History        `set -o vi`

 Screen Management      `screen` • `tmux`

 Web Commands           `lynx` • `curl`

 Find Files             `find` • `**`

 Analyze Files          `file` • `strings`

 Compress Files         ``

:::co-faq

## Why `perl` and no `sed` or `awk`?

Modern Bash 4+ and [Perl destroy the usefulness of `sed` and
`awk`](/advice/dont/perlhate/) to anyone who has been paying attention.
Many don't.

:::
