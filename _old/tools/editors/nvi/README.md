---
Title: 'NewVi `nex`/`nvi` Text Editor'
Subtitle: '*Exactly* Compatible with Original Vi'
Query: true
---

The `nvi` editor is the current default editor on OpenBSD and available on several other Unix and Linux distributions. The fact that it is "bug-for-bug compatible" with the original `vi` makes is a *very* useful editor to install and learn since it makes it very clear when you are using a [Vimism](/tools/editors/vim/vimisms/) such as `Control-V` or even forgetting to activate your `set ruler` or `set showmode` (which are not on by default in `vi`.

:::co-tip
It is interesting to note that arrow keys *are* a part of the original `vi` --- even in insert mode --- even though learning to use `hjkl` is obviously preferable. Testing with `nvi` was the way this was discovered.

Also, you might want to create a `~/.exrc` section in your `~/.vimrc` so that you have the lines available to pull over if you are ever on a system that does not have `vim` installed. Many BSD systems don't by default.
:::
