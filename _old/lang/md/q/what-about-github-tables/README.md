---
Title: What About GitHub Tables?
Tags: faq
---

Just don't use them. They are far too complicated and poorly designed. People have created tools just to be able to create GHF Markdown tables. They are a disaster but you might not realize how bad they are until you have experienced [Pandoc Markdown tables](https://pandoc.org/MANUAL.html#tables) instead. GitHub tables were *never* supported by Markdown and are still *not* supported by CommonMark for obvious reasons. So if you ever want to move to another Git hosting site (like GitLab, and you should) you will quickly discover your tables do not render. If and when you are ready to be okay with your tables not rendering on Git hosting services graduate up to [Pandoc Markdown](/lang/md/pandoc/) instead. It has much broader and deeper support than GFM and includes several table formats that are simpler *and* more powerful than GitHub's.
