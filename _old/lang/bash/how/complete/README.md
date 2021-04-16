---
Title: Easy Bash Programmable Tab Completion
Subtitle: A Tab is Worth a Thousand Words
Query: true
---

Bash tab completion is perhaps one of the most overly and unnecessarily complicated topics in all of Linux. The best documentation on it remains the Bash man page itself, but even that is complex and gives almost *zero* attention to the humble `complete -C foo foo` approach even though implementing tab completion using `-C` instead of `-F` and others means that the code that prints the possible completions is *also* the code that executes the completed command line. Bundling completion *with* the command is simple, concise, and more sustainable. There isn't an additional script to maintain and install (typically in `/etc/bash_completion.d`). Users just add the `complete` to their Bash configuration and works. Those making software packages can still install a separate completion script, but it becomes a single line instead of an overly complicated Bash function. Moreover, your code --- with encapsulated, detectable completion mode --- will be ready to integrate into any number of other shells that support it besides Bash. When you consider these advantages its rather surprising everyone doesn't use this method probably because most just don't know about. Now you do.

Here are some simple examples written in Bash to get the idea. The `week` script just prints out the first argument, which can be any days of the week. (Normally this would be simple enough to just implement with `complete -W` but this is to illustrate how it works.)

## Setup a Test Script

First you'll need to find a place to create a simple Bash script, preferably not in your `PATH`.

```sh
cd /tmp
touch week
chmod u+x week
```

Then go ahead and execute the `complete` command in the currently running shell. Note that we are using the local directory path just for our test.

```sh
complete -C ./week ./week
```

You can already test this by running `./week` and tapping `Tab` a couple times to see that it does nothing, having taken over the default Bash tab completion that uses files and directories.



