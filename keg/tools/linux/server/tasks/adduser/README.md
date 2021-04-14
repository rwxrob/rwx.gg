---
Title: Add a New User Account on Linux
Query: true
---

Login as `root` into the system. You can do this remotely with `ssh`
sometimes or others you might have to do `sudo su -` instead and type
your password.

```
ssh root@<serverip>
```

Then run the `adduser` command.

```
adduser <username>
```

For example, to create an account named `mc` you would do `adduser mc`.

```{.out}
Adding user `mc' ...
Adding new group `mc' (1002) ...
Adding new user `mc' (1002) with group `mc' ...
Creating home directory `/home/mc' ...
Copying files from `/etc/skel' ...
Enter new UNIX password:
```

Type in the new UNIX password. Note that you will *not* see anything on
the screen for security reasons --- not even stars --- but it *is*
taking in what you are typing. Remember this is just the password for
the *specific* (`mc`) user account *not* `root`.

```{.out}
Enter new UNIX password:
Retype new UNIX password:
passwd: password updated successfully
Changing the user information for mc
Enter the new value, or press ENTER for the default
        Full Name []:
```

You don't really need any of the prompted information but you can
provide it if you want it. Otherwise, just tap enter for everything to
accept blanks.

```{.out}
Enter the new value, or press ENTER for the default
        Full Name []:
        Room Number []:
        Work Phone []:
        Home Phone []:
        Other []:
Is the information correct? [Y/n] y
```

Okay. Now you are ready to test it out. 

First we will just change to the account without even logging in.
Replace `<username>` with the user name you created above.

```
su - <username>
```

So in this example, `su - mc` and should see the prompt change. (This is
not something you type in. It is what you should see after you type in
what's above.)

```{.out}
mc@live:~$
```

That means everything worked.

Now while we are here let's setup our SSH public keys which we probably
already have setup in the `root` account.

First check to see if the `.ssh` directory exists.

```
ls .ssh
ls -ad
```

If it says `no such file or directory` then we need to create it. There
are a couple of ways to create it, but let's just use `mkdir` and
`chmod`.

```
mkdir .ssh
chmod 700 .ssh
ls -ld .ssh
```

You should see the permissions are locked down.

```{.out}
drwx------ 2 mc mc 4096 Jul 23 18:43 .ssh
```

We're not done yet. Now we need to add the public key from our
workstation system into a file called `.ssh/authorized_keys`. First
let's use `touch` to create the file.

```
touch .ssh/authorized_keys
ls -l .ssh/authorized_keys
```

Don't forget to lock down the permissions on this file as well.

```
chmod 600 .ssh/authorized_keys
```

Now we can paste your ***public*** SSH key into the file. To be safe,
just cut it from your terminal that is not logged into the remote
server. Open a new terminal if needed.

On your *local* system (*not* the server) `cat` your public key. Any of
them will do, but you should have an `id_ed25519.pub` if you are current
and using the best.

```
cd
cat .ssh/id_*.pub
```

```{.out}
ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIOxsQnKQcnJ3zymnCdfdGvYZX898LVkDhz9gvMY5FRFj rwxrob@live
```

Copy that however you want. Most will just select it and later middle or
right-click to paste using UNIX cut and paste style.

***YOU NEED THE WHOLE THING!*** Yes, the `ssh-` part. Yes, the `...@...`
part. [I really am so tired of answering that
question.]

Now it's time to write the public key into the `.ssh/authorized_keys`
file on the *server*.

Open the file for editing.

```
cd
ls -a
vi .ssh/authorized_keys
```

::: Remember
Any file or directory with a dot in front is hidden by
default from the `ls` command.
:::

Paste it in. If you have stuff already just add it to the end. Save and
exit your editor. 

::: Smile
Aren't you glad you learned `vi` instead of `emacs` or VS Code first?
:::

Now we need to test the remote connection.

Open a *local* terminal. (You might still have one open from cutting the
SSH public key.)

Test the connection with the `ssh` command and the server IP.

```
ssh mc@<serverip>
```

Accept the man in the middle warning and boom, you should be on the
server ready to work.
