package cli

import (
	"database/sql"
	"errors"
	"os"

	"github.com/larrymacrich/gator/internal/config"
	"github.com/larrymacrich/gator/internal/database"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func RunCommand() error {
	// Read config file
	cfg, err := config.Read()
	if err != nil {
		return err
	}

	// Open DB
	dbURL := cfg.DBURL
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return err
	}
	dbQueries := database.New(db)

	// Init CLI State & CMD
	cliState := state{
		db:  dbQueries,
		cfg: cfg,
	}
	cliCommands := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}

	// Register handler
	cliCommands.register("login", handlerLogin)
	cliCommands.register("register", handlerRegister)
	cliCommands.register("reset", handlerReset)
	cliCommands.register("users", handlerList)
	cliCommands.register("agg", handlerAgg)
	cliCommands.register("addfeed", handlerAddFeed)

	// Setup CMD
	cliArgs := os.Args
	if len(cliArgs) < 2 {
		return errors.New("Missing command.")
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
		return err
	}

	return nil
}
