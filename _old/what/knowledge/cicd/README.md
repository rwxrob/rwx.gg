---
Title: Continuous Integration and Deployment of Knowledge Source
---

Due to the amount of frequent changes made to knowledge source and the necessity for testing to be done by humans it does not make sense to automate integration of testing. Instead create tools for auditing and checking broken links, spelling, grammar, and such that are used during the content creation process, while it is being written. 

## Single Editor

Like most traditional publications knowledge source applications benefit most from a single editor making all the final decisions and often doing the actual work of placing the content into the application. This is a stark contrast from the [wiki](/what/wikis/) approach and ensures the overall [PKA](/what/knowledge/apps/) maintains all the advantages of traditional publications:

* Coherent voice
* Single final decision maker
* High quality
* Attention to spelling and grammar
* Less redundancy

## Keep Things Static

Keeping all content pushed to the Git hosting provider and then indirectly to a service [Netlify](/services/netlify/) is essential to maintain what would otherwise be an overuse of processing power to build out the changes during the deployment process. 

## No Merge Requests

The ultra-dynamic nature of knowledge source development does not lend itself well to the merge request process --- particularly because each release must include a full meta data generation process to create the localized `meta.json` file. Content developers must coordinate their efforts in ways that are not conducive to the typical merge request process. Consider using issues instead.

## Issues as Errata

Issues can and should be opened containing corrections and suggestions that a *single* editor can then place into the main content. Since the issues themselves using the same [Markdown](/lang/md/) format used for the knowledge source itself this becomes a matter of cutting and pasting from the issue into the master source branch.
