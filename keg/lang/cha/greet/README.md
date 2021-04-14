---
Title: '"Greetings" Challenge'
Tags: easy
---

Create a command line program that when run greets the user and prompts
them to enter their name. Then read the name and print a nice
personalized greeting using it. If the name is not entered or contains
only whitespace then print a message about not understanding and prompt
again for a good name. Repeat the message and prompt forever until a
good name is entered. 

## Requirements

* Must include a *safe* shebang line if a script
* When `./greet` then print `Greetings! What's your name?`
* Read user name from one line of standard input
* When `You` entered then prints `Welcome, You! It's good to meet you."`
* When name is empty then print `Sorry, didn't catch that.`
* When name is just spaces then print `Sorry, didn't catch that.`
* When name is tabs then print `Sorry, didn't catch that.`
* Prompt again forever until name is entered if empty
* Use an infinite loop design breaking out when name is good.

## Bonus

* Write the algorithm first as comments using natural language
* Validate using language test framework.
