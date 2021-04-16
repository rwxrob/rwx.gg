---
Title: Basic Markdown
Subtitle: Priority on Simplicity and Compatibility
---

Here is quick overview that you can complete in 20 minutes covering the
most basic, compatible [Markdown](/lang/md/) everyone should learn. The
priority is on simplicity and maximum compatibility. It allows your
content to be used *everywhere*. There is no more ubiquitous form of
knowledge storage.

:::co-pwz
If you really want to learn Markdown while reading about it then get
on your Linux command line right now and make a directory (maybe called
`lang/md/basic`) that has a `README.md` file in it. Then put your
practice markdown into it. `README.md` is something of a standard name
used by most people.
:::

## Paragraphs

A paragraph is either a long *single* line of text followed by a blank
line *or* multiple lines one after another. You pick but it's a good
idea to use one consistently throughout the same document or knowledge
base.

## Inline Formatting

One star for *italics*:

```markdown
one star for *italics* 
```

Two stars for **bold**:

```markdown
two stars for **bold**
```

Three stars for ***bold italics***:

```markdown
three stars for ***bold italics***
```

Backticks for `code` also known as `monospaced`:

```markdown
backticks for `monospaced`
```

[[What about \_ underscore?]](../q/what-about-underscore/)

## Escaping

Any time you don't want, say, a star (`*`) or a dollar sign (`$`) to be
recognized by the Markdown parser just put a backslash in front of it.
This goes for all the other special symbols used in Markdown. In fact,
you can put a backslash (` \ `) in front of any character whatsoever in
Pandoc Markdown to cause it to be used as is.

## Headings

Headings (often incorrectly called headers) begin with 1-6 hashtags
(`#`) followed by a space and then the title text followed by a single
blank line. 

```markdown
# Level One

Paragraphs and such here.

## Level Two

Paragraphs and such here.
```

Formatting is allowed in headings but can be problematic with some
renderers. Avoid if you can.

:::co-tip
Generally you should never have more than one first level heading (`#
Heading One`) because search engines prioritize it. When using Pandoc
you will not even need a level one heading because the `Title` is better
placed in the YAML meta-data property instead and rendered with the
Pandoc Template.
:::

## Links / Hyperlinks 

Hyperlinks (stuff you click on) come in three basic forms: 

1. Words
1. Raw URLs
1. Images

### Hyperlinked Words

