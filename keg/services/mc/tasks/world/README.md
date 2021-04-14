---
Title: Upload a Minecraft World
Query: true
---

First you have to have a `world` to upload. You could just upload a world that you have created locally on your computer. Or, you could copy the world over from another server host computer. Maybe you just started another Minecraft server and just need to copy over the `world` from the previous one.


## Upload from Local Computer

*Need to add steps.*

## Transfer from Another Server Host

*Need to add steps.*

## Copy from Previous

First of all stop your server.

Change into the server directory if you have not already. The name is what you named it. It might be named `mc`, maybe `paper`, or perhaps you just used the home directory.

```sh
cd paper
```

Now go ahead and make sure you are in the right place by listing the files.

```sh
ls
```

You *should* see a `world` directory unless you changed the name in `server.properties`.

Then move your `world` directory out of the way without deleting it by simply renaming it.

:::co-pwz
Make sure the name does already exist or you will either destroy that file or, if it is a directory, you'll move `world` into the directory instead of renaming.
:::

```sh
mv world oldworld
```

Now list everything to see if it worked.

```sh
ls
```

Now is the time to decide if you *also* want to copy over the `world_nether` and `world_the_end` but usually it is not needed.

Next we need to *copy* the world we are transferring into this server. 

:::co-pwz
Note that we are *copying* and *not* moving so that we still have a backup.
:::

```sh
cp -r ../mc/world .
```

The `-r` means to *recursively* copy the directory and everything that it contains. Without it you won't get the directories contents. The `../` means parent directory. The trailing `.` means the current directory that I'm in right here, right now.

Now list everything to see if it worked again.

```sh
ls
```

You should see a new `world` directory.

Now you should simple be able to restart your Minecraft server to connect and check it out.

