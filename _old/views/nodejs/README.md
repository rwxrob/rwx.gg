# View on Node.js

Our view on Node.js can be summarized best by its creator [Ryan Dahl](/#ryan-dahl) and one of the biggest Node.js "defectors" in history, [TJ Holowaychuk](/#tj-holowaychuk):

## Ryan Dahl's Views on Node.js

> 💬 Yeah, I think that’s a really interesting question. Now, it’s been several years, and I haven’t worked on Node myself since like, 2012, or 2013. And Node, of course, is a big project at this point. So, yeah, I think… when it first came out, I went around and gave a bunch of talks, trying to convince people that they should. That maybe we were doing I/O wrong and that maybe if we did everything in a non-blocking way, that we would solve a lot of the difficulties with programming. Like, perhaps we could forget about threads entirely and only use process abstractions and serialized communications. But within a single process, we could handle many, many requests by being completely asynchronous. I believe strongly in this idea at the time, but over the past couple of years, I think that’s probably not the end-all and be-all idea for programming. In particular, when Go came out.
>
> Well, I think Go came out a long time ago, but when I first started hearing about Go, which was around 2012. They had a very nice runtime that had proper green threads and easy to use abstractions around that. That I think makes blocking I/O – again, blocking I/O in quotes, because it’s all in green threads at the interface of… between Go and the operating system, I think it is all non-blocking I/O.
> 
> But the interface that they present to the user is blocking, and I think that that’s a nicer programming model. And you can think through what you’re doing in many situations more easily if it’s blocking. You know, if you have a bunch of following actions, it’s nice to be able to say: do thing A, wait for a response, maybe error out. Do thing B, wait for a response, error out. And in Node, that’s harder, because you have to jump into another function call.
> 
> ...
>
> Yeah, I think it’s ... for a particular class of application, which is like, if you’re building a server, I can’t imagine using anything other than Go. That said, I think Node’s non-blocking paradigm worked out well for JavaScript, where you don’t have threads. And I think that a lot of the problems with kind of the call-back soup problem, where you have to jump into many anonymous functions to complete what you’re doing has been alleviated these days, with the async keyword, the async feature that’s in Javascript now.
> 
> So, kind of the newer versions of Javascript has made this easier. That said, I think Node is not the best system to build a massive server web. I would use Go for that. And honestly, that’s the reason why I left Node. It was the realization that: oh, actually, this is not the best server-side system ever.

[*Mapping the Journey* Interview on August 31, 2018](https://mappingthejourney.com/single-post/2017/08/31/episode-8-interview-with-ryan-dahl-creator-of-nodejs/)

## TJ Holowaychuk's Views on Node.js

> As far as platforms are concerned, Go is in my opinion far more robust than Node.js, with a rich stdlib and thriving community, so much less is needed on that front. It is roughly the same age as Node, so it doesn’t really need much! [Quora, June 23, 2017](https://www.quora.com/Has-TJ-Holowaychuk-been-as-prolific-in-the-Golang-community-as-he-was-in-the-Node-js-community)

TJ's now [Farewell Node.js](https://medium.com/@tjholowaychuk/farewell-node-js-4ba9e7f3e52b) has reached icon status in pop tech culture with his most important message being "don't get stuck in your own bubble," some of the most relevant tech wisdom in history by one of the more humble, respected technologists in history.

> 😲 Incidently, some point to this change in interest and focus and being parallel with [TJ Holowaychuk](/#tj-holowaychuk) as evidence that Ryan is perhaps one part of the possible myth that TJ is a [hive mind](/hive-mind).

