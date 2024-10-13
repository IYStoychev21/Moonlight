package actions

import (
	"fmt"
	"path/filepath"
	"sync"

	"moonlight/internal/types"
	"moonlight/internal/utils"

	"github.com/fatih/color"
)

type ActionManager struct {
	ConfigFile types.ConfigFile
}

var (
	instance *ActionManager
	once     sync.Once
)

func GetInstance() *ActionManager {
	once.Do(func() {
		instance = &ActionManager{}
	})

	return instance
}

func (am *ActionManager) GenerateFiles() {
	color.Magenta("Coping file to your desired directory")
	fmt.Println()

	for i := 0; i < len(am.ConfigFile.ConfigEntreis); i++ {
		currentEntry := am.ConfigFile.ConfigEntreis[i]

		src, _ := utils.ExpandPath(currentEntry.To)
		dst, _ := utils.ExpandPath(currentEntry.From)

		err := utils.CopyPath(src, dst)
		if err != nil {
			panic(err)
		}

		color.Green("Successfully Copied: %s", filepath.Base(currentEntry.From))
	}
}
