---
Title: Create Secure Shell Keys
Subtitle: Using [`ssh-keygen`](https://duck.com/lite?kae=t&q=`ssh-keygen`)
Query: true
---

Creating a public-private secure shell key pair is required for many applications and services such as authenticating to GitLab. SSH has many different key types  which are all based on different mathematical algorithms that vary in security and complexity. These days the Elliptical Curve is most frequently recommended. It's fast, small, and "quantum resistant" meaning it is expected to be harder to eventually crack when quantum computers become common. Here are the steps to creating your first key pair from the command line.

:::co-stop
It's possible you might have already set up your SSH keys. You can check quickly by doing a [Git Ping](/tools/git/ping/).
:::

Open a Bash shell terminal to get a command line.

:::co-fyi
If at any time during the following something goes wrong and you want to cancel try hold `Control` and type `c` (which sends an *interrupt* signal) or hold `Control` and press `d` key (which sends an *end-of-file* character).
:::

Enter the following:

```sh
ssh-keygen -t ed25519
```

You should see something like the following:

```{.out}
Generating public/private ed25519 key pair.
Enter file in which to save the key (/home/you/.ssh/id_ed25519/):
```

Press the `Return/Enter` key at each prompt to accept the defaults for everything, which should like this:

```{.out}
Enter passphrase (empty for no passphrase):
Enter same passphrase again:
Your identification has been saved in /home/you/.ssh/id_ed25519.
Your public key has been saved in /home/you/.ssh/id_ed25519.pub.
The key fingerprint is:
SHA256:zel859ZHCUi6Bp83Ao1iHvWUZAK9Qs68ocGOLNjNtjM you[@host](https://twitch.tv/host)
The key's randomart image is:
+--[ED25519 256]--+
|     .o..o.      |
|     . ooo .     |
|  . = . * o .    |
|   o X = * o .   |
|o.oo= * S *   . .|
|oo.o+o   O o   ..|
|.  . .  . = o .o |
|    E      . o. o|
|     o       .. .|
+----[SHA256]-----+
```

Congrats! You now have a public/private key pair.  You can view your **public** key using the `cat` command after changing into your home directory with `cd`.

```sh
cd
cat .ssh/id_ed25519.pub
```

Or, if you are fancy and have `xclip` available you can use that instead to copy it to the clipboard.

```sh
xclip < ~/.ssh/id_ed25519.pub
```

:::co-warning
Be *really* careful not to display your ***private*** key to *anyone*. No one needs to see it but you.
:::

Here's an example of a public key:

```{.out}
ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIOxsQnKQcnJ3zymnCdfdGvYZX898LVkDhz9gvMY5FRFj you@host
```

You can now share this public key with services like [GitLab](/services/gitlab/pubkey/) or other servers like [PicoCTF](https://duck.com/lite?kae=t&q=PicoCTF). Don't forget to copy *all* of it including the `ssh-ed25519` at the start and the `you@host` stuff at the end. 

:::co-btw
You don't need to worry that *this* private key doesn't have a passphrase. It is a convenience for authenticating to services that are not highly critical. When using public-private key authentication for servers that have higher security you would want a rather strong passphrase since it is your last line of security if someone gets your private key. But if you are that concerned about the security of your private key you should probably look at [KeePassXC](https://duck.com/lite?kae=t&q=KeePassXC) for storing it instead of ever writing it to disk.
:::
