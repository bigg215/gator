package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/bigg215/gator/internal/config"
	"github.com/bigg215/gator/internal/database"

	_ "github.com/lib/pq"
)

type state struct {
	cfg *config.Config
	db  *database.Queries
}

func main() {
	cfg := config.Read()
	db, err := sql.Open("postgres", cfg.DbURL)

	if err != nil {
		log.Fatal(err)
	}

	dbQueries := database.New(db)

	programState := state{
		cfg: &cfg,
		db:  dbQueries,
	}

	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerListUsers)
	cmds.register("agg", handlerAGG)
	cmds.register("addfeed", handlerAddFeed)
	cmds.register("feeds", handlerListFeeds)

	if len(os.Args) < 2 {
		fmt.Println("Usage: cli <command> [args...]")
		os.Exit(1)
	}
	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	err = cmds.run(&programState, command{
		Name: cmdName,
		Args: cmdArgs,
	})

	if err != nil {
		log.Fatal(err)
	}
}
