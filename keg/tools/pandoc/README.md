---
Title: Pandoc, Universal Document Converter
Subtitle: World's Best Tool for Knowledge Source
Query: true
---

Pandoc is not only the world's most comprehensive [Markdown](/lang/md) tool it allows you to convert documents *from* any format *to* any other format. That's not an exaggeration. In fact, Pandoc is widely known to be the leading tool used by those creating textbooks and other academic writing, but it is also great just for rendering a simple, static web site quickly or taking notes in class because of its powerful [LaTeX markup](https://duck.com/lite?kae=t&q=LaTeX markup) support for math notation and more.

## Installing Pandoc

Getting the latest Pandoc is a little tricky on Debian-based systems because the system packages don't seem to be kept up to date very well. Instead go [download the Debian package](https://github.com/jgm/pandoc/releases/latest) and install it.

If you know how to install and run scripts, here's a Bash script to always get the latest Debian package from the Pandoc project GitHub repo:

```sh
#!/bin/bash

loc=$(curl -o /dev/null -sIw "%{redirect_url}" 'https://github.com/jgm/pandoc/releases/latest')
vers=${loc##*/}
deb=pandoc-$vers-1-amd64.deb
uri=${loc%/tag*}/download/$vers/$deb

echo "Downloading $uri"
curl -sL "$uri" -o "/tmp/$deb"

echo "Installing /tmp/$deb"
sudo dpkg -i "/tmp/$deb"
```
