package setup

import (
	"os"
	"path"
	"runtime"
)

type ConfigPaths struct {
	ConfigDir         string
	ConfigFileDefault string
	ConfigFileConfig  string
	ConfigFileBackup  string
}

func GetPaths() *ConfigPaths {
	var p = func(e error) {
		if e != nil {
			panic(e)
		}
	}

	cp := new(ConfigPaths)

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
