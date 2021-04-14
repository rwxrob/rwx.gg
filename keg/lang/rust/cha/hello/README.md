---
Title: Hello Rust! Challenge
Subtitle: Create a Hello World Command in Rust
Category: Challenge
Summary:
    Install the Rust language tools. Write a `hello.rs` program that prints 'Hello World!'. Format it with `rustfmt`. Compile it with `rustc` and run it with `./hello`. If possible, add `rustfmt` to your editor so that it automatically formats your code when you save it. Replace `World` with a variable. Now replace `World` with the first command argument so that `./hello First` prints `Hello First!` and `./hello 'First Last'` prints `Hello First Last!`. Finally, print `Hello World!` if there are no command arguments.
Prereqs:
- /lang/rust/install/
- /tools/linux/mkdir/
- /tools/linux/touch
- /lang/rust/rustc/ 
- /lang/rust/rustfmt/ 
- /lang/rust/how/string
- /tools/editors/vim/how/fmtonleave/
---

Every Rust program *must* have a `main` function.

:::co-tip

Rust --- like a few other languages --- lumps [subroutines](https://duck.com/lite?kae=t&q=subroutines), [procedures](https://duck.com/lite?kae=t&q=procedures), and [functions](https://duck.com/lite?kae=t&q=functions) all together with the `fn` keyword even though they are *very* different from a language construct perspective.

:::

```rust

fn main() {
    println!("Hello World!")
}

```

Strings must be used *inside* of functions. Outside of functions only constants are allowed.

```rust

fn main() {
    let name = "World";
    println!("Hello {}!", name)
}

```

A *vector* in Rust is essentially the same as an *array* in other languages.

```rust

use std::env;

fn main() {
    let args: Vec<String> = env::args().collect();
    println!("Hello {}!", args[1])
}

```

Nothing in Rust can be changed by default. That's what makes the language so safe. You have to tell Rust explicitly that you want to change something with the `mut` keyword.

The `&` takes a pointer reference to the second argument of the `args` vector.

```
use std::env;

fn main() {
    let args: Vec<String> = env::args().collect();
    let mut name = "World";
    if args.len() > 1 {
        name = &args[1];
    }
    println!("Hello {}!", name)
}
```

