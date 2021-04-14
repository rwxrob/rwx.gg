---
Title: Set Up PaperMC Server on Linux
Query: true
---

Here's how to set up a [PaperMC Minecraft server](https://duck.com/lite?kae=t&q=PaperMC Minecraft server) in a simple way to get you started. 

First, you will need to [set up a Linux server](/tools/linux/server/tasks/setup/) if you haven't already.

Then you should first [create an
account](/tools/linux/server/tasks/adduser/) specifically to run your
Minecraft server. Let's call the account `mc`. If you are doing this as
another user or at home you can skip past it.

Login to server host.

```sh
ssh mc@myhost
```

Change into your home directory.

```sh
cd
```

Create a `paper` directory to contain server.

```sh
mkdir paper
```

Change into the `paper` directory.

```sh
cd paper
```

Find the paper downloadable Java `jar` file on the Web.

Copy the latest link address by selecting it. 

:::co-pwz
You *do not* need to do `Control-C` to copy on Linux! Simply select the text by highlighting it. Make sure you get the latest download link URL. *Don't just use the following. It might be out of date.*
:::

Download the server jar file and name it `paper.jar` by using the command and pasting your URL. Make sure to follow any server redirects.

```sh
curl -L -o paper.jar https://papermc.io/api/v1/paper/1.15.2/325/download
```

Check the human-readable size and type of the file after downloading it.

```sh
ls -lh paper.jar
file paper.jar
```

Run the server `jar` file to create the `eula.txt` file.

```sh
java -jar paper.jar
```

List the files. Find `eula.txt` and review the content (without editing just yet).

```sh
ls -lh eula.txt
cat eula.txt
```

Agree to the EULA by editing the `eula.txt` file and changing `false` to `true`.

:::co-tip
Although you probably have not yet learned about Perl and [regular expressions](https://duck.com/lite?kae=t&q=regular expressions) this is a perfect example of something that does not even need to be opened up with an editor. It is a safe enough change since the file is simple and the text string to be changed is unique. You could edit the `eula.txt` file *in place* like this:

```sh
perl -pi -e 's/false/true/' eula.txt
```

Perl is the command line tool. The `-p` means to "print" it. The `-i` means to "print" it back to the file "in line" instead of just to the screen. The `-e` means to just do the command `s/false/true/`, which means to substitute the first thing with the second thing.

Be *really* careful if you try this out on other stuff because it is *very* powerful. You will eventually learn to be able to change things in 100s or even 1000s of files at once.
:::

vi eula.txt



* Update Server Properties
    * Set `server-port`
    * Set `difficulty`
    * Set `gamemode`
    * Set `force-gamemode` 
    * Set `motd`
    * Set other settings to taste.
* Run Server Jar with Java
* Test Client Connection to Server
* Op yourself.

* Restart your server under screen so it does not go down.

:::co-pwz
You'll probably want to hone and improve things over time. This is just for absolute beginners who barely know how to use the terminal command line. For example, no discussion of Java memory constraints is included.
:::

## See Also

* <https://papermc.io>
