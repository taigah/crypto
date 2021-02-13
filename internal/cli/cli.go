package cli

import (
	"fmt"
	"log"
	"os"

	"github.com/taigah/crypto/internal/crypto"
)

func printHelp() {
	fmt.Print(`
Usage: crypto [options] <command>

Options:

	-h,--help	display help

Commands:

	crypto price <pair>	display pair price
	crypto ls		display the available pairs

`)
}

func parse() (command string, commandArgs []string) {
	for i, arg := range os.Args {
		if i == 0 {
			continue
		} else if arg[0] == '-' {
			if arg == "--help" || arg == "-h" {
				printHelp()
				os.Exit(0)
			} else {
				log.Fatalf("Unrecognized option %v\n", arg)
			}
		} else {
			if command == "" {
				command = arg
			} else {
				commandArgs = append(commandArgs, arg)
			}
		}
	}
	return
}

// Run cli
func Run() {
	command, commandArgs := parse()

	if command == "price" {
		pair := commandArgs[0]
		pairPrice, err := crypto.GetPairPrice(pair)
		if err != nil {
			log.Fatalf("%v", err)
		}
		fmt.Printf("%v\n", pairPrice)
	} else if command == "ls" {
		pairs, err := crypto.GetPairList()
		if err != nil {
			log.Fatalf("%v", err)
		}
		for _, pair := range pairs {
			fmt.Printf("%v\n", pair)
		}
	} else {
		log.Fatalf("Unrecognized command")
	}
}
