---
Title: Linux Isn't Dead. GNU Just Killed Itself.
Subtitle: And GNU Deserves It
Query: true
---

The [GPLv3](https://duck.com/lite?kae=t&q=GPLv3) license placed on all [GNU software](https://duck.com/lite?kae=t&q=GNU software) --- that was originally intended to cover [the Linux kernel](https://duck.com/lite?kae=t&q=the Linux kernel) as all previous GPL licenses had done until [Linus Torvaldz vetoed it](https://youtu.be/PaKIZ7gJlRU) --- has destroyed the reputation of the GNU project and indirectly a lot of interest in Linux as well. Software projects and organizations are fleeing GNU as fast as they can. But Linux will remain alive as the short-sighted, over-reaching GNU project and unethical FSF organization slowly die the death they deserve.

:::co-mad
[Stop calling it GNU/Linux. It's disrespectful.](/what/gnulinux/)
:::

## Tivoization Ignores Hardware Constraints

The [Tivoization clause](https://duck.com/lite?kae=t&q=Tivoization clause) is poorly-conceived, overreaching, and *limits* freedoms rather than promoting them by going completely against the original intent of the GPLv2 license that was mostly about giving back any changes made to any *software*. GPLv3 forces *hardware* makers to provide some way to update the software without any consideration or concern for what a specific user needs related to that hardware. 

### ARM with Linux is Getting Smaller and Smaller

The FSF's draconian decision to force upgrades of GNU software ignores the reality that devices running GNU software are getting smaller and smaller. 

:::co-dum
Perhaps this oversight is because some FSF members don't *use* modern devices like mobile phones, or they are too old or isolated to understand their value. 

"But you don't need devices that size and you should give up your 'Smart' phones like me!"

"Ok, boomer."
:::

[ARM chips](https://duck.com/lite?kae=t&q=ARM chips) have become something of a standard in the IoT world and Linux is usually the pick. But as ARM and RAM chips continue to shrink and [microcontrollers](https://duck.com/lite?kae=t&q=microcontrollers) further fall out of favor due to their limitations compared to a *full* operating system it seems obvious Linux will continue to be installed on smaller and smaller devices. Eventually *any* physical interface will be impossible. Imagine a lapel pin, baseball cap, belt, conference credential, or door key running Linux. 

### We Don't Need No Stinking Network

Many of these devices *do not need network connectivity at all.* In fact, allowing network access to such a device would be a gross and unnecessary security risk, as is the case with most microcontrollers included in consumer products today. Assuming all devices running *any* GPLv3 software will have network access is not just short-sighting, it's fucking idiotic. GPLv3 stops anything licensed with it from being used on *any* modern or future device that does not have an interface.

:::co-btw
The ["smallest Linux computer in the world"](https://duck.com/lite?kae=t&q="smallest Linux computer in the world") right now is still huge compared to what is possible without providing *any* network or upgrade interface.
:::

### GPLv3 Software Libraries are at Greatest Risk

GPLv3 software libraries lose the most.

Most of the [GNU Core Utilities](https://duck.com/lite?kae=t&q=GNU Core Utilities) affected by GPLv3 are wasted bloat when the size of the computer running Linux reaches the size of a door key. But who cares about interactive GNU tools at that point. A product designer can create [Linux From Scratch](https://duck.com/lite?kae=t&q=Linux From Scratch) and leave all the useless GNU "utilities" out of it.

:::co-btw

Most GNU "core" utilities are ancient boomer tech full of ancient
technical debt in need of a safer and more maintainable rewrite anyway
at this point. Many of them have *already* been superseded by better
alternatives like [`k3os`](https://k3os.io) which is based on Alpine
Linux that contains zero GNU code, not even `glibc`. It's just a matter
of time before the entire GNU core library has a plug-n-play
replacement, which will probably be written in Go and/or Rust. Most of
`k3os` and OSes like it are almost entirely in Go these days (except for
the Linux kernel, of course).

:::

However, when the product designer begins to choose or develop software for that tiny, no-interface device anything licensed under GPLv3 is automatically ruled out since it's illegal. Doesn't matter how good it is. It's out.

[Software engineers choosing to release under GPLv3 have permanently banned themselves from many modern and future Linux devices too small to have an interface.]{.important}

Stay relevant. Release your software under GPLv2.

## Unethical License Naming and Release

The [Free Software Foundation's](https://duck.com/lite?kae=t&q=Free Software Foundation's) [dirty treatment of the GPLv3 license roll-out](https://youtu.be/PaKIZ7gJlRU?t=196), which should have clearly required a new name, was completely unethical and sought to *undermine the rights and freedoms* of those who had already decided to use GPLv2 "or later" software already installed. Linus Torvaldz and many others rightfully  abandoned any work or association to the FSF and GNU and now actively seek software written *without* the GPLv3 license at all. 

## GPLv3 is *Bad* for Linux

What and how software is to be installed and upgraded on hardware is *not* an FSF decision. Hell, it has more to do with *hardware* than software. You know, the *S* in FSF. No doubt the issue of corporate control over devices is a very real concern but the resolution and debate is not related to the software at all. It is a separate issue that needs attention. Forcing a *software* license to suddenly reduce the freedoms of hardware manufacturers is *not* the way to resolve it. In fact, it forces the GPLv3 into a hardware licensing space where it simply does not belong --- at all.

:::co-fun
Maybe the Free *Software* Foundation should be legally required to change its name to Free *Software and Hardware* Foundation instead. They are operating way out of their original charter at this point.
:::

All of this means many more device designers are overlooking Linux entirely, not just [Apple](https://duck.com/lite?kae=t&q=Apple GPLv3 GNU) and [Google](https://duck.com/lite?kae=t&q=Google GPLv3 GNU Fuscia). This is really too bad because it is happening at the same time the capacity for smaller devices to run an entire Linux OS is increasing. GNU licenses were already problematic for most companies, big and small --- including those which aren't as evil as others --- but now GNU is simply not an option, and too often they will continue to throw Linux out along with it.

Linux is not dead. GNU just killed itself.

## See Also

* [Linus Explains Why He's Against GPLv3 and FSF](https://youtu.be/PaKIZ7gJlRU)
