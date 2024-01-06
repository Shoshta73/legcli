package setup

import (
	"os"
	"path"
	"runtime"
)

type ConfigPath struct {
	ConfigDir         string
	ConfigFileDefault string
	ConfigFileConfig  string
	ConfigFileBackup  string
}

func GetPaths() *ConfigPath {
	var p = func(e error) {
		if e != nil {
			panic(e)
		}
	}

	cp := new(ConfigPath)

	userConfigDir, err := os.UserConfigDir()
	p(err)
	userHomeDir, err := os.UserHomeDir()
	p(err)

	cp.ConfigDir = path.Join(userConfigDir, "legcli")
	cp.ConfigFileDefault = path.Join(cp.ConfigDir, "config.ini")
	cp.ConfigFileConfig = path.Join(userConfigDir, "legcli.config.ini")
	if runtime.GOOS != "windows" {
		cp.ConfigFileBackup = path.Join(userHomeDir, ".legcli.config.ini")
	} else {
		cp.ConfigFileBackup = path.Join(userHomeDir, "legcli.config.ini")
	}

	return cp
}
