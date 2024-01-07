package cmd

import (
	"flag"
)

type CliArgs struct {
	Init          bool
	Verbose       bool
	InitBenchmark bool
	Name          string
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

	flag.Usage = func() {
		println("Usage: legcli [options]")
		println("Options:")
		println("   --help    | -h      Print this help message")
		println("   --init    | -i      Initialise app and create config file")
		println("   --verbose | -v      Verbose output")
		println("   --name    | -n      Name of the user to override the default name")
	}

	flag.Parse()
	return cliArgs
}
