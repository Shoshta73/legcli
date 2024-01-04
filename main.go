package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"runtime"
	"syscall"

	"gopkg.in/ini.v1"
)

var println = fmt.Println

type APPDATA struct {
	initialized      bool
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
	appdata.initialized = false
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
	if runtime.GOOS == "windows" {
		appdata.configFileBackup = path.Join(homeDir, "legcli.config.ini")
	} else {
		appdata.configFileBackup = path.Join(homeDir, ".legcli.config.ini")
	}

	confDirInfo, err := os.Stat(appdata.configDir)
	if os.IsExist(err) {
		if confDirInfo.IsDir() {
			confFileData, _ := os.Stat(appdata.configFile)
			if confFileData.Mode().IsRegular() {
				appdata.initialized = true
			}
		}
	}
	confFileData, err := os.Stat(appdata.configFileBackup)
	if os.IsExist(err) {
		if confFileData.Mode().IsRegular() {
			appdata.initialized = true
		}
	}
}

type CFGDATA struct {
	fullname       string
	defaultLicence string
}

func getConfigFileData() CFGDATA {
AGAIN:
	var data CFGDATA

	scanner := bufio.NewScanner(os.Stdin)
	var p = fmt.Println

	p("Enter your full name")
	scanner.Scan()
	data.fullname = scanner.Text()
	p("Enter your default licence")
	scanner.Scan()
	data.defaultLicence = scanner.Text()

	p("Is this OK? (y/n)")
	p("Full name:", data.fullname)
	p("Default licence:", data.defaultLicence)
	scanner.Scan()
	ok := scanner.Text()

	switch ok {
	case "n":
		fallthrough
	case "N":
		fallthrough
	case "no":
		fallthrough
	case "NO":
		goto AGAIN
	case "y":
		fallthrough
	case "Y":
		fallthrough
	case "yes":
		fallthrough
	case "YES":
		return data
	default:
		return data
	}
}

func writeConfigFile(file string, d CFGDATA) {
	cfg := ini.Empty()

	section, err := cfg.NewSection("Default")
	if err != nil {
		println("Error creating section:", err)
		panic(err)
	}
	section.NewKey("fullname", d.fullname)
	section.NewKey("default_licence", d.defaultLicence)

	err = cfg.SaveTo(file)
	if err != nil {
		panic(err)
	}
}

func main() {
	args := os.Args[1:]

	parseArgs(args)
	if cliArgs.init {
		if appdata.initialized {
			return
		}
		confDirInfo, err := os.Stat(appdata.configDir)

		var hideFile = func(f string) {
			if runtime.GOOS != "windows" {
				return
			}

			var p = func(e error) {
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

		var writeFile = func(f string) {
			cfd := getConfigFileData()
			writeConfigFile(f, cfd)
			hideFile(f)
		}

		if err != nil {
			if os.IsNotExist(err) {
				err := os.Mkdir(appdata.configDir, 0755)
				if err != nil {
					panic(err)
				}
				writeFile(appdata.configFile)
				return
			} else {
				panic(err)
			}
		}

		if confDirInfo.Mode().IsDir() {
			writeFile(appdata.configFileBackup)
			return
		}
		if confDirInfo.Mode().IsRegular() {
			writeFile(appdata.configFileBackup)
			return
		}
	}
}
