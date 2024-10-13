package main

import (
	"fmt"
	"log"

	"github.com/BurntSushi/toml"

	"moonlight/internal/actions"
	"moonlight/internal/flags"
	"moonlight/internal/types"
)

func main() {
	flagManager := flags.GetInstance()

	flagManager.InstantiateFlags()
	flagManager.ParseFlags()

	var config types.ConfigFile
	_, err := toml.DecodeFile(flagManager.ConfigFile.Value, &config)
	if err != nil {
		log.Fatal(err)
	}

	actionManager := actions.GetInstance()
	actionManager.ConfigFile = config

	switch mode := flagManager.Mode.Value; mode {
	case "generate":
		actionManager.GenerateFiles()
	case "link":
		fmt.Print("Not yet implemented")
	}
}
