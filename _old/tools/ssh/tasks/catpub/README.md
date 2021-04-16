---
Title: Print Your Public Secure Shell (SSH) Key with `cat`
Subtitle: Careful, *Not* Your Private Key
---

Here's how to quickly display the content of your ED25519 Secure Shell
public key.


```sh
cat ~/.ssh/id_ed25519.pub
```

You should see something like the following:

```{.out}
ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIOxsQnKQcnJ3zymnCdfdGvYZX898LVkDhz9gvMY5FRFj rwxyou@home
```

If that doesn't work then you either don't have a key at all or you
don't have the right *type* of key. Either way you should go [generate
one](../keygen/). Otherwise, proceed.

When you provide your public key *all* of it is required ***including
the `ssh-...` at the beginning and the `rwxyou@home` (or whatever) at
the end***. Frequently people only send the middle part.



