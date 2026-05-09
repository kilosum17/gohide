package core

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	Encrypt           bool
	Decrypt           bool
	Zip               bool
	Verbose           bool
	Rename            bool
	TestMode          bool
	IgnoreHiddenFiles bool
	MaxDepth          int
	Password          string
	IgnorePattern     string
	IncludePattern    string
}

func ParseArgs() *Config {
	if len(os.Args) < 2 {
		printUsage()
		return nil
	}

	conf := Config{Rename: true, IgnoreHiddenFiles: true}
	args := os.Args[1:]

	for i := 0; i < len(args); i++ {
		item := args[i]

		if strings.HasPrefix(item, "-") {
			switch item {
			case "-i":
				if i+1 < len(args) {
					conf.IncludePattern = args[i+1]
					i++
				}
			case "-g":
				if i+1 < len(args) {
					conf.IgnorePattern = args[i+1]
					i++
				}
			case "-p":
				if i+1 < len(args) {
					conf.Password = args[i+1]
					i++
				}
			case "-m":
				if i+1 < len(args) {
					depth, err := strconv.Atoi(args[i+1])
					if err != nil {
						conf.MaxDepth = depth
					}
					i++
				}
			default:
				parseShorthand(strings.TrimPrefix(item, "-"), &conf)
			}
		} else {
			// If it doesn't start with -, treat as shorthand (e.g., ezn)
			parseShorthand(item, &conf)
		}
	}

	validateConfig(&conf)
	return &conf
}

func parseShorthand(s string, conf *Config) {
	for _, char := range s {
		switch char {
		case 'e':
			conf.Encrypt = true
		case 'd':
			conf.Decrypt = true
		case 'z':
			conf.Zip = true
		case 'v':
			conf.Verbose = true
		case 'n':
			conf.Rename = false
		case 't':
			conf.TestMode = true
		case 'h':
			conf.IgnoreHiddenFiles = false
		}
	}
}

func validateConfig(conf *Config) {
	if conf.Encrypt && conf.Decrypt {
		fmt.Println("Error: Cannot encrypt and decrypt at the same time.")
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("Usage: ./gohide [commands] [options]")
	fmt.Println("\nCommands (can be combined):")
	fmt.Println("  e : Encrypt")
	fmt.Println("  d : Decrypt")
	fmt.Println("  z : Zip result")
	fmt.Println("  v : Verbose")
	fmt.Println("  n : No rename")
	fmt.Println("  t : Run In test mode")
	fmt.Println("  h : Include hidden files & directories")

	fmt.Println("\nOptions:")
	fmt.Println("  -i [pattern] : Include pattern")
	fmt.Println("  -g [pattern] : Ignore pattern")
	fmt.Println("  -p [password]: Provide password (insecure: shows in history)")
	fmt.Println("  -m [depth]: Provide max-recursive depth, defailt is 3")
}
