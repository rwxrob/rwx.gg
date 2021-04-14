---
Title: Survive Vi
Subtitle: The Absolute Minimum You *Must* Learn
Query: true
---

People like to joke about how hard it is to exit Vi. It's true that for years people would be suddenly forced to use Vi without expecting it, for example, when editing a [`crontab`](https://duck.com/lite?kae=t&q=`crontab`) was a common one. Here are the absolute basics to help you survive your first encounter with Vi with a smile.

:::co-fun

Now that [Nano terminal editor](https://duck.com/lite?kae=t&q=Nano terminal editor) is the default on all major Linux [distros](/tools/linux/distros/) Vi users are the ones stuck not being able to exit.

:::

Keep in mind that the arrow keys are now fully supported by Vi itself (not just Vim) even though you should learn to use `hjkl` navigation as soon as possible.

------------------ --------- ------------------------------------------
 Action              Mode    Description
------------------ --------- ------------------------------------------
 `i`                Command  Change into `-- INSERT --` mode.

 `ESC`             `INSERT`  Escape `-- INSERT --` to Command.

 `←`                 Both    Move left.

 `↓`                 Both    Move down.

 `↑`                 Both    Move up.

 `→`                 Both    Move right.

 `:wq`             Command   Save and quit.

 `:q!`             Command   Quit without saving.

 `:w`              Command   Save without exiting.
------------------ --------- ------------------------------------------

See that's not so bad! Here's a [PNG](./visurvive.png) if you want a copy.

Seriously, that's all you need to get started with Vi/m and never bother with Nano at all. In fact, that's *easier* than learning Nano. Just say no to Nano!

:::co-btw
The reason `ZZ` and `ZQ` and `Control-[` were removed from the survival guide is that too many beginners would experiment with different combinations having forgotten how to exit and end up accidentally doing `Control-Z` which backgrounds the process. This is yet another reason `nano` is so bad for beginners.
:::
