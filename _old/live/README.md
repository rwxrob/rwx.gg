---
Title: Twitch Live Streaming
Subtitle: Madness and More 
Query: true
---

A big part of the community is joining [\@rwxrob](https://twitch.tv/rwxrob) and other [streamers](./streamers/) on Twitch. There you can just get answers, randomly banter about stuff, or perhaps join a live "studio" audience for a proper video recording.

:::co-faq

## I Can't Watch Your Stream at 1080?

First of all, just understand that *nobody* controls the transcoding of your Twitch but Twitch. It is entirely up to them to provide transcoding to affiliates and others in way that is not *entirely* random. The best you can do is try to fix with [StreamLink](./streamlink/) or hope that streamers are able to get the transcoding.

While it is true that restarting a stream over and over again until you get transcoding is *one* way to help this is less and less possible now that so many more people are streaming at regular times during the day and night. To help we've tried to stream 24 hours a day once we get a good stream and transcoding rather than resetting and risking losing the good transcoding. Unfortunately that means having a system setup and streaming all day, which isn't always practical.

Here's another suggestion from [\@tafetanes](https://twitch.tv/tafetanes) that might work for some. You'll just need to know a little about scripting:

```sh
ffmpeg -i $(youtube-dl -f 1080p $rwxstream -g) ./livestream.ts & ; mplayer ./livestream.ts
```

:::


