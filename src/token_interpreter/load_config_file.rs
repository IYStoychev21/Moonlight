use serde::{Serialize, Deserialize};
use serde_yaml::{self};

#[derive(Debug, Serialize, Deserialize)]
struct Config {
    config_paths: Vec<String>
}

pub fn load(path: String) -> Vec<String> {
    let file = std::fs::File::open(path).expect("Could not open file.");
    let paths: Config = serde_yaml::from_reader(file).expect("Could not read values.");

    return paths.config_paths;
}
