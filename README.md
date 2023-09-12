# Moonlight

<p center>Moonlight is a simple dotfiles management tool written in Rust</p>

## Installation
For now, the only way to install Moonlight is via the installation script provided in this repo

1. Clone the repo
```
git clone https://github.com/IYStoychev21/Moonlight
```
2. Cd into the directory
```
cd Moonlight
```
3. Give permission to execute the install.sh script
```
sudo chmod +x install.sh
```
4. Run the install.sh script
```
sudo ./install.sh
```

## Usage
First, you need a config.yml file where you specify the path to all of your configuration. Make sure to backup the config too because it is used to link the files later

Example config.yml file:
```
config_paths:
  - ~/.config/nvim
  - ~/.config/alacritty
  - ~/.config/fish
```
Moonlight has two main functionalities
1. Backing up your configs to specific location
  - To do this you need to give Moonlight the config file you created and the location where to store the files
    ```
    moonlight -g ~/example_path_to_config_file -dd ~/dotfiles_directory_path
    ```
2. Doing the symbolic links between the actuall files and the locations they are supposed to be.
    ```
    moonlight -l ~/example_path_to_config_file -dd ~/dotfiles_directory_path
    ```
