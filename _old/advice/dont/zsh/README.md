---
Title: Don't *Ever* Use Zsh
Subtitle: Seriously, It will Destroy You
Query: true
---

    "Bash is the only shell scripting language permitted for executables." ([Google Style Guide](Bash is the only shell scripting language permitted for executables.))

What you are *actually* saying is,"I'm an inexperienced, uninformed, probably Mac cultist who has trouble with basic spelling and prioritizes gaudy command prompts over professional effectiveness" like [these](https://www.howtogeek.com/362409/what-is-zsh-and-why-should-you-use-it-instead-of-bash/), [uninformed](http://fendrich.se/blog/2012/09/28/no/), [people](https://www.pixelstech.net/article/1348854378-9-reasons-to-use-Zsh).

Here is a list of some of the advantages *they* claim. These four advantages are copied verbatim from [Anthony Heddings](https://twitter.com/anthonyheddings) surprisingly high-SEO-ranked write-up on [HowToGeek](https://www.howtogeek.com/362409/what-is-zsh-and-why-should-you-use-it-instead-of-bash/), a publication that has lost serious credibility for allowing this. Even though Anthony is probably a great dude these ideas are completely stupid.

> 1. Automatic cd: Just type the name of the directory
> 1. Recursive path expansion: For example “/u/lo/b” expands to “/usr/local/bin”
> 1. Spelling correction and approximate completion: If you make a minor mistake typing a directory name, ZSH will fix it for you
> 1. Plugin and theme support: ZSH includes many different plugin frameworks

## Brain-Dead Muscle Memory

All of these "advantages" are actually horrible flaws any engineer who
values productivity on *all* Linux and Unix systems will immediately
identify. The main reason you learned [Vi properly](/tools/editors/vi/)
was the same reason you should absolutely shun Zsh.

## "But `cd` is Easier"

This shows you don't even know Bash because setting a `CDPATH` does the same thing. 

## "But I Can Do `/u/lo/b`"

You can also just use tab completion which prompts when you get things wrong and is *far* faster, supported everywhere, and safer --- especially if you have `set -o vi`. Usually you are doing this in conjunction with `cd` so `CDPATH` comes into play.

:::co-dum
Go ahead and ask the Zsh person if they even know what `CDPATH` and `set -o vi` are. They won't because they are *usually* inexperienced, vocal Mac cultists who couldn't tell you the difference between `if []` and `if [[ ]]` or worse, [recommend that you use `Control-L` to clear the screen](/advice/dont/controll/).
:::

## "It Fixes My Spelling"

Do you *really* need your shell interpreter to correct your spelling? Are you seriously openly admitting you are that pathetic? If you are then nothing is going to save you, least of which some broken, non-standard, dangerous shell that Apple prefers. At least we all now know that about you. And yes, we are openly judging you. Apparently you don't prioritize accuracy *or* cross-platform productivity in your production work. "Yeah --- *sure* --- we'd *love* to have you join our pentesting or systems operations team. Umhummm."

## "It has Plugins and Pretty Themes"

> "Zsh is the new hotness. Well newer and hotter than Bash anyway, since the first version of Bash was released in June 1989, while the young and peppy Zsh was released in December 1990. In large parts thanks to the configuration “skin” oh-my-zsh, Zsh has gained a lot of popularity during the last year or so. I have used it for a few months myself and could not be happier, unless it produced chocolate ice cream  note to shell developers. 
>
> This is a guide on why you need it and how you install, configure and use it. Sometimes just with links to the relevant sites.
>
> This is written partly for my colleagues, who I think would benefit from using Zsh instead of Bash on their desktops and on our servers.

*Smacks forehead.* Hey [Fendrich](http://fendrich.se/blog/2012/09/28/no/), we all pity your "colleagues" and want to remind you that `https` (not `http`) is actually a thing in 2020. Did you *seriously* just recommend swapping out Bash on "servers" as well as desktops? 

Using *Oh-My-Zsh* is almost immediately the first advantage *most* Zsh users cite. Some of them have learned not to lead with that because any *actual* engineer is laughing his or her ass off when they list this. 

This has nothing to do with color and customization. We all love it and do it. We just don't depend on a bloated, broken, insecure plugin system for it. 

Plugins are for lazy, incompetent Mac users and junior Linux engineers who don't know how to minimally customize their own shell --- let alone write any *actual* shell code in any language. Besides, if they *did* know how to write enough shell to customize without using a plugin they would quickly realize all of Zsh's other *massive* engineering and design flaws.

## Official Shell of Apple

Zsh is now the official shell of Apple. You know. The company that *removed* the Escape key because their technical design team is just that completely brain dead. "What's Vi?" *shakes head* We *know* they are stupid because they were practically *forced* to add it back the same year Apple lost a massive class-action lawsuit for actually selling such a completely broken keyboard. 

Why did they remove it? Because they preferred a super pretty, battery-draining piece of colorful crap in its place. 

"I'm sure Apple has their reasons. Their engineers are *geniuses*! How dare you besmirch the memory of the great Steve Jobs! Apple support centers are filled with geniuses just like him!" 

Okkkkaaay.

## Apple Hates Giving Back

The *true* reason Apple picked Zsh is because they are a heartless greedy company that never gives *anything* back to the [open source](/what/open/source) community but *loves* to take and take and take claiming all the credit for itself. The idea that a license might *require* them to do the right thing terrifies them --- again, because Apple simply has no soul at all.

## Bash has Been Default for Decades

Bash continues to be the official default shell of all Linux systems. Ask yourself, do you want to learn some broken new default shell from a company that removed the `Escape` key and just wants to avoid GPLv3 or do you want to learn and use the Linux shell available on millions of systems *by default*? It's as simple as that. 

But if you are still here and care about some objective and *selfish* reasons *not to ever look at zsh* then keep reading.

## Zsh Actually Adds *Nothing*

Zsh provides no additional value over Bash, period. That is the biggest reason of all. The burden of proof is on Zsh to prove it is worth throwing out the default shell on *all* Linux systems for over two decades.

## No Exported Functions

Zsh handicaps you by preventing you from using critical Bash functionality such as exported functions (which is how all Bash completion on every Linux system is created in `/etc/bash_completion.d`). 

## No Support for Variable Name References

Another amazing thing Bash has that Zsh throws up on for no good reason
is something called *Variable Name References*.

Did you know that Bash has pointers? Well, effectively that's what they
are. They are almost unheard of in the wild despite their *incredible*
power. Here's how they
work:

```bash
#!/bin/bash

declare -n avar
some=thing
avar=some
echo $avar # prints "thing"
```

There are so many applications of this it would be a full article just
to describe them, passing names of variables to a function, saving a
`case` with an associative array lookup containing the name of a
function instead, the list could to on and on. Knowing `nameref` exists
in Bash and *don't* exist in Zsh is almost enough *by itself* to kick
Zsh off your computer.

## Dangerous Floating Point Support

Zsh tempts you with floating point math in `$((3/2))` arithmetic variable expansion --- which seems great at first --- until you realize that people will attempt to use scripts code on other Linux and Unix systems where `/bin/sh` is *not* Zsh and you scripts will *horribly and disastrously fail*. The safest way to use floating point *has always been and continues to be the `bc` command*. Always assume *any* shell script is not floating point safe.


## Encourages Bad Practices

There are a lot of little bad practices that Zsh users begin to adopt without even realizing they are destroying their ability to work on a default Linux system.

For example, some YouTubers have encouraged users to use `<` instead of `cat` to output a file confusing beginners who didn't know this was *never* supported in any other shell but Zsh. (It's actually a bug that should never have worked in the first place.)

```sh
# WRONG!
< somefile
```
Use `cat` instead.

```sh
cat somefile
```

## Failure to Meet POSIX Compliance

Zsh is *not* POSIX compliant as people claim. Use of that broken redirection suggestion is proof of that. The floating point math is another. If you are going to learn non-POSIX things they should be for the *default* Linux shell, not some broken upstart that can't decide what it wants to be. If you want POSIX use `/bin/dash` instead, which is what `/bin/sh` is linked to on all Debian-based systems.

