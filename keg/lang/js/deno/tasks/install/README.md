---
Title: Install Deno for JavaScript/TypeScript Development
Query: true
---

Since Deno is so new it still uses a simple shell script to install. They recommend to download and run the shell script directly but you know better than to *ever* download and run *anything* without ensuring that the file you are downloading has not been tampered with by a man-in-the-middle attack. To do this we have to break apart the steps and validate the checksum of the file to ensure it is the right one. 

First download the file

```sh
curl -sSL https://deno.land/x/install/install.sh 
```

Then you can put the following into a file like [checksums.txt](checksums.txt).

```sha
cc99de2df5af20fd65b287293d9ca89121044cd5cb832d3c7f52175d0f7fc6ab  install.sh
```

After that just check that the file hasn't been tampered with.

```sh
sha256sum -c checksums.txt
```

This should print something like the following to ensure your `install.sh` file hasn't been tampered with.

```{.out}
install.sh: OK
```

Only now is it safe to run the `install.sh` script which will place `deno` into `~/.denu/bin` by default.

```sh
chmod +x install.sh
./install.sh
```

```{.out}
################################ 100.0%
Archive:  /home/rwxyou/.deno/bin/deno.zip
  inflating: deno
Deno was installed successfully to /home/rwxyou/.deno/bin/deno
Run 'deno --help' to get started
```

You're all set.
