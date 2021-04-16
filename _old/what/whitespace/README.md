---
Title: What is Whitespace?
Subtitle: Tabs, Spaces, Blank Lines
Query: true
---

*Whitespace* is just the empty space around stuff including indentation. Using whitespace properly is the single best way to make your code readable and avoid [syntax](/what/syntax/) errors. It is so important that some languages have designed it into the [language itself](#significant-whitespace).

## Significant Whitespace

Several languages --- notably Python and Haskell but even Nim --- have included whitespace as a syntactical part of the language. This means that if your indentation fails your code won't even compile or run. While this appears like a good idea initially it quickly falls on its face when working with large codebases where code must be cut and paste in large blocks between nested constructs such as functions and if statements. A good editor is not enough to make up for bad language syntax design.

Allowing lines to be combined (as most other languages in the world allow) provides the syntactical flexibility to grow the syntax as needed as has been the case with modern JavaScript, now more than 20 years old and still the most used language on the planet.

*Not* requiring whitespace as a part of the language syntax also allow command-line one-liners to exist. Bash, Perl, Ruby, Awk all allows for this. Python does not, making it a horrible language pick for quick shell scripting.

:::co-tip
Avoid coding in any language that requires significant whitespace unless you are absolutely required to. Remember CoffeeScript? Jade? Be glad they died for all the right reasons.
:::

:::co-fun
If you want to get a good laugh from the misfortune of a team that actually made the idiotic decision to include significant whitespace in a major language check the Python team choking on their decision while seriously struggling to add multi-line anonymous functions to the language. The forced indentation stopped them dead cold. The mailing list thread is absolutely hilarious as everyone basically just skips the issue and pretends like it never happened as if to say "We didn't need multi-line anonymous functions anyway, one line is fine" rather than concede defeat. 

Meanwhile, JavaScript has been able to grow its modern syntax into one of the most elegant to have existed with fat-arrow functions and first-class function definitions with the assignment operator denoting *exactly* what is happening and how they should be viewed, as just another data type that can be assigned and passed as arguments. This elegance cannot be expressed at all in the Python language because of significant whitespace.
:::
