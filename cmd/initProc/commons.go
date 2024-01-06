package initProc

import (
	"fmt"
	"io/fs"
	"os"

	"github.com/Shoshta73/legcli/config"
	"gopkg.in/ini.v1"
)

func checkFile(file string) bool {
	fi, err := os.Stat(file)
	if err == nil {
		if fi.Mode().IsRegular() {
			return false
		}
		if fi.Mode().IsDir() {
			return false
		}
	} else {
		if os.IsNotExist(err) {
			return true
		} else {
			panic(err)
		}
	}
	return true
}

func tryConfigDir(di fs.FileInfo) bool {
	if di.Mode().IsRegular() {
		return false
	}
	if di.Mode().IsDir() {
		return true
	}
	return false
}

func GetDirInfo(d string) fs.FileInfo {
	di, err := os.Stat(d)
	if err != nil {
		fmt.Println("Error getting directory info: ", err)
		panic(err)
	}
	return di
}

var getConfigAndWriteFile = func(f string) {
	d := config.GetConfigData()

	cfg := ini.Empty()

	section, err := cfg.NewSection("Default")
	if err != nil {
		fmt.Println("Error writing config file: ", err)
		panic(err)
	}

	section.NewKey("fullname", d.Fullname)
	section.NewKey("default_licence", d.DefaultLicence)

	err = cfg.SaveTo(f)
	if err != nil {
		panic(err)
	}
}
