package main

import (
	"flag"
	"fmt"
	"os"
)

const rc_OK = 0
const rc_ARGUMENT_ERR = 1
const rc_RUNTIME_ERR = 2

type arguments struct {
	filePath string
	list     bool
	params   []string
}

func main() {
	args := fetchArgs()
	os.Exit(realMain(args))
}

func realMain(args *arguments) int {
	if args.filePath == "" {
		fmt.Println("Please set path to YAML file by -f option.")
		showUsage()
		return rc_ARGUMENT_ERR
	}
	if !fileExists(args.filePath) {
		fmt.Println("YAML file does not exists.")
		return rc_RUNTIME_ERR
	}

	parsed, err := ParseFile(args.filePath)
	if err != nil {
		fmt.Printf("YAML parse error: %v\n", err)
		return rc_RUNTIME_ERR
	}

	if args.list {
		properties := List(parsed)
		for _, property := range properties {
			fmt.Printf("%s = %v\n", property.Path, property.Value)
		}
	} else if len(args.params) == 1 {
		value, err := Query(parsed, args.params[0])
		if err != nil {
			fmt.Printf("YAML property get error: %v\n", err)
			return rc_RUNTIME_ERR
		}
		fmt.Printf("%v\n", value)
	} else if len(args.params) == 2 {
		after, err := Set(parsed, args.params[0], args.params[1])
		if err != nil {
			fmt.Printf("YAML property set error: %v\n", err)
			return rc_RUNTIME_ERR
		}
		ioErr := UnparseToFile(args.filePath, after)
		if ioErr != nil {
			fmt.Printf("YAML output error: %v\n", err)
			return rc_RUNTIME_ERR
		}
	}
	return rc_OK
}

func fetchArgs() *arguments {
	args := new(arguments)
	flag.StringVar(&args.filePath, "f", "", "Path to YAML file.")
	flag.BoolVar(&args.list, "l", false, "Show list of properties.")
	flag.BoolVar(&args.list, "list", false, "Show list of properties.")
	flag.Parse()

	args.params = flag.Args()
	return args
}

func showUsage() {
	fmt.Println("Usage:")
	flag.PrintDefaults()
}

func fileExists(filePath string) bool {
	info, err := os.Stat(filePath)
	return err == nil && !info.IsDir()
}
