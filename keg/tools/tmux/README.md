---
Title: TMUX
Subtitle: Modern Terminal Multiplexor
Query: true
---

*TMUX* is the world's best terminal multiplexor today. A descendent of `screen` TMUX allows a wide and powerful set of features beyond `screen` that include the following:

* Splitting a window into panes
* Scripted sessions and keystrokes
* Millions of colors
* Nested sessions
* Full customization

:::co-warn
Even though TMUX is a descendent of `screen` the creators decided not to use the same keys as `screen`, which is unfortunate because so many people have memorized them. You might seriously consider using a `screen`-centric [tmux.conf](./tmux.conf) file from the beginning and learn the defaults later. That way when you are on a system that does not have TMUX (nor allow it to be installed) you will be right at home with the `screen` that comes on such systems by default.
:::

:::co-tip
The fastest way to get up and running with TMUX is to copy someone else's [`tmux.conf`](./tmux.conf) file until you understand enough to configure your own. Just put it in your home directory as a `~/.tmux.conf`.
:::

## See Also

* <https://github.com/zolrath/wemux>
