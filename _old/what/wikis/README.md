---
Title: What are the Pros and Cons of Wikis?
Subtitle: Good for Consensus, Bad for Reading and Learning
Query: true
---

> "Wikis are the knowledge management equivalent of herding cats." [\@loworderbit](https://twitch.tv/loworderbit)

Today a wiki has come to be known as "a knowledge base website" where "users ... modify content directly from the web browser." ([Wikipedia](https://en.wikipedia.org/wiki/Wiki)) 

However, both the conventional wiki software and overall wiki philosophy have some pretty major disadvantages and you should think carefully before investing time in one.

## Problems with Conventional Wiki Software

Conventional wiki software served a purpose for its time and many legacy wikis are engrained into the [modern Web](https://duck.com/lite?kae=t&q=modern Web) but this conventional, legacy wiki software has largely been superseded by better approaches and tools for most documentation needs, notably [static site generators](/what/ssg/) with content written in [Markdown](/lang/md/), hosted on a [Git provider](https://duck.com/lite?q=Git provider), and delivered over a [content delivery network](https://duck.com/lite?q=content delivery network) such as [Netlify](/services/netlify/). This modern approach fulfills the same needs as conventional wiki software:

* Content storage
* Standardized content format
* Multiple author support
* Authenticated and reviewed changes

However, conventional wiki software comes with the following heavy disadvantages:

* Database dependency
* Cannot be deployed to a [CDN](https://duck.com/lite?kae=t&q=CDN)
* Changes permitted without full review
* Difficult and non-standardized change tracking
* Forced layouts and styles
* Inadequate yet overly complex content format
* Content indefinitely locked into a non-portable format
* Not beginner friendly at all

In short, avoid the use of conventional wiki software and fulfill the same needs with a [JAMStack](/what/web/jamstack/) approach based on [Pandoc Markdown](/lang/md/pandoc/) instead.

## Challenges to Wiki Philosophy

Crowdsourcing knowledge initially seems like a great idea. Everyone contributes. No one is excluded. Opinions are discouraged. The best, balanced knowledge wins. But these advantages comes with several strong disadvantages that are often overlooked.

### Falsification and Trust

This is the \#1 biggest disadvantage cited by most educators about why they penalize students for even considering wiki information. While it is likely an over zealous response, the fact remains people do regularly lie, troll, and falsify content on Wikipedia and presumably others.

The problem exists because there is no relationship of trust between *any* of the authors as well as none between the reader and collective authors. Without trust you can never provide quality content, period.

### Content is Inconsistent

Wikis often contain contradictory and inconsistent writing. One person changes one area but (for whatever reason) decides not to change the other related information. 

### Opinions are Silenced

Despite the overwhelming sentiment to the contrary, opinions are good. They are the seeds of scientific hypotheses. Books covering the same material as a wiki are allowed to contain [strong, well-researched opinions](/what/learning/opinions/) from the author. Opinions are the basis of innovation and fundamentally not something the crowd will generally agree upon --- especially ground-breaking ideas. Imagine what a group of "collaborative authors" would say about connecting two computers together in the 80s, or saying "fat is *good* for you*, or saying "the world is actually round" (Okay that's a rough one.) Point is *well-researched* opinions are not only *essential* to good content, they are the very basis of scientific discovery.

### No Voice or Audience

Ever notice some Wikipedia pages will have all sorts of words a 12-year-old would have to look up? Others are so simply written anyone could understand. This is because Wikipedia is the functional equivalent of the Borg trying to write something for everyone. Collective authoring is *always* devoid of voice, style, and (in that case) a target audience with a specific target vocabulary. Now consider all the things that are taught at a good, novice writers' seminars:

* Know your audience?
* People like stories and anecdotes.
* Help the reader relate to you.

The result is something no one really enjoys reading. If they read it it's likely because they have no other choice. This is the polar opposite of the very successful and proven [*Head First* approach](/reviews/books/hf/).

### Just Too Many Authors

In software development this disadvantage can be compared to having several developers all working on a project using different style conventions. The comparison holds for the number of people working on the project. Throwing more people at a software project often results in missed deadlines and reduction of quality because the effort to manage all the participants detracts from the core development goals.

Instead, having a core team (or a core team of committers) meets the demand for consistency in the code base. This same approach can (and should) be applied to collaborative knowledge content development.
