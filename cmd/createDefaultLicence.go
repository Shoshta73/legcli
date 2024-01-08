package cmd

import "github.com/Shoshta73/legcli/licences"

func CreateDefaultLicence(fn string, l string) {
	if fn == "" || l == "" {
		panic("Something happened to your config file.")
	}

	switch l {
	case "AGPL-3.0":
	case "Apache-2.0":
	case "BSD-2.0-Clause":
	case "BSD-3.0-Clause":
	case "BSD-4.0-Clause":
	case "BSL-1.0":
	case "CC-BY-4.0":
	case "CC0-1.0":
	case "EPL-2.0":
	case "GPL-2.0":
	case "GPL-3.0":
	case "ISC":
	case "LGPL-2.1":
	case "LGPL-3.0":
	case "MIT":
	case "MPL-2.0":
	case "Unlicense":
	case "WTFPL":
		licences.WTFPL(fn)
	}
}
