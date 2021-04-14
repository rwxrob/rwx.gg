---
Title: Language Learning Challenges
Subtitle: Gamified Way to Drive Your Coding Skills
Query: true
---

These challenges are designed to be researched and coded in order to
incrementally learn new concepts and skills for any language. The
challenges require completing several combined tasks which you will need
to research and learn for your target language. You can do them all in
one language at a time or learn several languages at the same time. (Mr.
Rob recommends Bash, Go, C, JavaScript/TypeScript, Python, Ruby, and
Rust.)

Note that all of these challenges assume you will be doing them from a
Linux command line so you will first need to become proficient with Bash
and a terminal editor such as the recommended Vim. This is by design. If
you can learn to edit, compile, and run code from the command line you
can do it *anywhere* else, but the reverse is simply not true. Since
most coding these days involves Linux and the command line you will be
ahead of everyone else who doesn't learn it that way. It's a simple
fact.

::: Remember
Remember to include the name of the language in your research about the
given topics. Otherwise, your search results will often be much too
broad.
:::

* [**Hello World**](/lang/cha/hello/)  
  Create program that prints `Hello World!` when run from
  the command line or console locally.  
  [*SheBang, Standard Output, Printing, Strings*]

* [**Hey There**](/lang/cha/hey/)  
  Create a program that prints `Hey there` by default or `Hey there You`
  if `You` is passed as an argument and `Hey there You over there` if
  `'You over there'` is passed as the first argument (hint quotes).  
  [*Variables, Parameters, Arguments, Formatted Printing*]

* [**Hi You**](/lang/cha/hi/)  
  Create a program named `hi` that prints `Hi there` by default but
  prints `Hi <arg>` if there is a command line argument. If a one of two
  special names (like `Linux` and `Rob`) are passed then instead of the
  default the script should print something special like `Woah <arg>,
  you rock!` If a specific insult name (like `Dork`) is used the script
  should print `Um, no need to be rude.`  
  [*Algorithms, Condition Statements, Boolean Logic and Operators*]

* [**Nyan Cat**](/lang/cha/nyan/)  
  Write a command-line program that loops forever printing the word
  `Nyan` (or something).  
  [*Looping, Signals, Interrupts, `Control-C`, Printing to Same Line*]

* [**RGB Command**](/lang/cha/rgb/)  
  Create a command that takes three arguments for red, green, and blue
  that are numbers between 0 and 255. Combine the numbers into a vt100
  (ANSI) RGB color escape and print the escape sequence invisibly
  coloring anything printed after. Test by printing different things in
  different colors. Once that works go back and assign random default
  color values for each if less than three arguments are passed.  
  [*VT100 ANSI Escape Characters, Terminal RGB Color Escapes, Variables,
  Quoting Backslash Escape Characters*, *Random Number Generation*]

* [**Greetings**](/lang/cha/greet/)  
  Create a command line program that when run greets the user and
  prompts them to enter their name. Then read the name and print a nice
  personalized greeting using it. If the name is not entered or contains
  only whitespace then print a message about not understanding and
  prompt again for a good name. Repeat the message and prompt forever
  until a good name is entered. 


* [**Now**](/lang/cha/now/)  
  Create a program called `now` that prints a specific human-friendly
  time stamp suitable for including in an article or a blog (Monday,
  June 29, 2020, 8:28:26PM). Create *another* program called `hnow` that
  calls the first program and adds a default Markdown header prefix `##
  `. Have `hnow` accept one argument and use that instead of the default
  if set. Demonstrate using `hnow` from `vi` by combining it with `!!`
  correctly. 

* [**Comment Out**](/lang/cha/cmt/)   
  Create a program called `cmt` that reads in every line of input and
  adds a comment prefix to each line and outputs the newly prefixed
  line. Once that's complete have the program check for arguments and
  use them as the prefix instead.

* [**Do You Like Waffles?**](/lang/cha/waffles/)   
  Write a command-line program that simulates the now ancient video [Do
  You Like Waffles](https://youtu.be/UtlaTNI1TaU) prompting the user
  with questions and checking their responses. If `yes` then ask the
  next question. If `no` then print something snarky and end the
  program. If nothing is entered (or just empty spaces) then ask the
  same question again.

<!--

|('agecalc')||||
|('badgers')||||
|('basecount')||||
|('eightball')||||
|('quote')||||
|('hex2rgb')||||
|('roll')||||
|('enigma')||||
|('talltale')||||
|('replace')||||
|('story')||||
|('weather')||||
|('todo')||||
|('notes')||||
|('repo')||||
|('panmake')||||
|('serve')||||
|('torthare')||||
|('scrape')||||
|('graph')||||
|('cpick')||||
|('tabcomp')||||

-->
