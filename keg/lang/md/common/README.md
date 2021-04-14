---
Title: CommonMark
Subtitle: A Valiant Attempt to Standardize Markdown
Query: true
---

[*CommonMark*](https://commonmark.org) is the result of a collaboration lead by [JGM](https://johnmacfarlane.net) (the creator of [Pandoc](/tools/pandoc/)) to create some sort of agreement between all the different companies and organizations using [Markdown](/lang/md/) and splintering into dozens of different flavors all interpreting the non-standardized [original](https://daringfireball.net/projects/markdown/) in their own way. The effort was a success overall, but left out several design decisions including tables. That's right, [*CommonMark does not allow the use of any tables.*](https://talk.md/common.org/t/tables-in-pure-markdown/81) Tables were one of the main additions from [GitHub](/lang/md/github/) that people wanted (despite the pathetic design of GFM tables that often requires a tool of its own just to make them). Given how bad GFM tables are, it is no surprise that JGM politely left them out and quietly implemented a far superior --- and simpler --- design in [Pandoc Markdown](/lang/md/pandoc/), which is one reason why the resulting Pandoc Markdown is by far the simplest, cleanest, and most widely applicable Markdown version today. It has clearly demonstrated its design supremacy with its adoption into the R language and more (even if GitHub and the rest never catch up).

## See Also

* [CommonMark Specification](https://commonmark.org/)
* [CommonMark Tables Discussion](https://talk.md/common.org/t/tables-in-pure-markdown/81)
