use std::collections::HashMap;

/*
 * -g (generate) - generate config file (copy configs to specified directory)
 * -dd (dotfiles directory) - directory where the dotfiles are stored 
 * -l (link) - does the symboly links between the dotfiles directories and the config directories
 * */

pub fn interprete(tokens: Vec<HashMap<String, String>>) -> () {
    let mut paths: HashMap<String, String> = HashMap::new();

    for i in tokens {
        if i["flag"] == "-g" {
           paths.insert("config".to_string(), i["path"].to_string()); 
        } else if i["flag"] == "-dd" {
           paths.insert("dotfiles_dirs".to_string(), i["path"].to_string()); 
        } else if i["flag"] == "-l" {
           paths.insert("link_dirs".to_string(), i["path"].to_string()); 
        }
    }

    println!("{:?}", paths);
}
