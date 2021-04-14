---
Title: Dark Mode GPU PDFs with MuPDF
Subtitle: A Fast and Safe Viewer with Vim Mode
Query: true
---

The *MuPDF* viewer it wicked fast on all fronts, supports dark mode (inversion), and has [Vi](/tools/editors/vi/) bindings built in. However, it is maintained by Artifex, not hosted on GitHub or GitLab, and has ancient versions in most Debian-based distros. 

:::co-tip
Only the latest has a working `-I` inverted dark mode. It's worth the time to compile it and compiling will be a good learning activity.
:::

## Building from Source

First get all the dependencies. You should be doing this from a device that has a graphics driver supporting OpenGL, which you probably do.

```sh
sudo apt install -y xorg-dev freeglut3-dev
```

```{.out}
Reading package lists... Done
Building dependency tree
Reading state information... Done
The following NEW packages will be installed:
  freeglut3-dev xorg-dev
0 upgraded, 2 newly installed, 0 to remove and 146 not upgraded.
Need to get 128 kB of archives.
After this operation, 812 kB of additional disk space will be used.
Get:1 http://archive.ubuntu.com/ubuntu bionic/universe amd64 freeglut3-dev amd64 2.8.1-3 [124 kB]
Get:2 http://archive.ubuntu.com/ubuntu bionic-updates/main amd64 xorg-dev all 1:7.7+19ubuntu7.1 [4,300 B]
Fetched 128 kB in 1s (171 kB/s)
Selecting previously unselected package freeglut3-dev:amd64.
(Reading database ... 326158 files and directories currently installed.)
Preparing to unpack .../freeglut3-dev_2.8.1-3_amd64.deb ...
Unpacking freeglut3-dev:amd64 (2.8.1-3) ...
Selecting previously unselected package xorg-dev.
Preparing to unpack .../xorg-dev_1%3a7.7+19ubuntu7.1_all.deb ...
Unpacking xorg-dev (1:7.7+19ubuntu7.1) ...
Setting up xorg-dev (1:7.7+19ubuntu7.1) ...
Setting up freeglut3-dev:amd64 (2.8.1-3) ...
```

Now you have to get the code. 

:::co-stop
If you haven't already done so [setup your local repos](/services/gitlab/repos/).
:::

Then create a directory to contain the `mupdf` repo.

```sh
cd ~/repos
mkdir git.ghostscript.com
cd git.ghostscript.com
```

Now to a *recursive* clone to not only get the Git repo but also all the *submodule* repos within it.

```sh
git clone --recursive git://git.ghostscript.com/mupdf
```

This will have a *lot* of output.

Change into the new `mupdf` directory and then force the submodules to update as well (if they didn't already).

```sh
cd mupdf
git submodule update --init 
```

Again a lot more output.

Now you just need to build it. You can either make and install it for the whole computer and all users on it, or you can make a local copy. Just change the `prefix` to whatever you would like. Keep in mind that installing it for the system requires `sudo`, which is what this example shows.

:::co-fyi
The [LFH](/tools/linux/lfh/) calls for putting all installed user applications into `/usr/local`.
:::

```sh
sudo make prefix=/usr/local install`
```

This will produce a *massive* amount of output and often pause and possibly look like it is failing even though it is not. Just be patient. You are compiling a rather large application written in C and C++ from scratch. Take a moment to be thankful for precompiled binary packages for most things.

When the compilation finishes you will see two new binaries in `/usr/local`: `mupdf-gl` and `mupdf-xl`.  You can either use them this way or create an alias, functions, or small script to call them. 

For example, to use the fast `mupdf-gl` with inversion on every time you type `mupdf` try adding the following alias to your Bash configuration.

```sh
alias mupdf='mupdf-gl -I'
```

Perhaps you want to consolidate all your PDFs into a common `PDFDIR` that keeps them organized.

First create the directory or project repo. Then export the `PDFDIR` in your Bash configuration something like this.

```sh
export PDFDIR=$HOME/repos/gitlab.com/rwxyou/pdfs
```

Now you can create a script like the following and put it in your `dotfiles/scripts` directory or someplace else in your `PATH`.

```sh
cd ~/repos/gitlab.com/rwxyou/dotfiles/scripts
touch pdf
```

Here's one example, tweak it as you like.

```sh
#!/bin/sh

if [ -z "$PDFDIR" ];then
  warn 'The `PDFDIR` env var not set.'
  exit 0
fi

list () {
  find "$PDFDIR" -name "*.pdf" -exec basename {} \;
}

usage () {
  usageln 'list|usage|<pdf>'
}

if [ -n "$COMP_LINE" ]; then
  prefix=$(echo "$COMP_LINE" | cut -d " " -f 2)
  list | grep ^$prefix
  exit 0
fi

case "$1" in
  list) list; exit ;;
  usage) usage; exit ;;
esac

# actually we have a pdf name, not a command
# if gotten this far
pdfpath="$PDFDIR/$1"

# always open using GL in dark (inverted) mode
if [ -r "$pdfpath" ];then
  mupdf-gl -I "$pdfpath"  &>/dev/null &
fi
```
