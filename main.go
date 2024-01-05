package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path"
	"runtime"
	"syscall"

	"gopkg.in/ini.v1"
)

var println = fmt.Println

// app data
type APPCFG struct {
	fullname       string
	defaultLicence string
}

type APPDATA struct {
	initialized      bool
	configDir        string
	configFile       string
	configFileBackup string
	cfg              APPCFG
}

var appdata APPDATA

// command Line argumens
type ARGS struct {
	init          bool
}

var cliArgs ARGS = ARGS{
	init:          false,
}

func parseArgs() {
	flag.BoolVar(&cliArgs.init, "init", false, "initialise app and create config file")
	flag.BoolVar(&cliArgs.init, "i", false, "initialise app and create config file")
	flag.BoolVar(&cliArgs.initBenchmark, "init-benchmark", false, "log the time it took to initialize the app")

	flag.Usage = func() {
		println("Usage: legcli [options]")
		println("Options:")
		println("   --init    | -i      Initialise app and create config file")
	}

	flag.Parse()
}

func getConfigData(f string) {
	cfg, err := ini.Load(f)
	if err != nil {
		panic(err)
	}
	section := cfg.Section("Default")
	appdata.cfg.fullname = section.Key("fullname").String()
	appdata.cfg.defaultLicence = section.Key("default_licence").String()
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

	var useConfigFile string

	confDirInfo, err := os.Stat(appdata.configDir)
	if err == nil {
		if confDirInfo.IsDir() {
			confFileData, err := os.Stat(appdata.configFile)
			if err == nil && confFileData.Mode().IsRegular() {
				appdata.initialized = true
				useConfigFile = appdata.configFile
			}
		}
	}
	confFileData, err := os.Stat(appdata.configFileBackup)
	if err == nil {
		if confFileData.Mode().IsRegular() {
			appdata.initialized = true
			useConfigFile = appdata.configFileBackup
		}
	}
	if appdata.initialized {
		getConfigData(useConfigFile)
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
	}
	parseArgs()
	if cliArgs.init {
		if appdata.initialized {
			return
		}
		confDirInfo, err := os.Stat(appdata.configDir)

		var writeFile = func(f string) {
			cfd := getConfigFileData()
			writeConfigFile(f, cfd)
		}

		if err != nil {
			if os.IsNotExist(err) {
				println("Creating config dir at", appdata.configDir)
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

		if confDirInfo.Mode().IsDir() {
			writeFile(appdata.configFileBackup)
			hideFile(appdata.configFileBackup)
			return
		}
		if confDirInfo.Mode().IsRegular() {
			writeFile(appdata.configFileBackup)
			hideFile(appdata.configFileBackup)
			return
		}
	}
}
