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

var appdata APPDATA

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
