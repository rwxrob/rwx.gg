---
Title: Vimisms
Subtitle: Don't Ruin Your Vi Muscle Memory
Query: true
---

The whole point of using [vi](/tools/editors/vi/) is so that you can edit files anywhere on any computer and device. *Vimisms* defeat this entirely by training your muscle memory to be dependent on things that other systems won't have. (Might as well use Emacs if you are going to use them.) You won't have your precious `.vimrc` (which is another reason to keep your Vim configuration rather light and easy to [Secure Shell `scp`](https://duck.com/lite?kae=t&q=Secure Shell `scp`) over to any system when needed and allowed).

*Vimisms* are the things Vim would have you do --- or that people invent
--- that damage your ability to be productive on a system that has only
Vi (not Vim). The tragedy is that most of these are *not necessary and
not even efficient* but people just simply do not know the alternatives,
which very frequently are some variation of the Vi [magic
wands](/tools/editors/vi/how/magic/). Please consider avoiding all of
the following when using Vim.

------------------------------------- ------------------------------------- 
Vimism                                Vi Alternative
------------------------------------- ------------------------------------- 
Arrow keys                            `hjkl` as God intended 

Visual mode                           Ex commands with `:`

Remapping `ESC`                       Use built-in `Control` with `[`

Vimscript                             Bash shell with variations of `!`

Split panes                           Use TMUX or Screen instead

Number Increments (C-a/x)             Just change the number.
------------------------------------- ------------------------------------- 

Table Vimism Alternatives

:::co-mad
Few things anger experienced technologists who mentor others than seeing noobs corrupt other noobs muscle-memory with bad suggestions such as using these stupid customizations. *They are not worth the cost in multi-system proficiency.*
:::
