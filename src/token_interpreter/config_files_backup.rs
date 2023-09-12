use std::process::Command;
use whoami;

pub fn backup(for_backup: String, dotfile_directory: String) -> () {
    let new_str: &str = &(for_backup.clone()[1..]);
    let formated: String = format!("/home/{}{}", whoami::username(), new_str);
    Command::new("cp").args(["-r", &formated, &dotfile_directory]).spawn().expect("failed to launch process");
}
