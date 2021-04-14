---
Title: What is a Static Site Generator?
Subtitle: A Tool for Building Static Web Sites
Query: true
---

A *static site generator* is a tool for creating a modern web site that is primarily static content that only changes when it is generated again. An SSG can be as simple as a [Bash shell script](/lang/bash/) combined with [Pandoc](/tools/pandoc/) or as fast and powerful as [Hugo](https://gohugo.io) which can render 1000s of pages in milliseconds on most hardware. This generation speed has shifted web application designs placing more emphasis on content rendered statically into HTML views and less dependent on slow, complex database queries that create bottlenecks and performance issues unnecessarily for content that is mostly read and not written.

:::co-warn
Be *extremely* cautious when using any of the JavaScript SSGs which are not in fact *actual* SSGs (including GatsbyJS, Nuxt.js, and VuePress). Some of these tools regularly and openly violate long-standing HTML web standards producing content that is only consumable by web browsers with JavaScript enabled. The "static" content is created by slowly crawling a front-end single page web application. Not only is this ridiculously slow compared to a tool like Hugo, but leaves the sites at risk of not being included in search engines and archiving tools because some part of it was not "crawlable". This is well documented on the JavaScript Sites in Search mailing list. 
:::

:::co-tip
Using an SSG allows your content to get closer to the person using it more reliably. Everything else can be put into a [function as a service](/what/cloud/faas/) provider or on one or two servers that use a JSON Web API for any dynamic content or content submissions.
:::
