# Moonlight

<p align="center">  Moonlight is a simple dotfiles management tool written in golang  <p>


## Prerequisites
  - [Golang](https://go.dev/)

### For Arch
```
sudo pacman -S go
```
    
## Installation
### For Arch use the AUR
```
yay -S moonlight-git
```

### Build from source
```
git clone https://github.com/IYStoychev21/Moonlight
cd Moonlight
mkdir build
go build -o build/moonlight cmd/moonlight/main.go
cp build/moonlight /usr/bin
```

## Usage
### Configuration
``` toml
# config.toml

[[config_locations]]
from = "~/dotfiles/nvim"
to = "~/.config/nvim"

[[config_locations]]
from = "~/dotfiles/fish"
to = "~/.config/fish"
```

- from - where the files are or should be located in your dotfiles directory
- to - where the files should link to

### Generating
After you've defined your config file.
You can use Moonlight to automatically copy them to your dotfiles directory
``` 
moonlight --config /path/to/your/config/file --dir /path/to/you/dotfiles/directory --mode generate
```

### Linking
After you've generated you files you can use Moonlight to link them automatically
```
moonlight --config /path/to/your/config/file --dir /path/to/you/dotfiles/directory --mode link
```
