//go:build linux

package main

import (
	"fmt"
	"os"
	"path"
	"time"

	"github.com/Shoshta73/legcli/cmd"
	"github.com/Shoshta73/legcli/config"
)

type APPDATA struct {
	initialized       bool
	configDir         string
	configFileDefault string
	configFileConfig  string
	configFileBackup  string

	config config.ConfigData
}

var appdata APPDATA
var initTime time.Duration

func init() {
	start := time.Now()

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

	appdata.initialized = false

	userConfigDir, err := os.UserConfigDir()
	p(err)
	userHomeDir, err := os.UserHomeDir()
	p(err)

	appdata.configDir = path.Join(userConfigDir, "legcli")
	appdata.configFileDefault = path.Join(appdata.configDir, "config.ini")
	appdata.configFileConfig = path.Join(userConfigDir, "legcli.config.ini")
	appdata.configFileBackup = path.Join(userHomeDir, ".legcli.config.ini")

	cdi, err := os.Stat(appdata.configDir)
	if err == nil {
		if cdi.IsDir() {
			checkFile(appdata.configFileDefault)
		}
	}

	checkFile(appdata.configFileConfig)
	checkFile(appdata.configFileBackup)
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
