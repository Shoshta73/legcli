package cmd

import (
	"github.com/Shoshta73/legcli/cmd/initProc"
	"github.com/Shoshta73/legcli/setup"
)

func HandleArgs(args *CliArgs, initialized bool, filepaths setup.ConfigPaths) {
	if args.Init {
		ok := initProc.Init(initialized, filepaths)
		if !ok {
			panic("Failed to initialise app")
		}
		return
	}
}
