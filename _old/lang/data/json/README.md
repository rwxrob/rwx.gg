---
Title: JavaScript Object Notation / JSON
Subtitle: World's Most Prominent Structured Data Language
Query: true
---

```json
{
  "pie": "Apple",
  "numbers": [10, 3.141592653589, 6.673e-11],
  "likes": true,
  "nothing": null,
  "rhyme": "Little Jack Horner stuck in his thumb\nAnd pulled out a plumb"
}
```

*JavaScript Object Notation* is an open standard, human-readable file format used for structured data and data interchange usually over the Internet. It is widely regarded as *the standard* for such data exchange. The format itself includes four simple data formats:

* strings (one or more [Unicode](https://duck.com/lite?kae=t&q=Unicode) characters, but no line returns)
* numbers (including scientific notation)
* `true`/`false` (boolean)
* `null` (nothing)

Simple formats can be used in one of the two compound collection types:

* objects (key:value pairs) 
* arrays (simple lists of values)

This simplicity allows JSON to be used universally by all [programming langauges](/lang/) to read and write structured data.

## Human Readable, Yes. Writable? Not Really.

JSON was made to fast data storage and interchange between programs, not really humans. Being human-readable is a plus because humans can quickly see what is going on. But humans have a harder time writing it because it lacks any support for line returns. All string data is one big long line. For this reason [YAML](/lang/data/yaml/), which is 100% backward compatible with JSON, is preferred by the industry for data that must be *written and maintained* by humans.

:::co-fyi
JSON saved the world from [XML](/lang/data/xml/). Be very grateful.
:::

## See Also

* [The `jq` Tool](/tools/jq/)
