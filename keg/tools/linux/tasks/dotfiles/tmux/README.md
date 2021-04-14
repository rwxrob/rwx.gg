---
Title: Set Up TMUX Dotfiles Configuration Subdirectory
Subtitle: Single File, 100% Portable, No Blocking Dependencies
Query: true
---

Make sure you have setup your [dotfiles](../) repo.

Change into your `dotfiles` repo.

```sh
cd ~/repos/gitlab.com/rwxyou/dotfiles
```

Create a `tmux` directory.

```sh
mkdir tmux
```

Change into the `tmux` directory.

```sh
cd tmux
```

Create or obtain a [TMUX configuration file](/tools/tmux/) and place it within directory as `tmux.conf`.

```sh
curl -LO https://rwx.gg/tools/tmux/tmux.conf
```

Create a `setup` script that prints `Configuring TMUX`.

```sh
#!/bin/bash
echo Configuring TMUX
```

Test that your script works.

```sh
./setup
```
```{.out}
Configuring TMUX
```
Add a line to create a [symbolic link](https://duck.com/lite?kae=t&q=symbolic link) from your `$HOME` directory to the local `tmux.conf` file. Don't forget the full paths in the command. You can also add another line to confirm that the link was created successfully.

```sh
#!/bin/bash
echo Configuring TMUX
ln -s "$PWD/tmux.conf" "$HOME/.tmux.conf"
ls -la ~/.tmux.conf
```

Run it to create your link.

```{.out}
Configuring TMUX
lrwxrwxrwx 1 sample sample 55 May 27 22:12 /home/rwxrob/.tmux.conf -> /home/rwxrob/repos/gitlab.com/rwxyou/dotfiles/tmux/tmux.conf
```

That's it. You should just be able to edit any file now to see the differences.

You might also want to add a `README.md` describing the process of how to run `./setup` and perhaps make your dotfiles directory public so you can share your configuration with others.

If you want to get really fancy, you can actually add some code to detect if `tmux` has been installed on the system and prompt them to install it first. Here's the full script:

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


Once you know it runs. Add a line to create a [symbolic link](https://duck.com/lite?kae=t&q=symbolic link) to `~/.tmux.conf`.


