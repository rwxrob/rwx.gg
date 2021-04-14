---
Title: Authorize a Remote Public Key
Subtitle: Let Someone Into Your Remote SSH Server
Query: true
---

Often you want to let someone (or something) connect safely to your account on a remote computer. The most common way to allow this is to add a public key to your account on the computer to which you want to allow remote access. This is assuming you have already [enabled secure shell server](/services/sshd/) so there is something to connect to.

You'll need to do all of this locally on the computer to which you want to provide remote access. This can get tricky if that computer is in another state, say, at a Digital Ocean location in New York. Thankfully cloud providers usually have a web-interface to the "local" console on that machine for such things.

:::co-fyi
Actually in Digital Ocean's case you don't even need to do this. You can just add the public key through their web interface, but it's good to know for your own servers.
:::

Get on a local console and login to the account to which you want to allow people to connect, this can even be `root` (provided you made sure to disable password login because of how insecure that would be otherwise). Now check that you don't already have an `.ssh` directory in that account's home directory.


```sh
cd
ls .ssh
```

```{.out}
ls: cannot access '.ssh': No such file or directory
```

If you see that you don't have one and need to make it. The `chmod 700` forces the directory to only be browsable, readable, and writable by you. Without this secure shell will throw and error and block incoming connections.

```sh
mkdir .ssh
chmod 700 .ssh
```

Then you need to create the `.ssh/authorized_keys` file. 

```sh
touch .ssh/authorized_keys
chmod 600 .ssh/authorized_keys
```

Now just edit the file and add the public key.

```sh
vi .ssh/authorized_keys
```

Next you should test your ssh connection from the other computer. Don't forget to provide the name of the user on the destination system if it is different than that from the connecting system, for example, `ssh rwxrob@192.1.1.21` instead of just `ssh 192.1.1.21`.
