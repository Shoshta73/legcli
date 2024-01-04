package main

import (
	"fmt"
	"os"
	"path"
)

var println = fmt.Println

type APPDATA struct {
	configDir        string
	configFile       string
	configFileBackup string
}

type ARGS struct {
	init  bool
}
var appdata APPDATA
var cliArgs ARGS

func parseArgs(args []string) {
	if args[0] == "init" {
		cliArgs.init = true
	}
}

func init() {
	userConfigDir, err := os.UserConfigDir()
	if err != nil {
		panic(err)
	}
	appdata.configDir = path.Join(userConfigDir, "legcli")
	appdata.configFile = path.Join(appdata.configDir, "config.ini")

	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	appdata.configFileBackup = path.Join(homeDir, "legcli.config.ini")
}
func main() {
	args := os.Args[1:]

	parseArgs(args)
	if cliArgs.init {
		confDirInfo, err := os.Stat(appdata.configDir)

		if err != nil {
			if os.IsNotExist(err) {
				err := os.Mkdir(appdata.configDir, 0755)
				if err != nil {
					panic(err)
				}
			} else {
				panic(err)
			}
		}
	}
}
