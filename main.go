package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/simonlewi/gator/internal/config"
	"github.com/simonlewi/gator/internal/database"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println("Error reading config:", err)
		return
	}

	db, err := sql.Open("postgres", cfg.DBUrl)
	if err != nil {
		log.Fatal("Error connecting to database:", err)
		return
	}

	dbQueries := database.New(db)

	appState := &state{
		cfg: &cfg,
		db:  dbQueries,
	}

	cmds := commands{
		commandMap: make(map[string]func(*state, command) error),
	}
	cmds.register("login", loginHandler)
	cmds.register("register", registerHandler)
	cmds.register("reset", resetHandler)
	cmds.register("users", getUsersHandler)
	cmds.register("agg", aggCommand)
	cmds.register("addfeed", middlewareLoggedIn(addFeedHandler))
	cmds.register("feeds", getFeeds)
	cmds.register("follow", middlewareLoggedIn(followHandler))
	cmds.register("following", middlewareLoggedIn(followingHandler))

	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
		return
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	err = cmds.run(appState, command{Name: cmdName, Args: cmdArgs})
	if err != nil {
		log.Fatal(err)
	}
}
