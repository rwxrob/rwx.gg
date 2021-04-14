---
Title: Basic Markdown
Subtitle: Simplicity and Compatibility
---

Here is quick overview that you can complete in 20 minutes covering the most basic, compatible [Markdown](/lang/md/) everyone should learn. The priority is on simplicity and maximum compatibility. It allows your content to be used *everywhere*. There is no more ubiquitous form of knowledge storage.

1. [Paragraphs]
1. [Formatting]
1. [Escaping]
1. [Headings]
1. [Links / Hyperlinks]
    1. [Hyperlinked Words]
    1. [Autolinked URLs]
1. [Images]
1. [Lists]
1. [Separators]
1. [Hard Returns]
1. [Blocks]
    1. [Plain]
    1. [Code Fences]
       1. [Don't Forget the Blank Lines]
       1. [Exception for Markdown]
1. [Blockquotes]
1. [Comments]
1. [FAQ](#faq)
    1. [Why Not Use Tables?]
    1. [What About GitHub Tables?]
    1. [Can't I Use Reference Links at the Bottom?]
    1. [What About \_ Underscore?]
    1. [Can't I Just Use `http` for Links?]

## Paragraphs

A paragraph is either a long *single* line of text followed by a blank
line *or* multiple lines one after another. Either is fine and there are
solid arguments to use one or the other. You pick but it's a good idea
to use one consistently throughout the same knowledge base.


To move paragraphs around all you have to do is copy that single
line (`yy` or `dd` in [Vi](/tools/editors/vi/)). Single-line paragraphs
always properly wrap no matter what editor you are using or its width.
Don't forget to put one blank line before and after each paragraph to
separate it from the rest. Paragraphs frequently contain
[formatting](#formatting).


Also note that this does *not* apply to the YAML front-matter at the top of a Markdown document. This is explained more in [Pandoc Light Markdown](/lang/md/pandoc/light/) but just know that limiting text blocks to 72 --- small enough for allow for line numbering and a gutter on the left-side required by Vim and other editors --- columns is wise for such things since YAML depends so heavily on clean left-side indentation. 

:::

```

## Example of Good Wrapping

This is a rather long paragraph that extends well beyond the 80 characters of most terminals and in the old days would be chopped into smaller individual lines rather than this big long huge line.

```

Although every Markdown tool allows multiple lines to be grouped into a single paragraph so long as a they are followed by a blank line this practice is rather bad when editing those same files with editor or terminal or [tmux](/tools/tmux/) pane widths that are less than the line length. They are completely unreadable. Stay safe, just use single long lines. All modern (and even most ancient editors) support auto-wrapping the line in the terminal rather than effectively hard wrapping it in the display.

```chopme

## Example of Bad Wrapping

This is a rather long paragraph that has been chopped forcefully into
individual likes of less than 80 characters each. While this might look
reasoably well on a standard 80 character terminal it looks horrible when in
modern TMUX panes that might contain far less than that and wastes available
characters when more are available.

```

## Formatting

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

[What about \_ underscore?]

## Escaping

Any time you don't want, say, a star (`*`) to be recognized by the Markdown parser just put a backslash in front of it. This goes for all the other special symbols used in Markdown. In fact, you can put a backslash (` \ `) in front of any character whatsoever in Pandoc Markdown to cause it to be used as is.

:::co-pwz

One particularly common character to need escaping is the dollar sign (`$`). This is because Pandoc supports the inclusion of LaTeX for math notation and uses the dollar sign to mark it. Jut remember to always put a backslash in front of it and you are fine.

:::

## Headings

Headings (often incorrectly called headers) begin with 1-6 hashtags (`#`) followed by a space and then the title text followed by a single blank line. 

```markdown

# Level One

Paragraphs and such here.

## Level Two

Paragraphs and such here.

```

Formatting is allowed in headings but can be problematic with some renderers. Avoid if you can.

:::co-tip

Generally you should never have more than one first level heading (`# Heading One`) because search engines prioritize it. When using Pandoc you will not even need a level one heading because the `Title` is better placed in the meta-data property instead and rendered with the Pandoc Template.

:::

## Links / Hyperlinks 

Hyperlinks (stuff you click on) come in three basic forms: 

1. Words
1. URLs
1. Images

### Hyperlinked Words

The most common link in Markdown is just words you can click on that take you to [local](#top) places or [external sites](https://rwxrob.live). The web address must be either pointing to a remote site or so something on the same site that document is on.

```markdown

Here is a [link to rwx.gg](https://rwx.gg).

```

### Autolinked URLs

Sometimes you want to show the full web address.

```markdown

Here is URL to <https://rwx.gg/md> that will appear in full.

```

[Can't I just use `http` for links?]

This also works with other link types besides `http`. (Yes, there are several other URL schemas.)

```markdown

Mail me at <mailto:rwx@robs.io>.  
Phone me at <tel:555-555-5555>

```

## Images

Images are just links with an exclamation point in front. Make sure to put a blank line before and after any image for maximum compatibility. Inline images are not widely supported and mess up other formatting in almost all cases.

```markdown

![gnome](/assets/img/mr-rob-gnome.jpg)

```

![](https://rwx.gg/assets/img/mr-rob-gnome.jpg)

Images can also be used as links.

```markdown

[![RWXROB](/assets/img/mr-rob-gnome.jpg)](https://rwxrob.live)


```

[![gnome](/assets/img/mr-rob-gnome.jpg)](https://rwxrob.live){.nodecor}

:::co-tip

Since storing video files with your site is usually rather prohibitive due to their size, consider taking a screenshot of the first frame of the video hosted on a video hosting site and using that to link to the external video on that site. That way the page will still load properly even if you do not have Internet access which would otherwise block embedding the video in the page instead. Besides, embedding videos is generally a really bad idea because it adds HTML to your Markdown unnecessarily causing it to be incompatible with other potential rendering formats.

:::

## Lists

Simple lists are supported by pretty much everything. Put the list items one to a line. Make sure to put a blank line after the list. 

Use stars followed by spaces (`* `) for bulleted (unordered) lists:

```markdown

* some
* thing
* here

```

* some
* thing
* here

Use the number one followed by a period and a space (`1. `) for numbered (ordered) lists:

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

Always use `1.` so that if you change the order you do not have to renumber the source itself. It will automatically change the number order when rendered.

## Separators

Also called "horizontal rule." These just break up the page usually with a horizontal line.

```markdown

----

```

Here comes an example of a separator.

----

Use four dashes for consistency even though there are dozens of ways to indicate separation (some of which allow stars to be used as well). This consistency allows you to easily find your separators and keeps them from being confused with YAML markers (which use three dashes) and inline formatting (which uses stars `*`).

## Hard Returns

Hard returns are a way of starting a new line within a given paragraph. Type two spaces (`␣␣`) followed by the line return.

```markdown

Roses are red␣␣
Violets are blue

```

Rose are red  
Violets are blue

:::co-tips

The [Pandoc VIM plugin](/tools/editors/vim/#plugins) is particular useful at showing these with [ligatures](/what/ligature/).

:::

```

Roses are red↵
Violets are blue

```

## Blocks

Blocks separate text or code from the document usually as a box. There are two main block types to remember: *plain* (preformatted, as-is) blocks and *code fences*. Both use three [backticks](/key/) to "fence off" the text or code.

### Plain

When you just want the text to appear exactly as it is just use the triple-backtick fence posts.

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

Code blocks are perhaps the single biggest reason to use markdown for all your tech writing and note taking. If you want syntax highlighting in your notes you get it for free. This provides very high-quality publications very easily

:::co-pwz

Please be considerate of those who have trouble distinguishing between colors when making a syntax highlighting style selection. Many cannot see many of the colors making the text in those colors virtually invisible.

:::

When you want to add syntax highlighting or otherwise indicate how the text should be handle provide an information tag immediately following the first triple-backtick fence, so for JavaScript it would be:

~~~markdown

```js

console.log('hello world')

```

~~~

Note that color syntax highlighting here has disabled because it can distract from equal focus on everything being displayed and creates unreadable text for those challenged with distinguishing different colors --- especially in educational settings.

```js

console.log('hello world')

```

Although there are other ways to write blocks, using triple-backtick fences is the most consistent way to do them all. This allows quickly finding your blocks when editing as well as filtering them out easily with scripting or simple parsing. It is also the most widely supported. Discord, for example, only supports this format of code fence.

Here is a short list of supported language tags:

 Tag         Language
---------- ---------------------
 md          Markdown 
 json        JSON
 js          JavaScript 
 html        HTML 
 css         CSS 
 sh          Shell or Bash

#### Don't Forget the Blank Lines

Even though all forms of Markdown support code fences that immediately contain content without a blank line after the opening fence post or before the final fence post you really should always add one and adjust your rendering styles appropriately. This is because *without* that initial blank browsers like FireFox that correctly respect whitespace as a DOM node will show your code incorrectly in a pre-formatted style. Plus the separation makes your knowledge source *much* easier to read and search.

~~~markdown

Don't do this:

```js
console.log('Hello')
```

Do this instead:

```js

console.log('Hello')

```

~~~

#### Exception for Markdown

In the single exceptional case where you need your block to contain markdown code you should use three or four tildes (`~~~markdown` or `~~~~markdown`). Again, this consistency allows you to filter out blocks from simple scripts that examine each line which can be useful for coding keyword searches and such.

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

There are literally an infinite number of possible ways to indicate a block supported by the original and most derived Markdown parsers. Just stick with these two options. Consistency is far more important than artistic expression. Blocks are particularly important to keep consistent because you will frequently want to simply strip them out for keyword searches and such. Following these suggestions makes this trivial even from simple shell scripts.

:::co-warn

Make sure there is no space after the backticks and before the block identifier (`js` in the example).

:::

:::co-fyi

Technically paragraphs, lists, and even separators are also considered blocks when parsed.

:::

## Blockquotes

Blockquotes are for quotations and *only* quotations. Avoid the temptation to use them for anything else because if you do you can semantically identify all the *actual* quotes in your content.

Begin each line of the block with a greater-than sign (right [angle-bracket](/key/)).

Usually you will just have a single paragraph:

```markdown

> "One of the painful things about our time is that those who feel certainty are stupid, and those with any imagination and understanding are filled with doubt and indecision." (Bertrand Russell)

```

> "One of the painful things about our time is that those who feel certainty are stupid, and those with any imagination and understanding are filled with doubt and indecision." (Bertrand Russell)

:::co-tip

Use of quotation marks surrounding the text of the quote itself is completely up to you but is recommended so that multiple quotes can be combined next to one another without reader confusion.

:::

In the rare case that your quotation expands beyond a single line make sure to join separate paragraphs with a blank line that is also included:

```

> This is the first part of the quote.
>
> Here is the second part.

```

> This is the first part of the quote.
>
> Here is the second part.

## Comments

Comments are just simple HTML comments, which were inherited from SGML dating back decades.

```markdown

<!-- This is a comment -->

This is a Markdown paragraph with another <!-- another --> comment in it.

```

<!-- This is a comment -->

This is a Markdown paragraph with another <!-- another --> comment in it.

Even though you should seriously avoid making your Markdown HTML dependent by including HTML as the original Markdown allowed, comments are a fair exception to this because *no* Markdown has any other syntax for indicating comments and *every* Markdown parser engines supports them.

:::{#faq .co-faq}

## Why Not Use tables?

Tables are not officially supported by the original Markdown nor are they supported in the CommonMark standard. Use an image of a table if you really need *basic* compatibility.

When the time comes to use tables you will want to use [simple Pandoc Markdown tables](https://pandoc.org/MANUAL.html#tables) instead, which work even on systems that do not support them because of their use of dashes which get interpreted as lines in original Markdown parsers.

## What About GitHub Tables?

Just don't use them. They are far too complicated and poorly designed. People have created tools just to be able to create GHF Markdown tables. They are a disaster but you might not realize how bad they are until you have experienced [Pandoc Markdown tables](https://pandoc.org/MANUAL.html#tables) instead. GitHub tables were *never* supported by Markdown and are still *not* supported by CommonMark for obvious reasons. So if you ever want to move to another Git hosting site (like GitLab, and you should) you will quickly discover your tables do not render. If and when you are ready to be okay with your tables not rendering on Git hosting services graduate up to [Pandoc Markdown](/lang/md/pandoc/) instead. It has much broader and deeper support than GFM and includes several table formats that are simpler *and* more powerful than GitHub's.

## Can I Use HTML in My Markdown?

*You can but you never should.* There's never any need and doing so creates unnecessary dependencies on a specific rendering format that stop your content from being used in other ways. There is always a way around it. Reach for [Pandoc](/lang/md/pandoc/) when you think you need it. Consider using a template that looks for `tpl-style` in the YAML data and if set adds a CSS link element so that `styles.css` can be kept locally with your Markdown without polluting your Markdown instead. 

## Can't I Use Reference Links at the Bottom?

*Hell no!* Never use "reference links" that allow you to put the actual
link elsewhere in the document usually at the bottom. While this made
sense when Markdown was invented and the emphasis was on how pretty and
readable the Markdown source text itself was this emphasis has changed
today. Separating your link several hundred lines away from the text
that is linked is *really* bad practice because it decouples the link
making it much harder to copy and paste to another document --- or worse
--- you will edit your document and completely forget to remove them
leaving them orphaned.

## What About \_ Underscore?

*Do not use underscore ever.* Yes the original Markdown supported it,
but it was always a bad idea. Eventually it will bite you. It is not
widely supported and causes problems when used for actual names. 

Also, avoid combining or overlapping formatting as much as possible and
avoid formatting a portion of a single word, which is not handled
consistently --- especially by syntax-highlighting tools and editors.

## Can I Just Use `http` for Links?

*No.* Never trust that your external links will be detected and
autolinked automatically just because they start with `http`. There are
still Markdown engines out there that do not autodetect them since
autodetection was not a part of the original and has never been a part
of any standard Markdown version even though this is one of the dumbest
things GitHub added to their flavor of Markdown.

## Should I Use Multiline Paragraphs?

There are strong arguments for either one.

### Multiline

* Historically multiline has been much more popular and is the standard
  for many formal text document formats such as the IETF RFC
  specifications.

* Multiline wrapping can be synchronized with rendered text width --- 72
  `textwidth` to 700px for example. While no rendering style should ever
  depend on such a thing, it is helpful to avoid odd outliers in a
  bulleted list, table, or paragraph orphans that, while strictly
  speaking are okay, would still be nice to avoid.

* While it may seem counter-intuitive to hard wrap wasting the maximum
  width available on any terminal or TMUX pane, long lines are much more
  difficult reading while editing documents --- particularly if the one
  writing them is working with them for long spans of time.

* Seeing ugly hard wrapping when a terminal or TMUX pane is not wide
  enough provides an immediate visual queue that it needs to be widened
  in order to display the document correctly. This is *particularly*
  important if your Markdown has a lot of text in the YAML header.

* The `vim-pandoc` plugin makes for drop-dead simple formatting. Just
  make sure to `let g:pandoc#formatting#mode = 'hA'`{.vim} and `let
  g:pandoc#formatting#textwidth = 72`{.vim} so that formatting is not
  something you even have to think about.

### Single Line

* Paragraphs become one single long line allowing them to be easily
  copied and pasted with a single `dd` or `yy` in Vi/m.

* None of the terminal or pane width is wasted.

* Smaller panes --- such as with TMUX --- make hard wrapped paragraphs
  look *very* ugly. But this is also an indicator that your pane should
  be widened.

## What Should My Text Width Be?

*72 has been the standard.* This provides eight characters for line
numbering and other gutter symbols on an 80 character screen. Even
though many --- including Linus Torvaldz --- now widely agree limiting
code to 80 characters is no longer practical given the new standard
longer terminal widths.

:::
