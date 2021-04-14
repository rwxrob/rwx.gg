---
Title: Site Reliability Engineering
Subtitle: Google's Poorly-Conceived Title for the Obvious
Query: true
---

Google came up with a poorly-conceived title for what *good* [System Administrators](/jobs/sysadmin/) had been naturally doing in their own companies during the same time period in the early 2000s. Then they got O'Reilly to publish what is effectively a buzz-wordy, bullshit blog about it and call it a book. Lucky for us. They had the "courage" to "come out of the shadows" and grace us with this \$40 collection of *well-duh* insights.

Meh.

It's no surprise there's nothing particularly Earth shattering in this thing. They just happened to write about it and tout themselves as innovative just 'cuz. I love the ring-kissing forward thanking Google for having such "courage" to do what everyone else already naturally does: share their innovations for the betterment of the world, not prey on them for their personal lives and preferences. 

Despite this non-event of a "book," Google's bullshit SRE non-title has taken over the industry making a *System Administrator* automatically considered *less-than* an SRE even though Google's title makes absolutely no sense across the industry. This is the sad, unfortunate state of the operations and administration industry today. The worst part is that you actually have to know what it all means to even be considered a serious player in the field these days.

\*sigh\*

This collection of annotations at times agrees with the obvious and others laughs out loud at what everyone else had already decided was *the* way to do systems administration worldwide. 

:::co-care
Your first caution should be this. This is just *one* company's way of doing it. *Never* just apply what one company says to do. Evaluate your situation independently and make your *own* plan and strategy. Ironically Google would be the first to agree with this advice, at least on paper when their being quoted and published for cash.
:::

## Forward

> Google grew at a time when the traditional role of the system administrator was being transformed.

This is *very* true and innovations happened all over the industry quietly behind the doors of *all* corporations. IBM created its *Virtual Server Administrator* initiative within Global Services around that same time if not before Google. There were a lot of layoffs of traditional sysadmins who had not demonstrated an ability to do the extra stuff which minimally included the ability to code and create automations.

> wizard culture

Included the BOFH comics making fun of the elitism among system admins and the developers who had to bow down to get their applications deployed, which came from the previous systems operator culture where you were luck to get your stack of punch-cards loaded.

> Google's experience was unique.

Sure it was, but so was everyone else's. Google's wasn't particularly better, just got more coverage and attention.

> tribal nature of IT culture

Dogma is death. Google realized that as a company and responded well for sure.

> These articles are not rigorous, academic pieces; they are personal accounts

In other words, a bunch of stuff that could just as well have been blogged instead of stuffed in a book for \$40.

> The most impressive thing about this book is its very existence.

Pfffhahahahahahahaha. That level of uncheck hubris needs to be punished as much as possible. Seriously Mark, you are so completely off base assuming that a "culture of ask no questions" has risen when all you have to do is do a bit of basic research to realize the culture has *completely* changed and if anything *overshares* their *insights*. This is completely unfounded.

:::co-rage
Bowing to the cult of Google is *really* putting the industry and our world in jeopardy. Putting Google on a false-god's pedestal is seriously hurting our industry. Smaller innovations from companies motivated by far more than the 90% ad revenue income are *way* more compelling. Maybe that is why Harvard Business Review is focused on GitLab and *not* on lumbering, acid-tripping Unicorns like Google.
:::

## Preface

> and sometimes, our task is figuring out how to apply existing solutions to new problems.

This is the *last* thing this author lists implying that looking for existing solutions *not* the first thing Google thinks. This demonstrates classic not-invented-hear syndrome. Google seethes with that disease. It is exactly by projects like Fuscia will crash and burn a horrible death. Google is just like Apple and IBM teaming up to call their *massive fail* of an operating system "Taligent" (talented + intelligent). Does anyone need to understand why such hubris in the technical world might just destroy the world? 

So here's the take away:

* SREs are *engineers*.
* SREs are focused on *reliability*.
* SREs are focused on *services*.

> The "site" in our name originally referred to SRE's role in keeping the *google.com* website running, though we now run many more services.

:::co-rage
Then change the stupid title! How this broken title became industry standard is absolutely beyond rational thinking. Why not change that S to *services* instead of *site* since they are different?
:::

Reading that *Preface* was almost an entire waste of time. The only valuable take away is that "site" is "services" now and that the entire title is biased toward the "world of web services", which is rather shallow thinking indeed. The future is *much* larger than the current Web we use today.
 and that the entire title is biased toward the "world of web services", which is rather shallow thinking indeed. The future is *much* larger than the current Web we use today.

