package main

import (
	"log"
	"os"

	"github.com/larrymacrich/gator/internal/config"
)

func main() {
	// Read config file
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
		return
	}

	// Init CLI State & CMD
	cliState := state{
		cfg: cfg,
	}
	cliCommands := commands{
		mapCommands: make(map[string]func(*state, command) error),
	}

	// Register handler
	cliCommands.register("login", handlerLogin)

	// Setup CMD
	cliArgs := os.Args
	if len(cliArgs) < 2 {
		log.Fatal("Missing command.")
		return
	}
	cmdName := cliArgs[1]
	cmdArgs := cliArgs[2:]
	cliCommand := command{
		name: cmdName,
		args: cmdArgs,
	}

	// Run CMD
	err = cliCommands.run(&cliState, cliCommand)
	if err != nil {
		log.Fatal(err)
		return
	}

}
