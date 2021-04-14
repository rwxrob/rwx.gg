---
Title: Don't Pick Emacs Just Cuz
Subtitle: Ed and Vi are *the* POSIX Editors and *far* More Ubiquitous
Query: false
---

Emacs is used by people who rarely use more than one computer terminal. Vi is used by people who want to be functional on *any* computer with a Unix/Linux terminal. Emacs users are helpless without their bloated, customized, personalized editor on what is likely their bloated, customized, heavy workstation. Expecting to be able to install Emacs on any computer you are required to work on is ridiculously naive. Most big companies where you are likely to work on a lot of Linux and Unix systems actually still practice some form of change management. Getting your Emacs installed ain't gonna get approved.

## Possibly Misguided Influence from Educators

Unfortunately many have reported that a large number of professors and academics recommend Emacs to their students without even pausing to consider how those students will go on to use the editor. Fostering a dependency on a tool that fundamentally handicaps *most* technologists in *most* professions involving a Unix/Linux terminal is downright irresponsible.

## Vi/m Isn't Better, Except When It Is

It's not that Vi/m is *better* than Emacs --- except that it actually
is. It's definitely a different tool for a different type of user. But
the problem is *most* people have no idea how to access the incredible
power plain old `vi` provides by building on the Unix philosophy and
*[integrating with the shell](/tools/editors/vi/how/magic/).*

For example, *very* few people realize the immense value of Vi's
[wands](/tools/editors/vi/how/magic/). When combined with filtering
shell scripts and TMUX Vi/m easily surpasses Emacs in raw power being
able to fully integrate the *entire Linux system* rather than just
loading yet another Emacs plugin. Vi/m has plugins, but the power is
*not* in the plugins. It is in the raw integration with *anything* that
can be run from the command line.

:::co-tip
Next time you are fighting with an Emacs user ask them if they ever
actually learned Vi/m before they decided to attack it. Most have never
mastered it and 0.0001% even know how to use the
[wands](/tools/editors/vi/how/magic/). Then again, only 0.1% of Vi/m
users even know about them. :::
