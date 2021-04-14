---
Title: Vim
Subtitle: Vi Improved
Query: true
---

*Vim* is best modern terminal editor due to its compatibility with [Vi](/tools/editors/vi/) and extensibility allowing it to rival [Emacs](/tools/editors/emacs/). Be careful to avoid becoming too dependent on [vimisms](vimisms/) which can destroy the [ubiquity](/what/ubiquitous/) of your Vi muscle memory, which is the single greatest advantage of becoming proficient with Vi in the first place.

:::co-faq

## What About NeoVim?

Vim is the standard editor included on *all* Linux systems by default. That should be reason enough. But there is actually no compelling reason to use NeoVim anyway. 

* Use [TMUX](/tools/tmux/) instead of panes. 
* No one cares about the codebase and developer team drama. 
* Vim is more modern than NeoVim these days.

Hell, Vim even has floating panels now that NeoVim does not if you are into that sort of thing. NeoVim simply fails to provide a significant alternative value proposition.

## What About Graphical Vim?

Sure it's nice, but the main reason to learn a terminal editor is to
take advantage of the [powerful](/tools/editors/vi/how/magic/)
possibilities of full shell integration, which is simply not possible
from the graphical version --- especially when running from Windows.

## What About Emacs?

Even if your workflow and projects would benefit from learning [Emacs](/tools/editors/emacs/) learn Vi anyway.

:::

## Installation

The `vi` command is already on your Unix/Linux machine. That's the beauty of it. However, you might want to get the full version of `vim` if this is your workstation or you want to use `vimtutor` for learning.

```sh
sudo apt update
sudo apt install vim
```

## Configuration {#config}

Configuring Vi/m is critically important to be effective. It is usually
good to start with [a configuration from someone you trust](./vimrc) and
alter that to your needs. Keep your `~/.vimrc` configuration
light-weight and portable and add it to your [dotfiles configuration
project repo](/tools/linux/tasks/dotfiles/vim). When done right you
won't even have to install any additional plugins because they'll be
installed automatically. The linked configuration does just that and
keeps it under 184 lines that can be easily and quickly be copied to any
system. 

## Plugins

Vim plugins are a great source of power when used judiciously. The leading plugin manager is currently [Plug](https://github.com/junegunn/vim-plug), which can be automatically downloaded and called on to install plugins when it is first launched in a way that does not interfere with other [dotfiles](https://duck.com/lite?kae=t&q=dotfiles) that you might have saved in a [Git repo](https://duck.com/lite?q=Git repo).

## Learning

The best way to learn Vim is to use Vim. (Now there's a surprise.) This
is why [`vimtutor`](https://duck.com/lite?kae=t&q=`vimtutor`) is
currently the best learning tool. It isn't complete and has flaws but is
better than anything else because it actually *is* Vim. After that maybe
try some of the others.

:::co-warning
Make absolutely sure that you properly learn to use [shell integration
with the magic wands](/tools/editors/vi/how/magic/) before you conclude
your beginner Vi training. You simply do *not* know Vi without them and
*every single tutorial on the planet leaves them out*. You seriously
*cannot* say you know Vi without understanding the single most powerful
feature of the world's most ubiquitous editor. There is little need for
additional Vim plugins when shell integration is properly understood.
:::

