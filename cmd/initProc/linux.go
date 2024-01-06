//go:build linux

package initProc

import (
	"github.com/Shoshta73/legcli/setup"
)

func Init(initialized bool, filepaths setup.ConfigPath) {
	if initialized {
		return
	}

	cdi := GetDirInfo(filepaths.ConfigDir)
	switch tryConfigDir(cdi) {
	case true:
	case false:
	}
}
