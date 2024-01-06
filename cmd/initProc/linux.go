//go:build linux

package initProc

import (
	"fmt"

	"github.com/Shoshta73/legcli/setup"
)

func Init(initialized bool, filepaths setup.ConfigPath) bool {
	if initialized {
		fmt.Println("Config file already exists. Skipping init.")
		return true
	}

	cdi := GetDirInfo(filepaths.ConfigDir)
	switch tryConfigDir(cdi) {
	case true:
		if checkFile(filepaths.ConfigFileDefault) {
			getConfigAndWriteFile(filepaths.ConfigFileDefault)
			return true
		}
	case false:
		if checkFile(filepaths.ConfigFileConfig) {
			getConfigAndWriteFile(filepaths.ConfigFileConfig)
			return true
		}
		if checkFile(filepaths.ConfigFileBackup) {
			getConfigAndWriteFile(filepaths.ConfigFileBackup)
			return true
		}
	}

	return false
}
