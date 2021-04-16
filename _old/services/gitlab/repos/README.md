---
Title: Set Up Local Home GitLab Project Repos
Subtitle: Sustainably Manage Repos from Multiple Git Hosting Services
Query: true
---

If you haven't already done so, complete the following:

1. Create a GitLab account. 
1. Create a GitLab project repo.
1. [Create local SSH keys](/tools/ssh/tasks/keygen/) for this computer.
1. [Add your public key to GitLab.](/services/gitlab/pubkey/)

:::co-stop
GitLab allows you to create project repos *without* the web interface just by pushing them, but these instructions assume you are on a computer that does not have any repos cloned although you already have one on GitLab already. If you don't have any repo at all on GitLab go ahead and create that using the graphic web interface. These instructions are valid for setting up a new computer or adding another workstation after having already done work from another one.
:::

Change into your home directory and make an empty `repos` directory.

```sh
cd
mkdir repos
```

Because you might have repos from multiple Git hosting services (GitLab *and* GitHub, for example) let's make a directory to match the name of the service.

Change into your `repos` directory and make a new `gitlab.com` directory and then change your current directory into it.

```sh
cd repos
mkdir gitlab.com
cd gitlab.com
```

:::co-fyi
Remember you can tap the `Tab` key instead of typing out all of `gitlab.com`
:::

Now we need to make another directory that matches the user or group name on GitLab that we are going to clone from and then change into that.

```sh
mkdir rwxyou
cd rwxyou
```

Now we are ready to clone the project repo using Git.

```sh
git clone git@gitlab.com:rwxyou/notes.git
```

```{.out}
Cloning into 'notes'...
remote: Enumerating objects: 3, done.
remote: Counting objects: 100% (3/3), done.
remote: Total 3 (delta 0), reused 0 (delta 0), pack-reused 0
Receiving objects: 100% (3/3), done.
```

Now we have a *local* project that we can work on. Make your changes and don't forget to save them with a Git commit and push.
