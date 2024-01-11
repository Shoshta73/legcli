//go:build windows

package main

import (
	"fmt"
	"os"
	"time"

	"github.com/Shoshta73/legcli/cmd"
	"github.com/Shoshta73/legcli/config"
	"github.com/Shoshta73/legcli/setup"
)

type APPDATA struct {
	initialized bool
	filePaths   setup.ConfigPaths
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

	if !appdata.initialized {
		checkFile(appdata.filePaths.ConfigFileConfig)
	}
	if !appdata.initialized {
		checkFile(appdata.filePaths.ConfigFileBackup)
	}
	if appdata.initialized && useConfigFile != "" {
		cfg, err := config.GetConfigFileContents(useConfigFile)
		p(err)
		appdata.config = *cfg
	}

	initTime = time.Since(start)
}

func main() {
	if !appdata.initialized && len(os.Args) == 1 {
		fmt.Println("Config file not found. Run 'legcli --init' to create one.")
		return
	}
	if len(os.Args) == 1 {
		cmd.CreateDefaultLicence(appdata.config.Fullname, appdata.config.DefaultLicence)
		return
	}
	args := cmd.ParseCliArgs()
	if args.InitBenchmark {
		fmt.Println(initTime)
	}
	cmd.HandleArgs(args, appdata.initialized, appdata.filePaths)
}
