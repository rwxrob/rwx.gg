---
Title: Install a Minecraft Server Plugin
Query: true
---

First you need to have set up a Minecraft server.

Then you need to *find* the plugin `.jar` file on the Web that you would download (but don't download it). Make *sure* the version is compatible with the server you are running otherwise it will just spew errors even though your server will still start up.

:::co-pwz
Be *really* careful not to accidentally get a virus or malware from a bad site during this process. Also the Bukkit site is outdated and *rarely* the place to get stuff these days.
:::

So copy the web link to the downloadable plugin file (usually right click).

Connect to server host computer.

If you server is running you need to stop it. 

Now find the `plugins` directory and change into it.

```sh
cd plugins
```

Now you can use the `curl` command to download the file from the link location.

```sh
curl -LO https://ci.athion.net/job/FastAsyncWorldEdit-1.15/lastSuccessfulBuild/artifact/worldedit-bukkit/build/libs/FastAsyncWorldEdit-1.15-256.jar 
```

Now check what you got.

```sh
ls
```

```{.out}
FastAsyncWorldEdit-1.15-256.jar 
```

Let's see what type of file it is.

```sh
file *.jar
```

```{.out}
FastAsyncWorldEdit-1.15-256.jar: Zip archive data, at least v1.0 to extract
```

:::co-pwz

If it says `HTML document` then you didn't actually get the plugin file and will have to look harder on the site for the `.jar` file. If you cannot manage to find the `.jar` file you can always just download it and upload it with `scp` but it takes a few more steps than just using `curl`.

```{.out}
...: HTML document, ASCII text, with very long lines
```

:::
