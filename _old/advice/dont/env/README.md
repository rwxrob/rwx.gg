---
Title: Don't Use `/usr/bin/env`
Subtitle: Collective Insecure Stupidity Gone Global
Query: true
---

One of the most idiotic practices ever to gain global popularity is the use of `/usr/bin/env` followed by the assumed name of an interpreter instead of a proper *explicit* path to the *exact* interpreter program to be run. Beyond the wasteful inefficiency of forcing an additional subprocess every time your script runs, forever. This insanity is based on a completely flawed assumption that the *name* of the interpreter and its *discoverable location* are consistent and secure. They are not.

The problem stems from the innocuous assumption that your script will be easier to share with others and will just do the right thing by trusting `env` to find the right interpreter executable to run. The claim is that this makes it easier to have your script "just work" on other people's systems. But this logic has several major flaws.

First, you can never trust `/usr/bin/env` to even be on the system in the first place. It is an executable itself and might be located somewhere else, albeit unlikely.

Second, the name of the interpreter might not be what you want at all. Python is notoriously bad at this. `/usr/bin/env python` is so unpredictable that you can't even count on it returning the same *major version* of Python interpreter, let alone Python itself. This is the reason [Python Virtual Environments](https://duck.com/lite?kae=t&q=Python Virtual Environments) have become so popular --- not to mention replacing Python entirely with a modern compiled language like Go as [Dropbox did](https://about.sourcegraph.com/go/go-reliability-and-durability-at-dropbox-tammy-butow) before Guido "retired."

Third, you are dangerously assuming that the interpreter you want is even an interpreter *at all*. One example might with `/usr/bin/env fish` to run a `fish` shell script. But someone may have created a `fish` command that prints fish to the screen as a screen saver instead. Even though this specific example is an illustration it shows the logical flaw in that assumption.

Fourth, *never* trust `PATH` to find an interpreter properly. This is massively insecure. It is the reason that important commands used within scripts of any kind should always *use full path names* to ensure the proper command is being used. Promoting dependency on a secure `PATH` is irresponsible in most cases.

## Use Alternatives

If you are concerned with portability there are many far safer and more efficient alternatives. 

The best of all is to write the code in a modern, cross-compiled language so there is no use of the `PATH` to find the interpreter at all. 

Second is the use of Docker to create a container with all the requirements of the script inside of it. 

Third is to simply provide a full path name and *deliberately cause it to fail* on other systems that have the interpreter in a different place. The act of correcting the path done by the person installing the script onto their system is *explicit* by design. No one should ever put anything on their system that uses any interpreter they do not expressly identify. Trusting a `PATH` is not enough when it comes to something as significant as a language interpreter.

## See Also

* [Google's Shell Guide](https://google.github.io/styleguide/shellguide.html)
