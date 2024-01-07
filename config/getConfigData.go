package config

import (
	"bufio"
	"fmt"
	"os"
)

var p = fmt.Println

func GetConfigData() *ConfigData {
AGAIN:
	var data = ConfigData{
		Fullname:       "",
		DefaultLicence: "",
	}

	scanner := bufio.NewScanner(os.Stdin)

NAME:
	p("Enter your full name")
	scanner.Scan()
	data.Fullname = scanner.Text()
	goto CONFIRM

LIC:
	p("Enter your preffereddefault licence")
	scanner.Scan()
	data.DefaultLicence = scanner.Text()

CONFIRM:
	if data.DefaultLicence == "" {
		goto LIC
	}
	p("Fullname:", data.Fullname)
	p("Default licence:", data.DefaultLicence)
	p("Is this OK?")
	p("Change:\n" + "Fullname: name | n\n" + "Default licence: licence | l\n" + "All: all | a\n")
	scanner.Scan()
	choice := scanner.Text()

	switch choice {
	case "name":
		fallthrough
	case "n":
		goto NAME
	case "licence":
		fallthrough
	case "l":
		goto LIC
	case "all":
		fallthrough
	case "a":
		goto AGAIN
	default:
		return &data
	}
}
