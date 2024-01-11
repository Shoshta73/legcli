package config

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

var p = fmt.Println

func isNumeric(s string) bool {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}

func GetConfigData() *ConfigData {
AGAIN:
	var data = ConfigData{
		Fullname:       "",
		DefaultLicence: "",
	}

	var selection string

	scanner := bufio.NewScanner(os.Stdin)

NAME:
	p("Enter your full name")
	scanner.Scan()
	data.Fullname = scanner.Text()
	goto CONFIRM

LIC:
	p("Enter your preffereddefault licence")
	p("They are numbered so that you can choose from them")
LICNUM:
	p("Please enter number your preffered licence")
	p("If you leave blank it will default to ISC Licence")
	scanner.Scan()
	selection = scanner.Text()
	if selection == "" {
		data.DefaultLicence = "ISC"
		goto CONFIRM
	}
	if isNumeric(selection) {
		switch selection {

		}
	} else {
		p("Please enter a number")
		goto LICNUM
	}

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
