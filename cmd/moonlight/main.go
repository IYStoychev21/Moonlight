package main

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/fatih/color"

	"moonlight/internal/actions"
	"moonlight/internal/flags"
	"moonlight/internal/types"
)

func main() {
	if os.Geteuid() == 0 {
		color.Red("DO NOT RUN AS SUPER USER")
		return
	}

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
		actionManager.LinkFiles()
	}
}
