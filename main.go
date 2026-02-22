package main

import (
	"fmt"
	"gator-go/internal/config"
	"os"

	"database/sql"
	"gator-go/internal/database"

	_ "github.com/lib/pq"
)

func main() {

	state, err := newState()
	if err != nil {
		fmt.Printf("Error initializing state: %v\n", err)
		os.Exit(1)
	}
	db, err := sql.Open("postgres", state.Config.DBURL)
	if err != nil {
		fmt.Printf("Error opening DB connection: %v\n", err)
		os.Exit(1)
	}

	state.db = database.New(db)

	commands := NewCommands()

	commands.Register("login", loginHandler)
	commands.Register("register", registerHandler)
	commands.Register("reset", resetHandler)
	commands.Register("users", usersHandler)
	commands.Register("agg", aggHandler)
	commands.Register("addfeed", addFeedHandler)
	commands.Register("feeds", feedsHandler)

	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("No command provided")
		os.Exit(1)
	}

	cmd := command{
		Name: args[0],
		Args: args[1:],
	}

	err = commands.run(state, cmd)
	if err != nil {
		fmt.Printf("Error executing command: %v\n", err)
		os.Exit(1)
	}
}

type state struct {
	Config *config.Config
	db     *database.Queries
}

func newState() (*state, error) {
	cfg, err := config.Read()
	if err != nil {
		return nil, err
	}

	return &state{
		Config: &cfg,
	}, nil
}
