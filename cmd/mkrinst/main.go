package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if err := run(os.Args[1:]); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

var subCommands = map[string]func([]string) error{
	"qpkgconf":  doQpkgConf,
	"agentconf": doAgentConf,
}

func run(argv []string) error {
	if len(argv) < 1 {
		return fmt.Errorf("no subbcommand specified")
	}
	fn, ok := subCommands[argv[0]]
	if !ok {
		return fmt.Errorf("unknown sub command: %s", argv[0])
	}
	return fn(argv[1:])
}
