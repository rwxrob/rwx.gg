---
Title: What is Docker Container Engine?
Subtitle: First and Most Popular Container Application
Query: true
---

*Docker* is the first application to popularize the use of [containerization](../) (which actually existed long before). Docker added <https://hub.docker.com> a place where containers of all kinds are uploaded and available to run with a single command from the terminal, including a full Ubuntu Linux server with an interactive Bash shell. The benefit of containers is that they can be run on *anything* that supports the container engine. This provides another way to get Linux easily working on a Windows or Mac, for example.

:::co-warn
It is critically important that you remember that unless you specifically configure Docker to use some disk space (where your files and directories are) that it will *not* remember anything after you stop the container, which might be as simple as just exiting the interactive container shell. This is great to try things out quickly, but not enough for long term usage and development.
:::

```sh
sudo apt install docker.io
sudo docker run hello-world
```

Then you can add yourself to the `docker` group.

```sh
sudo usermod -aG docker $USER
```

After that you won't need to do `sudo` any longer.

Docker can be used to run any number of applications including a full Ubuntu instance with a Bash shell.

```sh
docker run -it -d --name mylinux ubuntu bash
```

Then after you disconnect from it you can go back to it.

```sh
docker exec -it mylinux bash
```

You can practice setting up services and connecting to them such as installing the `openssh-server` and activating it.

```sh
apt update
apt install openssh-server
service sshd start
```

Maybe practice adding a user.

```sh
sudo adduser $USER 
sudo usermod -aG sudo $USER
```

Then [setting up keys](/tools/ssh/tasks/keygen/) so you can treat it like a remote [cloud](/what/cloud/) server. But first you'll need `vim` installed.

```sh
su - <you> 
sudo apt install vim
mkdir ~/.ssh
chmod 700 ~/.ssh
touch ~/.ssh/authorized_keys
vi ~/.ssh/authorized_keys
```

You can paste your *public* key from another pane into the `~/.ssh/authorized_keys` file. Then from another pane on your host system attempt so `ssh` into it after you look up the IP with the `docker inspect` command from the host.

```sh
docker inspect mylinux
ssh <you>@<localip>
```

When you are done playing around with the container, make sure when you are done to `stop` and optionally `rm` it.

```sh
docker ps
docker stop mylinux
docker rm mylinux
```

:::co-rant
The hub.docker.com web site was made by brain-dead web designers who have no interest in preserving the Internet standards that gave us the World Wide Web in the first place. They are part of a growing minority of people think *all* of the Web *must* support JavaScript. Accessing the site without JavaScript, say from Lynx, produces a black whole. Thankfully the whole Docker organization and team is not as stupid as their web designers.
:::

## See Also

* [Docker.io](https://docker.io)
* [Is Linux Now Moby?](https://www.mirantis.com/blog/ok-i-give-up-is-docker-now-moby-and-what-is-linuxkit/)
* [Willy Wonka of Containers](https://www.youtube.com/watch?v=GsLZz8cZCzc)
