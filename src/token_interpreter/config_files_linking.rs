use std::process::Command;
use whoami;

pub fn link(link_to_path: String, link_from_path: String) -> () {
    let new_str: &str = &(link_to_path.clone()[1..]);
    let formated: String = format!("/home/{}{}", whoami::username(), new_str);

    Command::new("sudo").args(["rm", "-rf", &formated]).status().expect("failed to launch process");

    let mut app_name: String = formated.chars().rev().take_while(|&ch| ch != '/').collect::<String>();
    app_name = app_name.chars().rev().collect::<String>();

    let dotfile_location: String = format!("{}{}", link_from_path, app_name);

    Command::new("sudo").args(["ln", "-s", &dotfile_location, &formated]).status().expect("failed to launch process");
}
