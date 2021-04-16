---
Title: Head First C
Subtitle: Annotations
Query: true
---

## Chapter 1

### C Coding Styles

*Just use [RWX C Coding Style](/lang/c#rwx-c-coding-style).*

This chapter might leave you wondering a few things about the coding style being used:

* Why is the [curly bracket](/key/) on its own line after functions?
* Why does other C code in other books like [Modern C](/reviews/books/modernc/) look different?
* Who decided this?
* Which style should I use?

It is really too bad the authors of this book chose to not even weigh in on the different C style guidelines at all, but it is understandable. It seems they are following K&R, which is also followed by the [Linux kernel team](/reviews/books/kernstyle/).

The *only* things you need to know about styles are the following:

* Use same style on existing projects.
* Use whatever you want on new projects.
* Use a formatter and no one cares.

The amount of indentation has *nothing* to do with coding in C until you do it differently than *everyone* else on a collaborative project. That is all you really need to know about styles. Do it the way the project requires, and if it is your project do it however you like. You might even find simply two-space indentation is just fine. 

### C Versions and Compilers

*Just use `gcc --standard=c11 --pedantic`.*

The `gcc` compiler has been the standard for more than thirty years. C99 is the most widely supported version of C even though the standard itself is at C2X now. Most projects and documentation will be using C99.

#### Why not `clang`?

The `clang` tool might boast faster *perceived* compilation times, color terminal support, and very verbose errors but strictly speaking `clang` isn't even a true C compiler. It creates a small bit of binary code that is combined with something called the Low-Level Virtual Machine or LLVM. This is exactly what other LLVM languages do as well (Rust, Crystal, and Julia). To give you a sense of this, the first LLVM was created with `gcc`, the main Linux compiler. Considering that the main reason you are learning C at this point is to understand how languages work and interact with the underlying operating system and machine it is in our best interests to leave out the LLVM, which robs you of the opportunity to actually compile to machine code.
# RWX C Coding Style

The `rwx.gg` project follows a modern C coding style that is more consistent with that of JavaScript, Go, and Rust. It is the same as the [Linux Kernel Coding Style](/reviews/books/kernstyle/) with the following changes:

* Left curly brackets *always* end of line.
* *All* blocks require curly brackets.
* Indent two spaces.

All the other great advice about size and scope of functions and such still very much applies.

### Motivation

* Remain consistent with JavaScript, Go, and Rust
* Present C clearly on very small screens
* Maintain the spirit of K&R style

## See Also

* [There is No Arduino Language](https://hackaday.com/2015/07/28/embed-with-elliot-there-is-no-arduino-language/)
* [O'Reilly Errata](https://www.oreilly.com/catalog/errata.csp?isbn=0636920015482)

