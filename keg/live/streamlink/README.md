---
Title: StreamLink for Better Twitch Quality
Query: true
---

Twitch can be really heavy on your CPU. While watching a steam at 1080p 60fps CPU usage can fluctuate between 18-24% usage, but there's a better approach to watching Twitch streams. The only negative side affects is that Twitch.tv *may* not register you as a viewer if you use this method. The solution is to watch Twitch via a publicly available Twitch API called [StreamLink](https://streamlink.github.io/) which is available for Linux, MacOS, and Windows. This software allows *anyone* to watch a live-stream via Linux terminal and Windows PowerShell. By watching like this you remove the website part of the story thus saving a lot of resources. 

## Installation for Debian-Based Distros

Copy and paste the following code into your terminal one line at a time.

```sh
sudo add-apt-repository ppa:nilarimogard/webupd8
sudo apt update
sudo apt install streamlink
```

After you finish those steps proceed to your terminal and type `streamlink <streamurl> <quality>`. If a stream only has one quality option you should pick `worst` but you can test the different settings and see what works for you. The full documentation for options is on the [StreamLink web site](https://streamlink.github.io/cli.html#cmdoption-arg-stream).

## Examples

```
streamlink https://www.twitch.tv/rwxrob best  
streamlink https://www.twitch.tv/rwxrob worst
streamlink https://www.twitch.tv/rwxrob "720p,480p,best"
```

