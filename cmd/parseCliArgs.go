package cmd

import (
	"flag"
)

type CliArgs struct {
	Init    bool
	Verbose bool
	Version bool
	Name    string
	Licence string

	// for internal use only
	InitBenchmark bool
}

func ParseCliArgs() *CliArgs {
	cliArgs := new(CliArgs)

	flag.BoolVar(&cliArgs.Init, "init", false, "initialise app and create config file")
	flag.BoolVar(&cliArgs.Init, "i", false, "initialise app and create config file")
	flag.BoolVar(&cliArgs.Verbose, "verbose", false, "verbose output")
	flag.BoolVar(&cliArgs.Verbose, "v", false, "verbose output")
	flag.BoolVar(&cliArgs.InitBenchmark, "init-benchmark", false, "log the time it took to initialize the app")
	flag.StringVar(&cliArgs.Name, "name", "", "name of the user to override the default name")
	flag.StringVar(&cliArgs.Name, "n", "", "name of the user to override the default name")
	flag.StringVar(&cliArgs.Licence, "licence", "", "licence to override the default licence")
	flag.StringVar(&cliArgs.Licence, "l", "", "licence to override the default licence")
	flag.StringVar(&cliArgs.Licence, "lic", "", "licence to override the default licence")
	flag.BoolVar(&cliArgs.Verbose, "version", false, "Show version and exit")
	flag.BoolVar(&cliArgs.Verbose, "V", false, "Show version and exit")

	flag.Usage = func() {
		println("Usage: legcli [options]")
		println("Options:")
		println("   --help    | -h      Print this help message")
		println("   --init    | -i      Initialise app and create config file")
		println("   --verbose | -v      Verbose output")
		println("   --name    | -n      Name of the user to override the default name")
		println("   --licence |         ")
		println("               -l      ")
		println("               -lic    ")
		println("                       Licence to override the default licence")
		println("   --version | -V      Show version and exit")
	}

	flag.Parse()
	return cliArgs
}
