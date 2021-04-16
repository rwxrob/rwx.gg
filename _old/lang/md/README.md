---
Title: Markdown
Subtitle: Invented by Writers, Not Scientists
Query: true
---

Markdown is a simple [markup](/what/markup/) language invented by writers who didn't want to write [HTML](/lang/html/). Markdown has become the standard for most technical writing and is supported on Reddit, StackExchange, Discord, GitLab, GitHub, and pretty much everywhere else.

--------------------------- ----------------------------------------------
 [Basic](/lang/md/basic/)        Emphasis on simplicity and compatibility
 [Common](/lang/md/common/)      General industry standard (but no tables)
 [Pandoc](/lang/md/pandoc/)      Maximum power and flexibility, no HTML dep
--------------------------- ----------------------------------------------

:::co-fyi
Modern companies such as GitLab depend *heavily* on Markdown for everything they do. Watch any of their release conference videos to see examples of all the documentation and presentations that make extensive use of Markdown (specifically [CommonMark](/lang/md/common/) since GitLab as an organization does not have a need for the Math and other advantages of Pandoc Markdown.)
:::

## Types of Markdown

Eventually you'll want to use [Pandoc Markdown](/lang/md/pandoc/) for
everything. (In fact, [this
document](https://gitlab.com/rwx.gg/README/-/tree/master/md) is written
in it.) Pandoc Markdown is far and away the most powerful, sustainable,
and supported format for capturing knowledge source, which is why the R
language project adopted it as their language documentation format.

However, you can learn [basic Pandoc Markdown](/lang/md/basic/) in less
than 30 minutes and it will work *everywhere*.

Why the differences?

The original Markdown was never standardized and has evolved into more than a dozen different flavors many of which are incompatible with one another. Thankfully there are only three versions that really matter:

## Standard Knowledge Source Language

Markdown has become the closest thing to a universal standard for writing all documentation today. In addition, authors write entire textbooks and novels in Markdown these days. When combined with a tool like [Pandoc](/tools/pandoc/) your Markdown can be converted to *every other format on the planet*. Markdown is the *only* markup writing language that can claim this. Therefore, *everyone* should learn it from those needing a solid, sustainable way to take notes, to those creating entire [knowledge bases](/what/knowledge/base/). Ironically Markdown is practically not taught at all today in any educational setting.

:::co-faq

## What about wikis?

[All wiki languages](https://duck.com/lite?kae=t&q=All wiki languages) are at least as complicated as HTML to learn. It's no surprise that they were invented by programmers and not writers, like Markdown was. [Wikis](/what/wikis/) obviously do have a place, however.

## What about LaTeX?

It is fully supported by [Pandoc Markdown](/lang/md/pandoc/).

## What about reStructuredText?

You didn't really ask that, did you? 

RST lost because it is so damn horrible and idiotically designed. Clearly the creators had nothing but documenting Python code in mind when they designed it.

:::

## See Also

* [Grav, Modern Open Source Flat-File CMS](https://getgrav.org/)
