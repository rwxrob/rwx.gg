---
Title: What are Containerization and Containers?
Subtitle: Applications on Steroids
Query: true
---

*Containers* are essential applications composed of other applications and systems that are contained, portable, and runnable using a technology such as Docker. Everything from a simple Minecraft server, to an entire operating system can be placed within containers. Containers themselves are *not* [virtual machines](/what/virt/) but the program that runs and manages the containers *does* require virtualization to do it. As Mike Coleman explains, virtual machines are like "houses" and containers are like "apartments in an apartment building".

:::co-warn
Containerization is often confused with [virtualization](/what/virt/). Be sure to know the difference.
:::

:::co-fun
The command to display the running containers is `docker ps` just like `ps` is the command to show the running processes on any UNIX/Linux machine. This similarity between running programs (processes) and running images (containers) is a good way to get your head around what a container is, which is closer to an application than to a machine.
:::

## See Also

* [Containers are Not Virtual Machines](https://www.docker.com/blog/containers-are-not-vms/)
