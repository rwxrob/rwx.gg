---
Title: Set Up Vi/m Dotfiles Configuration Subdirectory
Subtitle: Single File, 100% Portable, No Blocking Dependencies
Query: true
---

Make sure you have setup your [dotfiles](../) repo.

Change into your `dotfiles` repo.

```sh
cd ~/repos/gitlab.com/rwxyou/dotfiles
```

Create a `vim` subdirectory inside of your [`dotfiles`](../) directory.

```sh
mkdir vim
```

Change into the `vim` directory.

```sh
cd vim
```

Create or obtain a Vi/m configuration file and place within directory as `vimrc`.

```sh
curl -LO https://rwx.gg/tools/editors/vi/vimrc
```

Create a `setup` script.

```sh
touch setup
chmod u+x setup
vi setup
```

Have it print `Configuring Vim` for now.

```sh
#!/bin/bash
echo Configuring Vim
```

Run it to test.

```sh
./setup
```
```{.out}
Configuring Vim
```

Add a line to create a [symbolic link](https://duck.com/lite?kae=t&q=symbolic link) from your `$HOME` directory to the local `vimrc` file. Don't forget the full paths in the command. You can also add another line to confirm that the link was created successfully.

```sh
#!/bin/bash
echo Configuring Vim
ln -s "$PWD/vimrc" "$HOME/.vimrc"
ls -la ~/.vimrc
```

Run it to create your link.

```{.out}
Configuring Vim
lrwxrwxrwx 1 sample sample 55 May 27 22:12 /home/rwxrob/.vimrc -> /home/rwxrob/repos/gitlab.com/rwxyou/dotfiles/vim/vimrc
```

That's it. You should just be able to edit any file now to see the differences.

:::co-stop
Note that the first time your `vim` runs it might need to download some things to configure itself. Don't be alarmed.
:::

You might also want to add a `README.md` describing the process of how to run `./setup` and perhaps make your dotfiles directory public so you can share your configuration with others.

If you want to get really fancy, you can actually add some code to detect if the full `vim` is on the system and prompt to have it installed and run `vim` for the first time to that it downloads and sets up the plugins before you run it later for the first time. Here's the full script:

```sh
#!/bin/bash

if [[ -z "$(which vim)" || -z "$(which vimtutor)" ]]; then
  echo Need to install full vim.
  exit 1
fi

echo Configuring Vim
ln -sf "$PWD/vimrc" "$HOME/.vimrc"
ls -l "$HOME/.vimrc"
```

