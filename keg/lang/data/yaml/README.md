---
Title: YAML Ain't Markup Language
Subtitle: Human Writable Structured Data
Query: true
---

```yaml
---
pie: Apple

numbers:
- 10
- 3.141592653589
- 6.673e-11

# Does this person like pie?
likes: true

# Actually nothing
nothing: null

rhyme: |
  Little jack horner stuck in his thumb
  and pulled out a plumb
---
```

YAML is a backward compatible [JSON](/lang/data/json/) replacement for use when a [structured data format](https://duck.com/lite?kae=t&q=structured data) must be *written and maintained by humans*. However, YAML is *horrible* when a program needs to *write* it due to the human-friendly use of comments, tabs, whitespace, line returns, and such that make it so popular with those creating configuration files and [flat-file database applications](https://duck.com/lite?kae=t&q=flat-file database applications).
