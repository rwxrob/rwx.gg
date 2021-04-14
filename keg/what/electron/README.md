---
Title: What is the Electron Framework?
Subtitle: Bloated Combination of Chromium and Node Connected by Websockets
Query: true
---

[Electron](https://www.electronjs.org) is the name of an all-web framework for creating desktop applications using nothing but web technology. Even though it has been adapted to allow other code to work with it Electron remains primarily built for [HTML](/lang/html/), [CSS](https://duck.com/lite?kae=t&q=CSS), [JavaScript Front-End frameworks](https://duck.com/lite?q=Front-End frameworks), and [NodeJS](https://duck.com/lite?q=NodeJS). It uses a technology called [Websockets](https://duck.com/lite?q=Websockets) to provide essentially a miniature and compact copy of everything required to create a web application: a NodeJS web server, a [Chromium](https://duck.com/lite?q=Chromium) client browser, and some wrapping to hold it all together.

## Well-Known Applications Built on Electron

Lots of applications that you know and use are built with Electron:

* [Discord](https://discord.gg)
* [Slack](https://slack.com)
* [Spotify](https://spotify.com)
* [Mullvad VPN Client](https://mullvad.net)

Just to name a few.

## Strong Industry Criticism

Although Electron has been widely adopted it has received extremely harsh criticism from well-informed technical architects because of its wasteful management of resources, frequent critical memory leaks, and sometimes disastrous affect on the computer such applications run on. 

Considering that *every* Electron application contains a *full version of Chromium* it's not hard for even a beginning technologist to understand why. It is not unusual for people to effectively have five or six distinct copies of Chromium sucking down resources from their computer at the same time. 

Given that Chromium is regularly the biggest resource hog of any desktop system this problem becomes clear pretty quickly. No one is thinking about what will happen if *everything* is made with Electron, only what they can get away with. And because corporations don't really care about the desktops of anyone, they release their apps this way anyway because they can deliver so much more quickly and just hire a bunch of web developers to make their applications instead of those who can develop using alternative technologies such as Qt, GTK, and others. 

## See Also

* [What is Electron?](https://youtu.be/NpQqxOxMVlo)
* [Electron Sucks?](https://youtu.be/_eJERvFTiJQ)
