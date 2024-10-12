package flags

import (
	"sync"

	"moonlight/internal/types"

	"github.com/spf13/pflag"
)

type FlagManager struct {
	ConfigFile        *types.Flag
	DotfilesDirectory *types.Flag
}

var (
	instance *FlagManager
	once     sync.Once
)

func GetInstance() *FlagManager {
	once.Do(func() {
		instance = &FlagManager{}
	})

	return instance
}

func (fm *FlagManager) InstantiateFlags() {
	fm.ConfigFile = &types.Flag{
		Name:         "conf",
		DefaultValue: "~/dotfiles/config.toml",
		Description:  "path to the config file",
	}

	fm.DotfilesDirectory = &types.Flag{
		Name:         "dir",
		DefaultValue: "~/dotfiles",
		Description:  "path to the directory storing your dotfiles",
	}
}

func (fm *FlagManager) ParseFlags() {
	pflag.StringVar(&fm.ConfigFile.Value, fm.ConfigFile.Name, fm.ConfigFile.DefaultValue, fm.ConfigFile.Description)
	pflag.StringVar(&fm.DotfilesDirectory.Value, fm.DotfilesDirectory.Name, fm.DotfilesDirectory.DefaultValue, fm.DotfilesDirectory.Description)
	pflag.Parse()
}
