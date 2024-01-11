//go:build windows

package initProc

import (
	"fmt"
	"runtime"
	"syscall"

	"github.com/Shoshta73/legcli/setup"
)

func Init(initialized bool, filepaths setup.ConfigPaths) bool {
	if initialized {
		fmt.Println("Config file already exists. Skipping init.")
		return true
	}

	var hideFile = func(f string) {
		if runtime.GOOS != "windows" {
			return
		}

		var p = func(err error) {
			if err != nil {
				panic(err)
			}
		}

		f_ptr, err := syscall.UTF16PtrFromString(f)
		p(err)
		attr, err := syscall.GetFileAttributes(f_ptr)
		p(err)

		attr |= syscall.FILE_ATTRIBUTE_HIDDEN
		err = syscall.SetFileAttributes(f_ptr, attr)
		p(err)
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
			hideFile(filepaths.ConfigFileBackup)
			return true
		}
	}

	return false
}
