package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/larrymacrich/gator/internal/config"
	"github.com/larrymacrich/gator/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	// Read config file
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
		return
	}

	// Open DB
	dbURL := cfg.DBURL
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
		return
	}
	dbQueries := database.New(db)

	// Init CLI State & CMD
	cliState := state{
		db:  dbQueries,
		cfg: cfg,
	}
	cliCommands := commands{
		mapCommands: make(map[string]func(*state, command) error),
	}

	// Register handler
	cliCommands.register("login", handlerLogin)
	cliCommands.register("register", handlerRegister)
	cliCommands.register("reset", handlerReset)

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
