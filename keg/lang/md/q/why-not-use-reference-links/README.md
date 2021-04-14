---
Title: Why Not Use Reference Links in Markdown?
Tags: faq
---

Never use "reference links" that allow you to put the actual
link elsewhere in the document usually at the bottom. While this made
sense when Markdown was invented and the emphasis was on how pretty and
readable the Markdown source text itself was this emphasis has changed
today. Separating your link several hundred lines away from the text
that is linked is *really* bad practice because it decouples the link
making it much harder to copy and paste to another document --- or worse
--- you will edit your document and completely forget to remove them
leaving them orphaned.

