---
Title: Integrate Shell Commands Into Vi Workflow
Subtitle: Exclamation Points are the Magic Wands of Vi, Learn Them
Query: true
---

:::co-pwz
You seriously don't know Vi until you know how to *fully* integrate the shell and transform lines by passing them to shell commands that then *replace* those lines with their output.

Shell integration is the single most powerful aspect of Vi and yet it is hardly known, documented, or even used by veteran users. 

Unfortunately, those who *don't* understand the power of full shell integration build stupid things like [NeoVim](/tools/editors/neovim/).
:::

*Magic wands* are a mnemonic device (originally used by Rob Muhlestein) for using the [exclamation point](/key/) to send the current line, section, or page to any command that can be run from the shell and replace those lines with the output of that command. 

In fact, once you master the keystrokes you can use the same to accomplish other tasks that use the `ex` command line --- like finding and replacing --- by just backspacing out the exclamation point. The automatic `ex` line range inference is a really powerful shortcut that is built into *all* Vi versions.

:::co-comment
Yeah, you don't need all those Vimscript macros. No really, you don't. Just use the shell instead.
:::

### Line Wand

The line wand is by far the most frequent wand you will use. It is the fastest and most common way to extend `vi`. Just spam bang twice to send the current line to the shell command or just replace the current line with the output of a utility command like `date`, `cal` or *any* text that comes out of *any* command.


#### Send Line to Bash: `!!bash`

Before:

```bash
for i in {1..20}; do echo Item $i.; done
```

After:

```{.out}
Item 1.
Item 2.
Item 3.
Item 4.
Item 5.
Item 6.
Item 7.
Item 8.
Item 9.
Item 10.
Item 11.
Item 12.
Item 13.
Item 14.
Item 15.
Item 16.
Item 17.
Item 18.
Item 19.
Item 20.
```

#### Send Line to Calculator: `!!bc`

:::co-warning
This requires the `bc` program to be installed, which is by default on most Linux systems. 
:::

Before:

```bc
232432 * 2342342 / 234
```

After:

```{.out}
2326646306
```

#### Send Line to Python3: `!!python3`

Before:

```python
[print(f"Item {x}.") for x in range (1,21)]
```

After:

```{.out}
Item 1.
Item 2.
Item 3.
Item 4.
Item 5.
Item 6.
Item 7.
Item 8.
Item 9.
Item 10.
Item 11.
Item 12.
Item 13.
Item 14.
Item 15.
Item 16.
Item 17.
Item 18.
Item 19.
Item 20.
```

### Section Wand

Sending a section is a great way to run part of a larger script (say for configuration) without having to cut it out and put it into another file. If you use this easy way, don't forget to undo it later to replace it with the original code (instead of the output of the lines of code).

#### Send Section to Bash: `!}bash`

Before:

```bash
for i in {1..20}; do
    echo Item $i.
done
```

After:

```{.out}
Item 1.
Item 2.
Item 3.
Item 4.
Item 5.
Item 6.
Item 7.
Item 8.
Item 9.
Item 10.
Item 11.
Item 12.
Item 13.
Item 14.
Item 15.
Item 16.
Item 17.
Item 18.
Item 19.
Item 20.
```

#### Send Section to Sort: `!}sort`

Before:

```markdown
* [\@zerostheory](https://twitch.tv/zerostheory)
* [\@OGLinuk](https://twitch.tv/OGLinuk)
* [\@elementhttp](https://twitch.tv/elementhttp)
* [\@elremingu](https://twitch.tv/elremingu)
* Jovan Vladislav
* [\@mtheory11dim](https://twitch.tv/mtheory11dim)
* [\@smartrefigerator](https://twitch.tv/smartrefigerator)
* [\@returntodust](https://twitch.tv/returntodust)
* [\@MousePotato](https://twitch.tv/MousePotato)
* [\@unres1gned](https://twitch.tv/unres1gned)
```

After:

```markdown
* [\@elementhttp](https://twitch.tv/elementhttp)
* [\@elremingu](https://twitch.tv/elremingu)
* Jovan Vladislav
* [\@MousePotato](https://twitch.tv/MousePotato)
* [\@mtheory11dim](https://twitch.tv/mtheory11dim)
* [\@OGLinuk](https://twitch.tv/OGLinuk)
* [\@returntodust](https://twitch.tv/returntodust)
* [\@smartrefigerator](https://twitch.tv/smartrefigerator)
* [\@unres1gned](https://twitch.tv/unres1gned)
* [\@zerostheory](https://twitch.tv/zerostheory)
```

### Line Number Wand: `!:<lineafter>`

Line number wand is just a larger version of the [section wand] but allows blank lines to be included. It takes more keystrokes, however. Start with a bang and follow it up with a colon and the line *after* the last line that you want to include.

For the following example imagine the first line is line 10 of the file and you want to send it to `pandoc` for rending as HTML. You would position your cursor on the first `#` (on line 10) and type `!:22<enter>` (because there are 11 lines to send). Then type `pandoc<enter>` after the `!` as the command to receive the lines and replace them.

```markdown
# This is a title with *emphasis*

* one
* two
* three

How about a paragraph

## Thursday, March 12, 2020, 7:46:39PM

something
```

```{.out}
<h1 id="this-is-a-title-with-emphasis">This is a title with <em>emphasis</em></h1>
<ul>
<li>one</li>
<li>two</li>
<li>three</li>
</ul>
<p>How about a paragraph</p>
<h2 id="thursday-march-12-2020-74639pm">Thursday, March 12, 2020, 7:46:39PM</h2>
something
```

## Creating Transformational Filter Commands

Rather than using a Vim-only macro or even tying yourself to Vim at all you can write simple shell scripts and use them much faster than any hotkey combination you could set and with greater flexibility. *Anything* you can code into an executable command that can run from the shell can be inserted using this method.

### Commenting Something Out 

Sure you could to `:.,+10s/^/# ` and make a mapping for that, or use visual mode and spam your way through the same sort of thing--- or you could just to `!}cmt` and be done with it. Your fingers will thank you later. Plus you already are used to using the wands for everything else.

Before: 

```sh
for i in {1..20}; do
    echo Item $i
done
```

Here's the command to add to your PATH:

```bash
#!/bin/bash
while IFS= read -r line; do
  echo "${1:-#} $line"
done <&1
```

And then `!}cmt`

```sh
# for i in {1..20}; do
#  echo Item $i
# done
```

No mess. No fuss. No Vimrc bloat. Just another command that follows the Unix philosophy of doing one thing really well that integrates easily. You can even integrate your `cmt` with other scripting that has nothing to do with your Vi editing session, which you could never do if you can built it into your bloated `vimrc` file.

You can change the comment character just by calling `cmt // ` instead.

:::co-tip
By the way, this makes the power of Perl and Bash regular expressions really obvious.
:::
