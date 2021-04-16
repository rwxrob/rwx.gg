use std::env;

fn main() {
    // TODO change to env::args().nth(1)
    let args: Vec<String> = env::args().collect();
    let mut name = "World";
    if args.len() > 1 {
        name = &args[1];
    }
    println!("Hello {}!", name)
}
