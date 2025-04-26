package main

import (
	"log"
	"os"

	"github.com/k4rldoherty/rss-blog-aggregator/internal/command"
	"github.com/k4rldoherty/rss-blog-aggregator/internal/config"
	"github.com/k4rldoherty/rss-blog-aggregator/internal/state"
)

func main() {
	conf, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v\n", err)
	}
	appState := state.NewState(&conf)
	appCommands := command.NewCommands()
	err = appCommands.Register("login", command.HandlerLogin)
	if err != nil {
		log.Fatalf("error registering login command: %v\n", err)
	}

	if len(os.Args) < 2 {
		log.Fatalf("not enough arguments provided: %v\n", err)
		os.Exit(1)
	}

	cmd := command.Command{
		Name: os.Args[1],
		Args: os.Args[2:],
	}

	err = appCommands.Run(appState, cmd)
	if err != nil {
		log.Fatalf("not enough arguments provided: %v\n", err)
		os.Exit(1)
	}
}
