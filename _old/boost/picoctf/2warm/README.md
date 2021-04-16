---
Title: 'PicoCTF: 2Warm'
Category: General Skills
Points: 50
---

Can you convert the number 42 (base 10) to binary (base 2)? 

[What is a Base?](/what/base/)

We could just use a [binary
converter](https://duck.com/lite?kd=-1&kp=-1&q=binary converter), but
let's actually do the counting, at least for a while.

Base 10 (Decimal)|Base 2 (Binary)|Base 16 (Hexadecimal)
-|-|-
0|0|0
1|1|1
2|10|2
3|11|3
4|100|4
5|101|5
6|110|6
7|111|7
8|1000|8
9|1001|9
10|1010|a
11|1011|b
12|1100|c
13|1101|d
14|1110|e
15|1111|f
16|10000|10
17|10001|11

Let's use `bc` (which is included on all UNIX/Linux systems) to do the
conversion.

All we have to do is set the output base (`obase`) to 2.

```
$ bc
bc 1.07.1
Copyright 1991-1994, 1997, 1998, 2000, 2004, 2006, 2008, 2012-2017 Free
Software Foundation, Inc.
This is free software with ABSOLUTELY NO WARRANTY.
For details type `warranty'. 
obase=2
42
101010
^C
(interrupt) use quit to exit.
```
