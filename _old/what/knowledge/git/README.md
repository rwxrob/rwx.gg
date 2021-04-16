---
Title: Git for Knowledge Management
Subtitle: Commit Often, Don't Branch, Log Changes
Query: true
---

[Git](https://duck.com/lite?kae=t&q=Git) is the world's best tool for managing source code and is therefore the best for managing knowledge captured as [Pandoc Markdown](/lang/md/pandoc) source. However, much of Git's functionality does not apply to the much more dynamic work of maintaining and editing knowledge source. Knowledge source is also much less able to benefit from the same level of fully automated [CI/CD](/what/knowledge/cicd/) because of this and the essential need for *human* editors to do what might otherwise be automated unit testing. Here are a few guidelines to get started:

* Commit often
* Create a [`save`](/tools/git/save/) command
* Don't bother branching
* Work entirely on `master` branch
* Log changes to a [changes](/changes/) [knode](/what/knowledge/node/)
* Consider a different group just for knowledge source to prevent flooding those watching your Git activity on your hosting service.

