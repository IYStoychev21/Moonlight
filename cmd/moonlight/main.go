package main

import (
	"fmt"

	"moonlight/internal/flags"
)

func main() {
	flagManager := flags.GetInstance()

	flagManager.InstantiateFlags()
	flagManager.ParseFlags()

	fmt.Println(flagManager.ConfigFile.Value)
	fmt.Println(flagManager.DotfilesDirectory.Value)
}
