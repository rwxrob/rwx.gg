---
Title: Should I Use Multiline Paragraphs in Markdown?
Subtitle: Strong Arguments For and Against, But Yes
Query: true

Summary:
  Longer terminal widths from wider monitors and smaller widths from
  TMUX panes have brought the question of whether to use single line
  paragraphs or multiple lines to many Markdown authors minds. For those
  creating and maintaining knowledge bases with hundreds of pages of
  knowledge source this is a big deal. So what should you do? As usual,
  the answer is "it depends" but generally you still should. Here's why.

---

## Arguments for Multiline

* Historically multiline has been much more popular and is the standard
  for many formal text document formats such as the IETF RFC
  specifications.

* Multiline wrapping can be synchronized with rendered text width --- 72
  `textwidth` to 700px for example. While no rendering style should ever
  depend on such a thing, it is helpful to avoid odd outliers in a
  bulleted list, table, or paragraph orphans that, while strictly
  speaking are okay, would still be nice to avoid.

* While it may seem counter-intuitive to hard wrap wasting the maximum
  width available on any terminal or TMUX pane, long lines are much more
  difficult reading while editing documents --- particularly if the one
  writing them is working with them for long spans of time.

* Seeing ugly hard wrapping when a terminal or TMUX pane is not wide
  enough provides an immediate visual queue that it needs to be widened
  in order to display the document correctly. This is *particularly*
  important if your Markdown has a lot of text in the YAML header.

* Consistent line lengths make line counts a more accurate measure of
  overall document length (even though words are probably still the best
  measure.

* The `vim-pandoc` plugin makes for drop-dead simple formatting. Just
  make sure to `let g:pandoc#formatting#mode = 'hA'`{.vim} and `let
  g:pandoc#formatting#textwidth = 72`{.vim} so that formatting is not
  something you even have to think about.

## Arguments for Single Line

* Paragraphs become one single long line allowing them to be easily
  copied and pasted with a single `dd` or `yy` in Vi/m.

* None of the terminal or pane width is wasted.

* Smaller panes --- such as with TMUX --- make hard wrapped paragraphs
  look *very* ugly. But this is also an indicator that your pane should
  be widened.
