package main

import (
	"fmt"
	"log"

	"github.com/BurntSushi/toml"

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

	fmt.Println(config.ConfigEntreis[0].To)
}
