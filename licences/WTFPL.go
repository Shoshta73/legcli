package licences

var WTFPL = func(name string) {
	contents := []byte(
		"            DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE\n" +
			"                    Version 2, December 2004\n" +
			"\n" +
			" Copyright (C) " + year + " " + name + " \n" + // TODO: add email"<sam@hocevar.net>\n" +
			"\n" +
			" Everyone is permitted to copy and distribute verbatim or modified\n" +
			" copies of this license document, and changing it is allowed as long\n" +
			" as the name is changed.\n" +
			"\n" +
			"            DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE\n" +
			"   TERMS AND CONDITIONS FOR COPYING, DISTRIBUTION AND MODIFICATION\n" +
			"\n" +
			"  0. You just DO WHAT THE FUCK YOU WANT TO.\n",
	)
	WriteFile(contents)
}
