---
Title: Knowledge Node
Subtitle: Finite Automaton of Knowledge States
Query: true
---

An RWX *knowledge node* is simply some amount of knowledge at its most granular level. If you imagine your brain as a [cognitive](https://duck.com/lite?kae=t&q=cognitive), [finite state machine](https://duck.com/lite?kae=t&q=finite state machine). Then a knowledge node is one finite state, which links to may others as well. Each time you follow a link you change the state of your cognitive circuitry. A [knowledge base](/what/knowledge/base/) is therefore cognitive programming you can load and run. Your brain processes the same nodes following different paths and KNode states.

The practical implementation of a knowledge node manifests as a simple subdirectory containing a single [README.md](/what/knowledge/readme/) within the subdirectory. Knowledge nodes may contain any other assets --- including other subnodes --- but should not contain other [Markdown](/lang/md/) files unless those files are examples or other reference material.

:::co-warn
***Only the main `README.md` Pandoc Markdown file will be rendered as HTML.***
:::

## Limited Size

Like a good program, a knowledge node should not be longer than a 2-3 screens of text (about 300 lines of [Markdown](/lang/md/)). This is to limit the required computing resources to display a single node and afford the best opportunity for node aggregation through hyperlinking.

## Aggregation, Not Composition

A single knowledge node can be logically aggregated from many others. [Aggregation](https://duck.com/lite?kae=t&q=Aggregation) takes the simple form of hyperlinking to other nodes. Even though the linked content does not *actually* become a physical part of the linking node the reference to other nodes *is* a critical dependency since the linking node cannot exist without the linked node.

:::co-fyi
[Aggregation](https://duck.com/lite?kae=t&q=Aggregation) is a critical dependency from one thing to others whereby if the thing depended on goes away it breaks. [Composition](https://duck.com/lite?kae=t&q=Composition) is when the that dependency goes both ways and if *either* thing goes away *both* break.
:::
