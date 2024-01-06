//go:build linux

package main

import (
	"os"
	"time"

	"github.com/Shoshta73/legcli/cmd"
	"github.com/Shoshta73/legcli/config"
	"github.com/Shoshta73/legcli/setup"
)

type APPDATA struct {
	initialized bool
	filePaths   setup.ConfigPath
	config      config.ConfigData
}

var appdata APPDATA
var initTime time.Duration

func init() {
	start := time.Now()
	appdata.initialized = false

	var p = func(e error) {
		if e != nil {
			panic(e)
		}
	}

	var useConfigFile string
	var checkFile = func(file string) {
		fi, err := os.Stat(file)
		if err == nil && fi.Mode().IsRegular() {
			appdata.initialized = true
			useConfigFile = file
		}
	}

	appdata.filePaths = *setup.GetPaths()

	cdi, err := os.Stat(appdata.filePaths.ConfigDir)
	if err == nil {
		if cdi.IsDir() {
			checkFile(appdata.filePaths.ConfigFileDefault)
		}
	}

	checkFile(appdata.filePaths.ConfigFileConfig)
	checkFile(appdata.filePaths.ConfigFileBackup)
	if appdata.initialized && useConfigFile != "" {
		cfg, err := config.GetConfigData(useConfigFile)
		p(err)
		appdata.config = *cfg
	}

	initTime = time.Since(start)
}

func main() {
	args := cmd.ParseCliArgs()
	fmt.Println(args)
	if args.Init {
		panic("not implented yet")
	}
}
