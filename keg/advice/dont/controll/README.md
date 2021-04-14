---
Title: Don't Use `Control-l` to Clear the Screen
Subtitle: Use `set -o vi` and `alias c=clear` If You Must
Query: true
---

Nothing shows how little you know about the Unix terminal than to recommend using `Control-l` or `^L` to clear the screen without understanding that `set -o vi` exists and what it does.

`Control-l` only works in Emacs mode (`set -o emacs`), which is the default, but you should seriously consider changing to `set -o vi` in order to leverage all your other [Vi](/tools/editors/vi/) muscle-memory. The search capabilities with `Esc - /` alone are worth it. 

The only reason Emacs mode is the default is because the GNU project
created Bash and the GNU project invented and uses the horribly bloated
and isolated [Emacs editor](/advice/dont/emacs/).

If you really are that bothered by typing `clear` then add an `alias c=clear` to your Bash config. It's fewer keystrokes.
