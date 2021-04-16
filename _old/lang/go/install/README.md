---
Title: Install Golang
Subtitle: Always Have the Latest 
Query: true
---

First check to see that you don't already the right version of Go installed:

```sh
go --version
```

The current version (as of April 23, 2020) is 1.14.2.

You can certainly following the [installation instructions](https://golang.org/doc/install) on the Go page. Or you can add this little script to manage your Go version installation like any other package on your system. Keep in mind that although this PPA is non-standard and not supported by Go officially it is used by thousands.

```sh
#!/bin/sh

sudo add-apt-repository ppa:longsleep/golang-backports
sudo apt update
sudo apt install golang-go
```

You'll need to do your own research for other non-Debian versions of Linux (but you are probably used to that if you do).

:::callout-fyi
If you absolutely have no other option you can always just create a Golang project on [REPL.it](https://repl.it).
:::