> [W]e are distinct from the industry term DevOps, because although we definitely regard [infrastructure](/what/infra/) as code, we have *reliability* as our main focus.

So all those postings that have both DevOps and SRE are covering two titles instead of one, nice. As if knowing what specific skills needed for the nebulous SRE title wasn't bad enough.

> Ultimately more reliability-oriented software and systems engineering is inherently good.

Obviously. Why waste a sentence on that?

> [T]here probably already is someone in your organization who is doing SRE work, without it necessarily called that name, or recognized as such.

Ya think?

> Margaret Hamilton

Oh great. So the person who gave the world the term *Software Engineer* (which continues to anger *actual* engineers working in mechanical, physical, chemical, civil and other realms) is the person first identified as actually being a *Site Reliability Engineer* even though she called herself specifically a "software engineer." This is such absolute bull crap.

> Welcome to our emerging profession!

This entire account paints everyone but Margaret Hamilton as blundering buffoons who never thought to think about planning for the inevitable failures that astronauts and others would make. This is a downright disgrace to the massive effort from multiple engineers providing multiple redundant systems for just such occasions. Cherry picking one anecdote and demonizing a fictional "higher up" is just irresponsible. And what is with the "Welcome" as if this is the first time anyone has done *this new profession*? The level of egocentricity is downright disgusting. Sorry Google, you did *not* invent this profession and you cannot claim the woman to made a name coining the term *software engineer* is your first club member. You have no idea what is going on in the world around you outside of Google.

Apparently it took "2 years" to write this "collection of essays."

## Chapter 1, Introduction

> Hope is not a strategy

The point about the cost and delay of change management is well taken.

> The outcome is a pathology.

This is the lead tech guy at Google stating that having a distinction between systems administrators and developers is a "pathology" that obviously needs to be cured.

This guy is describing the entire DevOps thing, and he wasn't the first. Interesting because in the Preface the explanation that DevOps is *not* the realm of an SRE due to its "services focus". So which is it?

> SRE is what happens when you ask a software engineer to design an operations team.

What do *you* think would be the first requirement every member of an operations team designed by group of programmers should be!? Ding ding! You guessed it *programming*! Oh my *God* the *innovative* thinking is over 9000! 

Again, this is *not* unique to Google. In companies around the world, small and large, software engineers have been creating effective, efficient, reliable ways to deploy while serving the role of operations. Perhaps the simplest is an immediate ability to release *any* change instantaneously with *zero* overhead to that release.

The serious danger in this whole way of thinking is the implication like saying "we don't need no stinking system administrators" which is obviously a moronic conclusion but directly implied by the entire SRE mania. You don't see these SREs providing hardware specs, testing power supplies, and scanning for back doors both in the hardware and applications and network layers. An SRE *cannot* do it all, yet they are championed as just that. The SRE title is similar to the Full-Stack Web Developer title. It means nothing because of its massive scope. The only people who like it are the increasing number of IT managers who refuse to access that work is about filling *roles* and *not* titles and that we need more fine-grained role identification and less over-arching, meaningless titles with books explaining them badly.

> By far, UNIX system internals and networking (Layer 1 to Layer 3) expertise are the two most common types of alternative skills we seek.

See, *those* are the system administrators and network administrators. You are just requiring them to have "software engineering" experience as well.

So the "SRE team" is comprised of 50-60% "software engineers" and "40-50% ... alternative skills" (which is just code for system engineers and network engineers). Congratulations you brilliant Googlers you've (re)invented the NOC and the SOC, put them into a new team called SRE and told the world about your wonderful discovery. OMG, the innovation is amazing!

They even "track" them individually as different roles within the SRE group.

> ... have a team of people who (a) will become quickly bored by performing tasks by hand and (b) have skills ... to replace their manual work ... even when solution is complicated

This overemphasis on complicated outcomes has to be one of Google's greatest flaws and will eventually take down the company. Fuscia is a perfect example of technical how hubris results is absolutely crappy software and systems.

> (DevOps 2008, SRE 2016) One could equivalently view SRE as a specific implementation of DevOps with some idiosyncratic extensions.

Finally an explanation of the difference.

> In general an SRE team is responsible for 
> 
> * availability,
> * latency,
> * performance,
> * efficiency,
> * change management,
> * monitoring,
> * emergency response,
> * and capacity planning.
> 
> of their service(s).

***Left off on Page 8. Sorry I couldn't stand to read any more of this trash.***
