---
Title: Set Up Lynx Text Web Browser
Subtitle: Browse and Search the Text Web
Query: true
---

Install Linux.

Fire up a terminal.

Install Lynx.

Configure Lynx.

Open or change into your `dotfiles` repo.

Create or copy a `lynx.cfg` file.

Create or copy a `lynx.lss` file.

Create  a `lynx` wrapper script that calls the *actual* Lynx pointing to the `lynx.cfg` and `lynx.lss` files.

:::co-pwz
It is easy to create a circular reference when creating the `lynx` wrapper. Just be careful.
:::

Tell your shell about the `lynx` wrapper and where to find it.

The easiest way might be to just add an extra entry to your path.

