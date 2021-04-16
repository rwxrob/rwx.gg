---
Title: Git Hosting Service Ping
Subtitle: Check if Your SSH Keys are Working
---

A *Git ping* is just a quick way to tell if your [SSH keys](/tools/ssh/tasks/keygen/) are working with a Git hosting provider like GitLab.

:::co-fyi
There is no *actual* `git ping` command, nor is the ICMP `ping` protocol or command involved at all.
:::

First you will have to have completed the following:

1. Created and account with the Git hosting provider
1. Created your [secure shell keys](/tools/ssh/tasks/keygen/)
1. Added your *public* key to your account with provider

If you have *not* done any of those correctly then the following will fail.

```sh
ssh git@gitlab.com
```

If prompted to trust and continue go ahead and answer yes (if you don't think someone is hacking you).

```{.out}
The authenticity of host 'gitlab.com (172.65.251.78)' can't be established.
ECDSA key fingerprint is SHA256:HbW3g8zUjNSksFbqTiUWPWg2Bq1x8xdGUrliXFzSnUw.
Are you sure you want to continue connecting (yes/no)? yes
Warning: Permanently added 'gitlab.com,172.65.251.78' (ECDSA) to the list of known hosts.
```

You should then see something like "Welcome! Now go away" then you have successfully added your SSH public key to the hosting provider.

```{.out}
PTY allocation request failed on channel 0
Welcome to GitLab,[@rwxyou](https://twitch.tv/rwxyou)!
Connection to gitlab.com closed.
```

:::co-stop
Remember that you still need to configure `git` in addition to adding the public key to the hosting provider.
:::

If you failed the configuration test you'll see something like this.

```{.out}
Warning: Permanently added 'gitlab.com,2606:4700:90:0:f22e:fbec:5bed:a9b9' (ECDSA) to the list of known hosts.
git@gitlab.com: Permission denied (publickey).
```
