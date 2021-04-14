---
Title: Can I Use HTML in Markdown?
Tags: faq
---

*You can but you really shouldn't.* There's never any need and doing so
creates unnecessary dependencies on a specific rendering format that
stop your content from being used in other ways. There is always a way
around it. Reach for [Pandoc](/lang/md/pandoc/) when you think you need
it. Consider using a template that looks for `tpl-style` in the YAML
data and if set adds a CSS link element so that `styles.css` can be kept
locally with your Markdown without polluting your Markdown instead. 
