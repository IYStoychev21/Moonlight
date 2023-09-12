use std::collections::HashMap;
pub mod load_config_file;

/*
 * -g (generate) - generate config file (copy configs to specified directory)
 * -dd (dotfiles directory) - directory where the dotfiles are stored 
 * -l (link) - does the symboly links between the dotfiles directories and the config directories
 * */

pub fn interprete(tokens: Vec<HashMap<String, String>>) -> () {
    let mut paths: HashMap<String, String> = HashMap::new();
    let mut creation_flag: &str = "0";

    for i in tokens {
        if i["flag"] == "-g" {
            creation_flag = "-g";
            paths.insert("config".to_string(), i["path"].to_string()); 
        } else if i["flag"] == "-l" {
            creation_flag = "-l";
            paths.insert("config".to_string(), i["path"].to_string()); 
        } else if i["flag"] == "-dd" {
            paths.insert("dotfiles_dirs".to_string(), i["path"].to_string()); 
        }
    }

    let directory_paths: Vec<String> = load_config_file::load(paths["config"].to_string());
    println!("{:?}", directory_paths);
}
