use std::env;
use std::collections::HashMap;
mod tokenizer;
mod token_interpreter;

fn main() {
    let args: Vec<String> = env::args().collect();
    let tokens: Vec<HashMap<String, String>> = tokenizer::tokenize(args);
    token_interpreter::interprete(tokens);
}