The most common link in Markdown is just words you can click on that
take you to [local](#top) places or [external
sites](https://rwxrob.live). The web address must be either pointing to
a remote site or so something on the same site that document is on.

```markdown
Here is a [link to rwx.gg](https://rwx.gg).
```

### Autolinked Raw URLs

Sometimes you want to show the full web address.

```markdown
Here is URL to <https://rwx.gg/md> that will appear in full.
```

[[Can't I just use `http` for links?]](/lang/md/q/can-i-use-http-for-links/)

This also works with other link types besides `http`. (Yes, there are
several other URL schemas.)

```markdown
Mail me at <mailto:rwx@robs.io>.  
Phone me at <tel:555-555-5555>
```

## Images

Images are just links with an exclamation point in front. Make sure to
put a blank line before and after any image for maximum compatibility.
Inline images are not widely supported and mess up other formatting in
almost all cases.

```markdown
![gnome](/assets/img/mr-rob-gnome.jpg)
```

![gnome](https://rwx.gg/assets/img/mr-rob-gnome.jpg)

Images can also be used as links.

```markdown
[![RWXROB](/assets/img/mr-rob-gnome.jpg)](https://rwxrob.live)
```

[![gnome](/assets/img/mr-rob-gnome.jpg)](https://rwxrob.live){.nodecor}

:::co-tip
Since storing video files with your site is usually rather prohibitive
due to their size, consider taking a screenshot of the first frame of
the video hosted on a video hosting site and using that to link to the
external video on that site. That way the page will still load properly
even if you do not have Internet access which would otherwise block
embedding the video in the page instead. Besides, embedding videos is
generally a really bad idea because it adds HTML to your Markdown
unnecessarily causing it to be incompatible with other potential
rendering formats.
:::

## Lists

Simple lists are supported by pretty much everything. Put the list items
one to a line. Make sure to put a blank line after the list. 

Use stars followed by spaces (`* `) for bulleted (unordered) lists:

```markdown
* some
* thing
* here
```

* some
* thing
* here

Use the number one followed by a period and a space (`1. `) for numbered
(ordered) lists:

```markdown
1. HTML 
1. CSS
1. JavaScript
1. Go
1. Bash
```

1. HTML
1. CSS
1. JavaScript
1. Go
1. Bash

Always use `1.` so that if you change the order you do not have to
renumber the source itself. It will automatically change the number
order when rendered.

## Separators

Also called "horizontal rule." These just break up the page usually with
a horizontal line.

```markdown
----

```

Here comes an example of a separator.

----

Use four dashes for consistency even though there are dozens of ways to
indicate separation (some of which allow stars to be used as well). This
consistency allows you to easily find your separators and keeps them
from being confused with YAML markers (which use three dashes) and
inline formatting (which uses stars `*`).

## Hard Returns

Hard returns are a way of starting a new line within a given paragraph.
Type two spaces (`␣␣`) followed by the line return.

```markdown
Roses are red␣␣
Violets are blue
```

Rose are red  
Violets are blue

:::co-tip
The [Pandoc VIM plugin](/tools/editors/vim/#plugins) is particular useful
at showing these with [ligatures](/what/ligature/).
:::

```
Roses are red↵
Violets are blue
```

## Blocks

Blocks separate text or code from the document usually as a box. There
are two main block types to remember: *plain* (preformatted, as-is)
blocks and *code fences*. Both use three
[backticks](/key/) to "fence off" the text or code.

### Plain

When you just want the text to appear exactly as it is just use the
triple-backtick fence posts.

~~~markdown
```
    Roses are red
    Violets are violet
```
~~~
```
    Roses are red
    Violets are violet
```

### Code Fences

Code blocks are perhaps the single biggest reason to use markdown for
all your tech writing and note taking. If you want syntax highlighting
in your notes you get it for free. This provides very high-quality
publications very easily

:::co-pwz
Please be considerate of those who have trouble distinguishing between
colors when making a syntax highlighting style selection. Many cannot
see many of the colors making the text in those colors virtually
invisible.
:::

When you want to add syntax highlighting or otherwise indicate how the
text should be handle provide an information tag immediately following
the first triple-backtick fence, so for JavaScript it would be:

~~~markdown
```js
console.log('hello world')
```
~~~

Note that color syntax highlighting here has disabled because it can
distract from equal focus on everything being displayed and creates
unreadable text for those challenged with distinguishing different
colors --- especially in educational settings.

```js
console.log('hello world')
```

Although there are other ways to write blocks, using triple-backtick
fences is the most consistent way to do them all. This allows quickly
finding your blocks when editing as well as filtering them out easily
with scripting or simple parsing. It is also the most widely supported.
Discord, for example, only supports this format of code fence.

Here is a short list of supported language tags:

 Tag         Language
---------- ---------------------
 md          Markdown 
 json        JSON
 js          JavaScript 
 html        HTML 
 css         CSS 
 sh          Shell or Bash
 go          Go (golang)

#### Exception for Markdown

*Skip this if you are just beginning.*

In the single exceptional case where you need your block to contain
markdown code you should use three or four tildes (`~~~markdown` or
`~~~~markdown`). Again, this consistency allows you to filter out blocks
from simple scripts that examine each line which can be useful for
coding keyword searches and such.

~~~~

~~~markdown

Here is *some* markdown.

```js

console.log('hello example')

```

~~~

~~~~

~~~markdown

Here is *some* markdown.

```js

console.log('hello example')

```

~~~

There are literally an infinite number of possible ways to indicate a
block supported by the original and most derived Markdown parsers. Just
stick with these two options. Consistency is far more important than
artistic expression. Blocks are particularly important to keep
consistent because you will frequently want to simply strip them out for
keyword searches and such. Following these suggestions makes this
trivial even from simple shell scripts.

:::co-warn
Make sure there is no space after the backticks and before the block
identifier (`js` in the example).
:::

:::co-fyi
Technically paragraphs, lists, and even separators are also considered
blocks when parsed.
:::

## Blockquotes

Blockquotes are for quotations and *only* quotations. Avoid the
temptation to use them for anything else because if you do you can
semantically identify all the *actual* quotes in your content.

Begin each line of the block with a greater-than sign (right
[angle-bracket](/key/)).

Usually you will just have a single paragraph:

```markdown
> "One of the painful things about our time is that those who feel
> certainty are stupid, and those with any imagination and understanding
> are filled with doubt and indecision." (Bertrand Russell)
```

> "One of the painful things about our time is that those who feel
> certainty are stupid, and those with any imagination and understanding
> are filled with doubt and indecision." (Bertrand Russell)

:::co-tip
Use of quotation marks surrounding the text of the quote itself is
completely up to you but is recommended so that multiple quotes can be
combined next to one another without reader confusion.
:::

In the rare case that your quotation expands beyond a single line make
sure to join separate paragraphs with a blank line that is also
included:

```
> This is the first part of the quote.
>
> Here is the second part.
```

> This is the first part of the quote.
>
> Here is the second part.

## Comments

Comments are just simple HTML comments, which were inherited from SGML
dating back decades.

```markdown
<!-- This is a comment -->

This is a Markdown paragraph with another <!-- another --> comment in it.
```

<!-- This is a comment -->

This is a Markdown paragraph with another <!-- another --> comment in
it.

Even though you should seriously avoid making your Markdown HTML
dependent by including HTML as the original Markdown allowed, comments
are a fair exception to this because *no* Markdown has any other syntax
for indicating comments and *every* Markdown parser engines supports
them.

