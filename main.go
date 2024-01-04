package main

import (
	"bufio"
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

				cfd := getConfigFileData()
			} else {
				panic(err)
			}
		}
	}
}
