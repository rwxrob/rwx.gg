---
Title: The `jq` Command-Line JSON Processing Tool
Subtitle: Yeah, It's `awk` and `sed` for JSON
Query: true
---

Every wish you could deal with JSON from the command-line as easily as with any other type of data? Well you can with `jq` which essentially combines the functionality of all your favorite shell parsing tools into one monolithic JSON tool. In fact, it is so exhaustive in what it can do that it probably pays for it a little in execution time even though it is written in compatible C code.

Given the following `pie.json` file:

```json
{
  "pie": "Apple",
  "numbers": [10, 3.141592653589, 6.673e-11],
  "likes": true,
  "nothing": null,
  "rhyme": "Little jack horner stuck in his thumb\nand pulled out a plumb"
}
```

You can integrate it like so:

```sh
echo "I like $(jq -r .pie pie.xml) pie."
```
```{.out}
I like Apple pie.
```

Or you can even do match with numbers:

```sh
echo "Can I have 2 pis? $(jq -r '.numbers[1]*2' pie.xml)"
```
```{.out}
Can I have 2 pis? Sure, 6.283185307178.
```

:::co-warn
Note that you need the single quotes around the queries that contain brackets and other characters that having special meaning to the shell otherwise.
:::

## See Also

* [Install `jq`](./tasks/install/)
* <https://stedolan.github.io/jq/>
* <https://jqplay.org/>
