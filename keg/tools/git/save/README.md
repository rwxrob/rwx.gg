---
Title: Saving a Git Project Repo
Subtitle: Understand and Simplify Common Git Commands
Query: title
---

Sometimes you just want to save your work and use Git as a sort-of backup system. This is perfectly fine --- especially if you are just starting out using Git in the first place. Still, before you use something like the [sample script below](#save-script) make sure you at least use the actual `git` commands a few times to initially learn them --- even if you don't fully understand all the details.

:::co-rage
**Don't feel bad or small for using Git as a backup system.**

Git is a tool and how you use the tool depends entirely on your use case, not some pedantic commandments from the Git gods. It is far better to frequently commit and push (annoying anyone following your GitLab feed) than to *fail* to commit and really messing things up because you made a common mistake that blows away a file or three. Eventually you will get good at making estimates of the best time to make a push so that the comments do make sense instead of `WIP` every time, but don't beat yourself up now while you are learning. Besides, there is always `squash` and `rebase` when drastic measures are needed on an actual initial release.
:::

## Individual Git Commands

Here are the steps and commands to do an effective *save* of your Git repo. But first, you need to have [your home repos set up](/services/gitlab/repos). This also assumes you have already created and/or cloned the repo and are now in it.

First `cd` into your project repo if you haven't already. It might look something like this command.

```sh
cd ~/repos/rwxrob/gitlab.com/somerepo
```

Okay, now let's check that we have some changes to actually save.

```sh
git status
```

It should be obvious if you don't have any changes. If not then make some changes at least so you can complete this exercise.

Now pull down any potential changes you might have made previously either from GitLab's editor or perhaps another workstation.

:::co-stop
Even if you are sure you haven't make any changes anywhere else get into the habit of doing this or one day you will be really surprised with an ugly merge conflict to fix.
:::

```sh
git pull
```

Now let's add all the changes in the current directory. Make sure you are at the top level of your repo directory and not in some subdirectory.

```sh
git add -A .
```

Try another `git status` and notice how things have changed.

Now commit and automatically stage (`-a`) with a new meaningful but short message (`-m`).

```sh
git commit -a -m "this is a meaningful message"
```

:::co-stop
Note that this might fail because you have not [configured your `git` user information](../config). Stop and do that now if so.
:::

That's it for everything on your local workstation. In fact, you could go ahead and do other work and commits before pushing them all up to the mother Git hosting ship. But the point is to stay safe by having that mother ship in the cloud updated so that even if your workstation melted your work would be fine.

The command for this is `git push` but if this is your first time, or you just want to be safe include exactly what you are pushing to.

```sh
git push -u origin master
```

You should have an update GitLab project repo at this point. Go check it out in the graphic web interface to be sure.

:::co-tip
That push command will also automatically create a new private project repo on GitLab (but not GitHub) if this was not cloned from there already and is the first push. It is one of the many things that makes [GitLab so superior](/services/gitlab/isbest/).
:::

## Save Script

Here is an example of a Bash script that might be like one you would create to incorporate your most common tasks. Consider copying it into a script named `save`

```bash
#!/bin/bash

# WARNING: This script assumes you are using GitLab for everything and
# only work on the master branch.

# The init function creates and initializes a git repo and then pushes it
# as a private project to GitLab (without needing to first create it on
# GitLab)

init () {
  declare repo y
  read -p "Not a git repo. Create? " y
  if [[ $y =~ ^[yY] ]]; then
      touch README.md
      read -p "GitLab path (<id>/<proj>): " repo
      if [[ -z "$repo" ]]; then
        return init
      fi
      git init
      git remote add origin "git@gitlab.com:$repo.git"
      git add .
      git commit -m initial
      git push -u origin master
  fi
  return 0
}

check-for-git () {
  declare user=$(git config user.name);
  if [[ -z "$user" ]]; then
    echo "Git doesn't look configured yet. Try:"
    echo "  git config -g user.name <name>"
    echo "  git config -g user.email <email>"
    return 1
  fi
  return 0
}

in-repo ()  {
 git rev-parse --count HEAD > /dev/null 2>&1
 return $?
}

has-local-changes () {
  test -n "$(git status --porcelain)"
  return $?
}

gitsave () {
  local comment=wip
  [ ! -z "$*" ] && comment="$*"
  git pull
  git add -A .
  git commit -a -m "$comment"
  git push -u origin master
}

main () { 
  declare message="$*"
  message=${message:-initial}
  check-for-git || exit 1
  if ! in-repo; then
    init
    return 0
  fi;
  if ! has-local-changes; then
    echo Already at the latest.;
    return 0;
  fi;
  gitsave "$message"
  return $?
}

main $*
```
