---
Title: What are Domain Names?
Subtitle: Easier to Remember than IP Numbers
Query: true
---

*Domain names* are what we use for websites, email, game servers, and such. They are friendlier than [IP numbers](https://duck.com/lite?kae=t&q=IP numbers). Getting one is pretty easy these days --- especially now that the [Netlify](/services/netlify/) and [Namecheap](https://namecheap.com) [services](/services/) exist.

:::co-rant
Don't even *think* about getting a GoDaddy domain. The company is infamous using women's scantily clad breasts to peddle domains and is notorious for holding domains hostage, pushing unwanted stuff on you, and having the worst customer service of any provider on the planet. It just needs to die the death it deserves.It is everything you might expect from the digital equivalent of a slimy real estate agent.
:::

First, make sure you have a *secure* email provider such as [ProtonMail](https://duck.com/lite?kae=t&q=ProtonMail). 

:::co-warning
Do *not* use Gmail. Google owns your email, not you and Gmail is regularly hacked, not something you want for the email used to manage your personal or professional domain names.
:::

Create an account on [namecheap.com](https://namecheap.com) (using your secure email address).

Then find a domain you like and buy it.

:::co-fyi
There are over 500 top-level domain suffixes these days. It's hard to pick. The `dev`, `tech`, `gg`, `live`, and `io` domains seem to be coolest latest, but find your own that matches your goals. Keep in mind that your domain will be likely used *both* for your web site *and* for email so avoid domains like `site`, etc. The `xyz` domain says, "I'm cheap." The old `com`, `biz`, and `net` domains mean you're old.
:::

You'll have to wait a while for the Internet to see your new fancy domain. Be patient. When it *does* see it you'll know because a placeholder from Namecheap will show up. You can also use the following command to check your domain.

```sh
dig <yourdomain>
```

Or the older version.

```sh
nslookup <yourdomain>
```

After your domain is ready you will want to point it at your site and possibly email.
