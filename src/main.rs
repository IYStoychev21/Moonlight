use std::env;
use std::collections::HashMap;
mod tokenizer;

fn main() {
    let args: Vec<String> = env::args().collect();
    let tokens: Vec<HashMap<String, String>> = tokenizer::tokenize(args.clone());
    println!("{:?}", tokens)
}
