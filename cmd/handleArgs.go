package cmd

import (
	"fmt"

	"github.com/Shoshta73/legcli/cmd/initProc"
	"github.com/Shoshta73/legcli/setup"
)

const VERSION = "dev-v0"

func HandleArgs(args *CliArgs, initialized bool, filepaths setup.ConfigPaths) {
	if args.Init {
		ok := initProc.Init(initialized, filepaths)
		if !ok {
			panic("Failed to initialise app")
		}
		return
	}
	if args.Version || args.Info {
		fmt.Printf("legcli version: %s", VERSION)
		if args.Version {
			return
		}
		fmt.Printf("\n")
		fmt.Printf("Authored by: Borna Šoštarić (Shoshta73) https://github.com/Shoshta73\n")
		fmt.Printf("https://github.com/Shoshta73/legcli\n")
		fmt.Printf("\n")
		fmt.Printf("Copyright (c) 2024 Borna Šoštarić and Contributors.\n")

		return
	}
}
