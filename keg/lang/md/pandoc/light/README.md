---
Title: Pandoc *Light* Markdown
Subtitle: Great Taste, Less Filling 
Query: true
---

* [Implicit Header Refs]
* [Tables]
* [Attributes]
* [Line Blocks]
* [Nested Lists]
* [Loose Lists]
* [Divs]
* [Spans]
* [Strikeout]
* [Subscript]
* [Superscript]
* [YAML]


| something
| thing here that
| is on another line


[Pandoc Markdown](/lang/md/pandoc/) is the supreme king of Markdown flavors providing the most *power*, *ubiquity*, and *simplicity*. But there is a *lot* in full Pandoc Markdown that most people simple do not need --- especially those maintaining simple documentation, blogs, and [knowledge apps](/what/knowledge/apps/). Here are some suggestions to help you get all the great taste, without the bloat.

:::co-tip
Even though we are keeping things light and simple we can sleep well at night knowing the ultimate power of Pandoc is available if needed later --- all the way up to coding our own Lua extensions. The overwhelming reason to pick Pandoc over all other engines is this ability to scale as the complexity of your knowledge source and applications scale all without blowing yourself up with standards-violating disasters like [GatsbyJS](https://duck.com/lite?kae=t&q=GatsbyJS), [VuePress](https://duck.com/lite?kae=t&q=VuePress), [NuxtJS](https://duck.com/lite?kae=t&q=NuxtJS), [ReactJS](https://duck.com/lite?kae=t&q=ReactJS) and other inferior alternatives.
:::

## Limit Yourself

Most people will only really benefit from the following Pandoc-only indulgences, which are fine because they still display well on GitLab and GitHub and are the best candidates for inclusion in the [CommonMark standard](/lang/md/common/) eventually.

:::co-tip
Remember, you should start with [Basic Markdown](/lang/md/basic/) first to maximize consistency and compatibility. Minimizing the number and type of Pandoc extensions is always a good idea for much the same reason that you might choose to write a shell script using [POSIX Dash shell](/lang/dash/) instead of Bash unless you need things that Bash reasonably provides better than overly complex subprocess alternatives. 
:::

*Need to complete.*

:::co-tip

## Don't Use GFM Todo Lists

Just don't use GFM list items. They are much harder to maintain than simply having a list of completed items and another of unfinished items and moving the one line to the other list, which is also much harder to misunderstand and clearer to read for everyone.

## Don't Use PHP Definitions

It is always better to hyperlink to a definition on its own [knowledge node](/what/knowledge/node/).

## Don't Use Numbered Examples

Their utility can be completely covered by an `id` inline attributes and simply using local links instead.

:::

## See Also

* [Pandoc Markdown](https://pandoc.org/MANUAL.html#pandocs-markdown)
