---
Title: Typing / Keyboards / Keystrokes
Subtitle: Grok the Keyboard Like a Coder
Query: true
---

You can't do anything these days without being able to use a keyboard. The selection of a keyboard, however, is a *very* personal decision.

## Symbols

Technical occupations require not only knowing about the lesser-known symbols on the keyboard but being able to type them relatively quickly. Most typing games unfortunately do not take this into account. Here they with some of their slang nicknames so you can know what people are talking about.

  Symbol      Name
--------- --------------------------------------------
   `/`        Slash / Wack
   `\`        Back Slash / Wack
   `.`        Dot
   `!`        Exclamation Point / Bang
   `?`        Question Mark
   `:`        Colon
   `;`        Semicolon
   `|`        Bar / Pipe
   `^`        Carrot
   `*`        Star / Asterix / Splat
   `-`        Dash / Hyphen
   `_`        Underscore
   `~`        Tilde / Squiggle
   `=`        Equals Sign
   `+`        Plus
   `<`        Less Than
   `>`        Greater Than
   `#`        Hashtag / Pound / Octothorpe
   `%`        Percent
   `$`        Dollar Sign
   `''`       Single Quotes
   `""`       Double Quotes
   ` `` `       Backticks
   `@`        At Sign
   `&`        Ampersand
   `[]`       Brackets / Square Brackets
   `{}`       Braces / Curly Brackets / Curlies / Mustaches
   `<>`       Angle Brackets
   `()`       Parens

:::co-fyi
You might be amazed how many people actually do not even know where the semicolon is on their keyboard, let alone the backtick.
:::

Also note the use of the following for the meta keys:

* `Space`
* `Control`
* `Alt` or `Option`
* `Super` or `Command` or `Windows`
* `Shift`
* `Return` or `Enter`
* `Escape` or `Control`+`[`
* `Tab`
* `Fn`
* `Home`
* `End`
* `PageUp`
* `PageDown`
* `Delete`
* `Backspace`
* `Insert`
* `F1` to `F19`
* `1` to `0`

There are a number of special keys for sound and such. Usually these are activated by default and to get function keys to work these days you need to hold down the `Fn` key. This can get pretty tricky when combining with `Option` and `Control` at the same time, for example, when trying to pull up a pseudo-tty (`Control`+`Option`+`F1` and `Control`+`Option`+`F7` to get your graphical desktop back)

:::co-warning
Keep in mine that what appear as the same keys on a numeric keypad are *not* the same keys when subscribing to key-press events.
:::

## Accidentally Hitting Insert

One of the most frequent tough situation beginners get into is by accidentally hitting the `Insert` key on their keyboard. This switches them into overwrite mode which just writes over their other text instead of behaving as usual. Just press the `Insert` key again to see the cursor change back.

## Getting Unstuck

Everyone will eventually have their terminal start doing weird things, usually because we have typed something it does not like. It's important that you understand all terminals started out as teletype machines that would actually type stuff out onto paper rather than sending the letters and numbers to a terminal screen for us to read.

### Interrupting (Stopping) with `Control`+`c`{#interrupt}

Sometimes something will be running and we need to cancel or stop it. Maybe we forgot a quotation mark or wrote an infinite loop and need to interrupt it. To send a running program (called a [process](/what/process/)) a signal to stop itself (an interrupt) hold down `Control` and tap `c`.

### Sending *End of File* with `Control`+`d`{#eof}

Sometimes a program wants to see a special *end of file* symbol (which cannot be printed or viewed). To send one of those do `Control`+`d`.

### Unsuspending with `Control`+`q`{#resume}

Inevitably you will type `Control`+`s` on accident some day. When you do your terminal will completely freeze. That's because you have sent it a *suspend* signal on accident. To quit suspending your terminal causing it to resume output use `Control`+`q` but be careful. What you have typed since the time it was suspended will *still* be printed to the screen and processed. For that reason it is often safer to [interrupt](#interrupt) instead.

:::co-fyi
The annoying suspend stuff is left over from teletype machines that could not keep up with the incoming data to be printed and didn't have buffers large enough to hold it all while the printer caught up. So once upon a time people actually used suspend to slow the output until the teletype printer could catch up.
:::

## Discussion Questions

* What is the minimum typing speed to work in technology?
* Do I *have* to learn touch typing?
* Should I customize my keyboard?
* What is the best keyboard?
* What about Dvorak?

## See Also

* <https://typing.com>
* <https://typing.io>
* <https://zty.pe>
* <https://nitrotype.com>
* <https://www.ostechnix.com/wpm-measure-typing-speed-terminal/>
* [Klavaro](https://klavaro.sourceforge.io/en/)
