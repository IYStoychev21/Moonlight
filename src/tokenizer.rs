use std::collections::HashMap;

pub fn tokenize(args: Vec<String>) -> Vec<HashMap<String, String>>{
    let mut tokens: Vec<HashMap<String, String>> = Vec::new();

    if (args[1].as_bytes()[0]) as char != '-' {
        panic!("\nError: This fist argument must be a flag.\nTip: Try beginning the command with a flag");
    }

    let mut is_there_flag: bool = false;
    let mut current_flag: String = "0".to_string();

    for i in args {
       if (i.as_bytes()[0]) as char == '-' {
           current_flag = i;

           if is_there_flag {
               panic!("\nError: A flag can not be followed by another flag.\nTip: Try adding path after the flag");
           }

           is_there_flag = true;
       } else {
           if is_there_flag {
                let mut token: HashMap<String, String> = HashMap::new();
                token.insert("flag".to_string(), current_flag.to_string());
                token.insert("path".to_string(), i.to_string());
                tokens.push(token);
           }

           is_there_flag = false;
       }
    }

    return tokens;
}
