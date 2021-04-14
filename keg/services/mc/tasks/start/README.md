---
Title: Start Minecraft Server
Query: true
---

First you to have set up a Minecraft server.

Connect to the server remotely with SSH.

Change into the directory containing the Minecraft server files.

```sh
cd paper
```

List the server files just to be sure.

```sh
ls
```

The Minecraft server is a program that must keep running in order to handle new connections. Normally when you disconnect from a remote host computer all your programs stop running. It is the same as if you just hit the power on your computer. Although it doesn't *actually* does not turn off the host computer, it *does* turn off and close all your running programs including your Minecraft server. So we need a way to keep the program running even after we disconnect.

Screen keeps running even after we disconnect. That includes any programs we started while using `screen`, which is called a "screen session." 

Go ahead and start a `screen` session.

```sh
screen
```

You'll see the Screen startup text including `Press Space or Return to End` which means just tap `Return` to continue to a shell that is running *under* screen. You can test this by pressing `Control-a` and then `w`. That will show you the windows you have. For this we only need one window.

:::co-btw
TMUX is another terminal multiplexor like screen but is not always available on host computers without installing it.
:::

Now you are ready to start the Minecraft server. Locate your `.jar` file that is your server.

```sh
java -jar paper.jar
```

:::co-btw
You will likely want to use a more complicated `java` start command that sets the memory later.
:::
