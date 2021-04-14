---
Title: To NeoVim, Or Not to NeoVim
Subtitle: An Honest Question, But No --- Hell No
Query: true
---

NeoVim, which uses the `nvim` command, is unfortunately a popular
replacement to [Vim](/tools/editors/vim/). It provides no additional
value for most users and can actually harm your Vi/m learning progress.

NeoVim's most significant failure is not technical at all. The NeoVim
design team has demonstrated a complete lack of understanding of Vi's
core value proposition as well as a total disregard for the fundamental
Unix philosophy. The bloated, buggy result is `nvim`. 

:::co-fun
But hey, at least you can make NodeJS plugins for it.
:::

Here are the advantages people usually present when asked why anyone
should consider using NeoVim and include an explanation about why they
are actually *dis*-advantages. They include reasons [documented by the
NeoVim project itself.](https://neovim.io/doc/user/vim_diff.html)

## "Better Configuration"

NeoVim looks in `$XDG_CONFIG_HOME` for its configuration files which
means that it follows the `~/.config/...` location convention that is
now the Linux standard.

You don't care though because you *already* are maintaining your Vim
configuration in a dotfiles repo and providing symbolic links.

Besides, moving the configuration file is downright stupid given the
decades of precedent with `~/.vimrc`. The main reason you picked Vim in
the first place was because Vim is on just about everything and copying
over your configuration is a simple matter of `scp ~/.vimrc you@remote:`
and your done. Boom. No other configuration needed. Somehow the NeoVim
team didn't figure that out, probably because they are used to
essentially turning their NeoVim editor into a watered down version of
Emacs that can only run on the single system on which it is configured.

## "Better Defaults"

The second thing listed on [NeoVim's comparison
page](https://neovim.io/doc/user/vim_diff.html) is the 42 different
defaults from Vim. These are completely and totally irrelevant because
*anyone* using Vim should *always* disable all the defaults and begin
with a clean slate in their `vimrc` file. The very fact that the NeoVim
team thought that having different defaults actually matters at all
shows how disconnected that design team is from how Vim is actually used
professionally. Again, a symptom of Emacs-envy.

## "Multiple API and Plugin Support"

Vim has this as well but you should never use it, that is, unless you
want to make Vim into VSCode or Emacs or Sublime. Seriously folks, the
entire emphasis of the NeoVim project and priorities demonstrates an
utter cluelessness about the actual value proposition of picking Vim in
the first place --- the biggest being *[full shell
integration](/tools/editors/vi/how/magic/) for extensibility*, not
supporting NodeJS plugins. NeoVim has made itself into a serious joke
among those who know and use Vi/m as has been down for decades for all
the right reasons.

Having a clearer internal API is a compelling reason to consider *any*
project, but it doesn't hold any weight with the end user. Think about
[Wayland vs X](https://duck.com/lite?kae=t&q=Wayland vs X), for example.

## "Changed Features"

Every single feature that has changed is ridiculously irrelevant to
anyone who actually knows how to use Vi. Things like `json_decode` are
just silly when commands like [`jq`](/tools/jq/) exist. They even
renamed `viminfo` to `shada` for nothing but vain not-invented-here
reasons. And Lua and Python support? Pffff. Please. You can be glad you
learned to use Vi/m correctly and without a bunch of unnecessary bloat
that would directly affect your performance on *every* other system with
Vi while *diminishing* your ability to actually use your most powerful
tool, the shell in which Vi/m is running.

What is even worse is that NeoVim has actually corrupted the *expected*
behavior of current Vi and Vim options and Ex commands. This is simply
dangerous and stupid. It creates an unnecessary rift in muscle-memory
that you *never* want to burn into your brain.

## "Missing Legacy Features"

The `if_perl` has been dropped. Nothing screams "we are all morons" more
than dropping Perl support from something that has had it for 2 decades
just because you buy into the trendy Perl-hate. Perl has the world's
most powerful regular expression engine as has been proven over and
again by its integration *into every other language* with regular
expressions including Python, NodeJS, even Bash. To blindly remove
support for syntax used by Vim users that integrate Perl (albeit
foolishly when they should have just integrated shell command filters
instead) is just plain clueless and downright stupid.

## "Removed Features"

NeoVim *removed* several core tools used regular by Vim users for
seriously important use cases:

* `ex`
* `view`
* `vimdiff`
* `:shell`

Again, incredibly inexperienced decisions from people who *never
actually learned to use Vim* for anything significant in the first
place. The fact that they removed `:shell` completely confirms they
don't value shell integration which is the basis of all of Vim's
[magical power](/tools/editors/vi/how/magic/).

Come on, they didn't need to remove `:smile`?! That's just low.

:::co-fun
The laughter from the Vim team behind the scenes must be so hilariously
loud given how ridiculously superior [Vim
8.2](https://duck.com/lite?kae=t&q=Vim 8.2) is to *any* NeoVim
script-kiddy release.
:::

## "More Accessible Team"

Yeah, just no. No one should believe that for a second. If anything, the
Vim team is simply more capable and discerning than the NeoVim team and
by the looks of the NeoVim project priorities this seems to be
objectively true.

## "External Plugins in Separate Process"

This might be true, but again it is irrelevant because no one should be
that dependent on plugins in the first place. In fact, making plugins
running in a separate process and thread is a symptom that they are
doing the internals completely wrong by allowing plugins to take a far
greater position in the overall runtime. This is not Microsoft Visual
Studio we are using here.

## "Better Support for Alacritty"

While Alacritty is an amazing project this claim makes absolutely no
sense whatsoever. It's an editor running in a terminal. Making a
decision to use a ubiquitous terminal editor based on which terminal
application you use it hilariously stupid and irrelevant. That's like
arguing back with "Well I use Vim because it works better with Xterm2."
Who cares?!

## "Gives Vim Healthy Competition"

Giving Vim competition certainly cannot hurt. But anyone who actually
knows how to use Vi/m and gives NeoVim's list of differences a solid
review will realize NeoVim is absolutely no competition at all. Vim is
installed on literally millions of Linux systems dating back more than
three decades. NeoVim *might* be installed on maybe 10,000 systems tops.
There is *zero* competition.

## "It Has Panes"

So does Vim, but you should never use them unless you are forced. They
are an unnecessary and useless [Vimism](/tools/editors/vim/vimisms/)
that are better replaced with learning to use TMUX panes or even
`screen` windows instead since they work for *any* application not just
Vim.

## Bugs? Stability?

It has been reported that NeoVim contains a lot of bugs and stability
issues. That is no surprise at all given the massive, unnecessary
scope-creep the NeoVim project team has *deliberately* chosen to
maintain. 

Any serious professional understand the importance of the Unix
philosophy of doing one thing well and making sure it integrates with
everything else. NeoVim --- with all its unnecessary bloat --- is a
serious departure from that philosophy and will continue to remain
unstable and buggy because of it. It really is a shame that the NeoVim
development team simply cannot see that.

:::co-faq

## Dude, why so harsh on NeoVim?

This review started out much more objectively. But as each point of
difference stated by the team itself was examined, the level of
hilarious collective cluelessness exceeded most authors' ability not to
completely roast it. There is simply nothing good to say about NeoVim at
all. It really is just that bad. 

Put comically, the NeoVim projects seems like a bunch of people got
their feelings hurt trying to get their stupid, bloated ideas accepted
into the Vim project so they started their own while waving their pretty
logo and chanting "We're more open. We're more open." 

Fact is. The clueless cult of over-engineering bloat makers [never knew
how to use Vi in the first place](/tools/editors/vi/how/magic/). Just
ask them what `!!` does from command mode in Vim. Most can't even tell
you. Their too busy dreaming up more ways to overcome their Emacs-envy.
They don't know Vi and frankly don't even understand Unix.

But hey, don't take all of this *too* seriously. As human beings every
one of us deserves respect even if our ideas are ridiculously stupid and
uninformed. Attack ideas, not people.

:::

## See Also

* <https://neovim.io/doc/user/vim_diff.html>
