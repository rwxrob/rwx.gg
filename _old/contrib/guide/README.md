---
Title: Style Guide
Subtitle: Follow These When Contributing
---

Here's a summary of styles used on `rwx.gg` so that contributors can keep things *extremely* consistent. The value proposition of this knowledge app (above, say, a wiki) is this consistency, readability, and the ability to aggregate nodes into other context-specific nodes. Feel free to [open an issue](https://gitlab.com/rwx.gg/README/-/issues) if you feel that anything violates this consistency. Thanks.

## Writing Style, Grammar, and Usage

Except where otherwise noted below, the [RWX.GG]{.spy} project follows [The Gregg Reference Manual](/reviews/books/gregg) for all written content.

## General Rules

- Keep main template at `/assets/templates/main.html`

- Keep top template at `/assets/templates/top.html`

- Keep bottom template at `/assets/templates/bottom.html`

* Have the main template begin with `${top()}`

* Have the main template end with `${bottom()}`

* Keep main style sheet at `/assets/styles.css`

* Add node local template override to `<node>/template.html`

* Keep node local style overrides as `<node>/styles.css`

* Never include HTML or CSS of any kind within any `README.md` file.

* Use any Pandoc Markdown whatsoever but limit to [Basic](/lang/md/basic/) or [Light](/lang/md/pandoc/light/) when possible.

* Never use `H1` (`#`) headings. Use `Title` in YAML instead.

* Do not use [gerunds](/what/task/) to describe tasks and actions.

* Use questions for all `/what/` terms.

* Use imperative verbs in IDs and instructional steps.

* Avoid parentheses at all costs. Use dash, comma or separate sentence.

* Only `README.md` files are rendered. Other markdown ignored.

* Keep nodes composable, don't include too much into one.

* Subnodes of [knowledge nodes](/what/knowledge/node/) are fine.

* Assume any node could be moved anywhere at any time.

* Only link to nodes, local files, and external URLs.

* Never link to a local HTML file.

* Use rooted directory refs (`/views/`)

* Always including the trailing slash (`/foo/` not `/foo`).

* Put local hashtag refs after slash (`/foo/#bar`).

* Use local (`./`) directory refs for local content and subnodes.

* Don't use parent (`../`) directory refs. 

* Use only full path or external links.

* Keep node directory names easy to type and remember:
    * Pick meaningful names
    * 12 lowercase characters max
    * Use abbreviations and acronyms when possible
    * First character must be letter
    * Letters and numbers only
        * No hyphens, dashes, or underscores
        * No emojis

* Use a default empty link populate line in your `.vimrc` to fill in empty links that pull up searches on Duck.com rather than not link at all. Later you can return and provide local content as time allows. This saves those learning on the site from having to type in all the links to search the Web for terms that are likely to need explanation and allows mobile users to simple tap as they go.

* Never depend on color for anything. Keep stuff black-ish and white.

* Do not put links of any kind in any heading.

* Use `Query: true` to activate link from `h1` heading to DuckDuckGo search.

```vimscript
autocmd vimleavepre *.md !perl -p -i -e 's,\[([^\]]+?)\]\(\),[\1](https://duck.com/lite?kae=t&q=\1),g' %
```

## Callout Types

Here are the different callout types that can be used by adding a `co-` to the beginning. Make sure you know what a Pandoc Markdown "Div" is. (There are a lot of examples already in the content.)

 Identifier    Description
------------   -------------------------------------------
 `faq`           Surround level two headers of FAQs
 `fyi`           Informational but not related
 `btw`           Informational and slightly related
 `fun`           Informational, related, and fun
 `yea`           Something to celebrate.
 `hap`           Something happy
 `sad`           Something sad.
 `mad`           Something mad.
 `doh`           Common problems and mistakes, trouble.
 `dum`           So dumb you'd smack your forehead
 `tip`           High value information, related ProTips
 `pwz`           Pause and pay attention, achtung

## Special `faq` Divs

In order for frequently asked question text to be included in the main meta data summations that are used for localized searches all FAQ questions must begin with a level-two heading (`##`) and should also be grouped and encapsulated in a div group as follows:

```markdown
:::co-faq

## How should I include FAQs throughout the content?

This is how

## Do I need them to always be second level headings?

Yes. Nothing forces it, but it is a good convention.

## Should I captialize the FAQ heading?

No. It should read and appear like any other sentence.

:::


```

## Special Link Types

The following link types can be added to a link as `{.<type>}` immediately after the link markdown causing the link to have extra formatting and `before` content including an emoji or special character.

 Identifier    Description
------------   -------------------------------------------
 `see`           Notable internal or external link
 `watch`         Direct link to a video to play 
 `download`      A link to a resource that will download

## YAML Headers

YAML headers contain both meta data information as well as display parameters. The headers from the main `/README.md` file can be considered global and are available to every node in the knowledge base. Header identifiers can literally be any valid YAML but the following are standard and expected when creating the `meta.json` file.

--------------------------------------------------------
 Identifier     Description         
--------------- ----------------------------------------
 `Title`        Becomes the title of the entire
                knowledge base. If not set will be set
                to `Short` if available. Overriden by
                nodes but always applies to whole base.

 `Short`        Short version of the title included in
                every page and usually combined with
                the node's own title.
--------------------------------------------------------

Table Standard YAML Headers

### Use of Case

Use initial caps for headers that are exportable and have relevance. Only initial caps will be added to the `meta.json` file.

Lowercase should be used for localized data usually considered arguments and configuration of the template itself. Use the `tpl-` prefix and whitespace to group these type of headers.

### Knowledge Nodes


:::co-faq

## Why No Side Pane Content Lists?

We want to encourage searching, not scanning. Such table of contents listings, which are popular on sites like <https://readthedocs.org> and VuePress are distracting, wasteful, and unnecessary. They distract the reader from the content and require  active removal to deal with --- when such options are allowed. They also do not deal well with frequent updates. When such indexes are required they should be *the entire node* and not something off to the side. Never forget these are not books we are writing so put all the common idioms from books out of your mind. Things like chapters and pagination and even the words "page", and "document" do not apply.

Instead we should be promoting very memorable URLs with optional shortcuts so people can memorize them quickly and go *right* to what they want later after finding it for the first time. Often those who design complex knowledge formats with massive tables of contents have long and obnoxious URLs ignoring the very *purpose* of a URL. PHP solutions are even worse with URLs that sometimes span multiple lines. ReadTheDocs, for example, rather stupidly shows the `index.html` in the name making memorization of a simple knowledge node path impossible.

:::
