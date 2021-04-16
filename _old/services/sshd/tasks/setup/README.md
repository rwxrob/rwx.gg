---
Title: Set Up Secure Shell Server Daemon
Subtitle: Enabling Remote Connections to Your Server
Query: true
---

You can't do much of anything on a remote server until you can connect to it. The standard way to make such remote connections is with `ssh` but first you have to setup a secure shell server called `sshd` to which you will connect.

:::co-stop
The `sudo` command is included in the following but if you are already logged in as `root` you obviously do not need to do that.
:::

:::co-fyi
The `d` in `sshd` is for [daemon](/what/daemon/).
:::

You might already have `sshd` installed. Here are a few ways to check. If you do you can skip installation.

```sh
which sshd
ls /usr/sbin/ssh*
```

```{.out}
/usr/sbin/sshd
/usr/sbin/sshd
```

## Install Secure Shell Server

For Linux you want the `openssh-server` package (or equivalent). As usual, you might need to update the list of packages with `sudo apt update`.

```sh
sudo apt update
sudo apt install openssh-server
```

## Disable Password Authentication

You should never allow remote password login for any reason. Sometimes this is the default but always check. Open up the configuration file.

:::co-mind
Don't forget the `d` at the end of `sshd_config` because there is also a `ssh_config` file.
:::

```sh
vi /etc/ssh/sshd_config
```

Now search for the following line to jump to it (`/Pass`)

```sshd_config
#PasswordAuthentication yes
```

Uncomment that line and change it to `no`. We don't ever want to allow password login with `ssh` because it is far too easy to defeat. We'll use SSH keys instead.

```sshd_config
PasswordAuthentication no
```

## Add a Public Key to the `root` Account

You might already have a `~/.ssh/id_25519.pub` key to use. If not [generate an SSH key pair](/tools/ssh/tasks/keygen/) on the system from which you will do the connecting, your laptop or workstation.

Then put that *public* key on a USB stick or even in a GitLab repo or a web page so you can get it onto your system that now has `openssh-server` installed.

Now you need to [authorize](/services/sshd/tasks/authpub/) the remote public key on your new server from the `root` account (or another one if you want to create it).

:::co-faq
**Why Allow Root Login?**

Allowing `root` to login *after* password login has been disabled is no where near as risky otherwise. Many cloud providers, including Digital Ocean are fine with having `root` login provided a public key has been provided and is the only means of connecting.

You may still need another account to run certain services with minimal permissions. If that is the case, *do not give that user sudo access*, which defeats the purpose of lesser permissions in the first place.
:::

## Enable the Secure Shell Service

First check to see that your server has not already been enabled.

```sh
sudo systemctl status sshd
```

Try one of the following to enable and activate your server.

```sh
sudo systemctl enable sshd
```

:::co-fyi
This is sometimes `ssh` instead of `sshd` for some reason, even on other Ubuntu derivatives.
:::

## Restart the Secure Shell Service

If your server was already running you need to restart it. Use one of these. 

```sh
sudo systemctl restart sshd
```

