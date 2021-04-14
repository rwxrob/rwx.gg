---
Title: Perl Haters Need to Shut Up and Learn It
Subtitle: Perl is *THE* Industry Standard For Regular Expressions
Query: true
---

Hating on Perl has become something of a trend these days and one thing
is very consistent. Most have no idea what Perl is or why *anyone* would
use it. There are even job posts that hate on Perl as a stupid,
uninformed, moronic joke:

> "Have a healthy disdain for Perl"

This pisses me off. Let's start with what Perl is:

>  In a single language, Perl combines some of the best features of C,
>  `sed`, `awk`, and `sh`. People familiar with these languages have
>  little difficulty being productive in Perl. Perl's expression syntax
>  is very C-like. Perl uses sophisticated pattern-matching techniques
>  to scan large amounts of data very quickly. Although optimized for
>  scanning text, Perl can also deal with binary data. If you have a
>  problem on which you would ordinarily use `sed`, `awk`, or `sh`, but
>  it exceeds these tools' capabilities or must run a little faster and
>  you don't want to write the program in a compiled language such as C,
>  Perl may be the language for you.

:::co-dum
That company making jokes about Perl is also pushing Python and proprietary Java tools, uses browser tracking, links over 50 different style sheets, and has a  horribly blinding white web site design full of `<div>` soup. That says everything you need to know about their cluelessness enough to dismiss them and find a company that actually knows what they are doing. 
:::

But lucky for you this helps you separate the dogmatic, uninformed, 1x, wannabe haters and script-kiddy companies from the 10x engineers you should prefer to work with and learn from. Here's why.

## Perl Compatible Regular Expressions / PCRE

Perl set the modern industry standard for regular expressions way back in 1987. The resulting [PRCE library](https://pcre.org) in 1997 has remained a "more powerful and flexible" replacement to other regular expression engines in *most* high-level projects and languages today. Some directly use the `libpcre` library while others like Rust use a syntax that "is similar to Perl-style regular expressions."

In other words, *almost everything* now uses Perl regular expressions --- the single most important reason Perl was created in the first place.

## Perl is Everywhere

Perl is installed by default on *every major Unix and Linux system since 1997.* The only scripting languages that can claim more ubiquity than that are the shell scripting languages and `awk` --- the language Perl was *specifically created to replace*.

This is the main reason why *good* hackers still learn Perl. They know when they get onto *any* Unix or Linux system that it is likely to be the most powerful tool available without pushing anything to the target system. Ruby ain't gonna be there. Java ain't gonna be there. And if Python *is* there it's a complete crap shoot that it will be a useful version and God knows Python sucks for command-line scripting.

## Perl is *Better* Than Sed and Awk

Sorry boomers, `sed` and `awk` are dead, and `perl` really nailed the coffin shut. Unless you need to write *truly* POSIX code that will run on *any* Unix system ever made you just don't need them. Anyone telling you otherwise has no idea what they are talking about.

Once upon a time --- decades ago --- `sed` and `awk` were the *only* POSIX way to write good shell code that would work everywhere. But that was a *long* time ago. Since then the *default* Bash shell far exceeds the needs for most and if not `perl` is *always* there to take over. 

Perl was created *specifically* to provide a better alternative to the ancient defunct Unix tools like `awk` and `sed` *and overwhelmingly succeeded at this goal.* This is why Perl was created with such an intense priority on *good* regular expressions. Try working with *basic* regular expression field replacements in `sed` or `awk` requiring crazy escapes to realize how much `perl` blew them away. Why on Earth would you waste time with crappy inferior tools?

The next time someone tells you to use `sed` or `awk` instead of `perl` just ignore them and move on. They have revealed their inferior skills and relevance. 

## Stop Blaming Perl for Being So Amazing

The fact that Perl blew up in the 90s gives a lot of insight into just how amazing this `awk` replacement really is. First of all, it was never intended to become a massive systems programming language. But people decided to use it that way because the alternatives at the time were ancient Unix shell or C. Perl was the *only* scripting language that provided what people needed between those two. 

The result was that this little `awk` replacement language became *the* back-end web processing language on literally millions of computers all over the world. It would be about a decade before anything even remotely better came along, and people will *still* debate that PHP and Python were ever better.

## Python *Never* Replaced Perl

Keep in mind that Python was created from the ground-up to fit that specific niche between C and shell for middle-ware and systems automation and systems programming *from the beginning*. It was *designed* to be that from the start. Python was *never* an `awk` replacement and continues *not* to be. 

Want proof? Try to do *anything* with Python from the command line. Python is a complete and total disaster as a command-line language because it was never meant to be. The ridiculously stupid decision to use significant white space permanently handicapped it. So stop comparing Python and Perl! Doing so just proves how *incredibly* junior you are and makes experienced engineers wince at your level of ignorance. Python and Perl serve two *completely different* needs.

The fact that Perl --- a humble little `sed` and `awk` replacement --- managed to become the *dominant* web language and remained so for *more than a decade* without any contenders proves its relevance and value and that it *far* exceeds the value and usefulness `sed`, `awk` and `python` for *anything* starting out on the command line. People ought never to forget that.

:::co-tip

Here's a little taste of Perl's power. Can you figure out what this does? Here's a hint: it comes from `~/.vimrc`. Note how the reason it is hard to read has *nothing* to do with Perl itself, but the fact that it uses a very powerful regular expression substitution that would make `sed`, `awk`, or `python` cry like a baby.

```vim
autocmd vimleavepre *.md !perl -p -i -e 's,\[([^\]]+)\]\(\),[\1](https://duck.com/lite?kae=t&q=\1),g' %
```

Much of Perl's "line noise" is caused by effective use of regular expressions, not Perl. Regexs are dense and hard to understand by nature. That's not Perl's fault. In fact, Perl is the *only* language that allows the `/x` addition to allow a complex regular expression to be written out with plenty of comments and white space making it *easier* to read. 

:::

## Inferior Engineers Criticize Perl

While regular expressions can certainly be overused --- especially in production code where a substring match would do just fine --- *not* understanding the purpose and power of regular expressions is a *major* skills gap. 

Those who criticize Perl for being "read only" do so because they have inferior skill and lack either the capacity or motivation to learn it. Hating calms their egos distracting them from realizing the harsh reality of just how inferior and lazy they really are --- you know, the kind of anti-[autodidactic](/what/autodidact/) incompetent engineer you *never* want to work with or hire, let alone be yourself. You are better than that.

## See Also

* [PCRE on Wikipedia](https://en.wikipedia.org/wiki/Perl_Compatible_Regular_Expressions)
* <https://pcre.org>
* [Sed is So Dead for Regular Expressions](/advice/dont/sed/)
