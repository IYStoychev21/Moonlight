package actions

import (
	"fmt"
	"os"
	"os/exec"
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
	color.Magenta("Coping files to your desired directory")
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

func (am *ActionManager) LinkFiles() {
	color.Magenta("Linking files")
	fmt.Println()

	for i := 0; i < len(am.ConfigFile.ConfigEntreis); i++ {
		currentEntry := am.ConfigFile.ConfigEntreis[i]

		to, _ := utils.ExpandPath(currentEntry.To)
		from, _ := utils.ExpandPath(currentEntry.From)

		err := os.RemoveAll(to)
		if err != nil {
			panic("failed to remove file")
		}

		cmd := exec.Command("ln", "-s", from, to)
		_, err = cmd.Output()
		if err != nil {
			panic("failed to link file")
		}

		color.Green("Successfully Linked: %s", filepath.Base(from))
	}
}
