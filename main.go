package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/k4rldoherty/rss-blog-aggregator/internal/command"
	"github.com/k4rldoherty/rss-blog-aggregator/internal/config"
	"github.com/k4rldoherty/rss-blog-aggregator/internal/database"
	"github.com/k4rldoherty/rss-blog-aggregator/internal/middleware"
	"github.com/k4rldoherty/rss-blog-aggregator/internal/state"

	// This is the postgres driver which allows the application to talk to the database
	_ "github.com/lib/pq"
)

func main() {
	conf, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v\n", err)
		os.Exit(1)
	}
	// This connnects to the database using the connection string from the config struct
	db, err := sql.Open("postgres", conf.DbUrl)
	if err != nil {
		log.Fatalf("error connecting to database: %v", err)
		os.Exit(1)
	}

	// Uses the generated database package to create a new *database.Queries
	// and stores it in the state struct
	dbQueries := database.New(db)
	appState := state.NewState(&conf, dbQueries)
	appCommands := command.NewCommands()
	err = appCommands.Register("login", command.HandlerLogin)
	if err != nil {
		log.Fatalf("error registering login command: %v\n", err)
		os.Exit(1)
	}
	err = appCommands.Register("register", command.HandlerRegister)
	if err != nil {
		log.Fatalf("error registering register command: %v\n", err)
		os.Exit(1)
	}
	err = appCommands.Register("reset", command.HandlerReset)
	if err != nil {
		log.Fatalf("error registering reset command: %v\n", err)
		os.Exit(1)
	}
	err = appCommands.Register("users", middleware.MiddlewareLoggedIn(command.HandlerUsers))
	if err != nil {
		log.Fatalf("error registering login command: %v\n", err)
		os.Exit(1)
	}
	err = appCommands.Register("agg", command.HandleAgg)
	if err != nil {
		log.Fatalf("error registering agg command: %v\n", err)
		os.Exit(1)
	}
	err = appCommands.Register("addfeed", middleware.MiddlewareLoggedIn(command.HandleAddFeed))
	if err != nil {
		log.Fatalf("error registering addfeed command: %v\n", err)
		os.Exit(1)
	}
	err = appCommands.Register("feeds", command.HandleFeeds)
	if err != nil {
		log.Fatalf("error registering feeds command: %v\n", err)
		os.Exit(1)
	}
	err = appCommands.Register("follow", middleware.MiddlewareLoggedIn(command.HandleFollow))
	if err != nil {
		log.Fatalf("error registering follow command: %v\n", err)
		os.Exit(1)
	}
	err = appCommands.Register("following", middleware.MiddlewareLoggedIn(command.HandleFollowing))
	if err != nil {
		log.Fatalf("error registering follow command: %v\n", err)
		os.Exit(1)
	}
	err = appCommands.Register("unfollow", middleware.MiddlewareLoggedIn(command.HandleUnfollow))
	if err != nil {
		log.Fatalf("error registering unfollow command: %v\n", err)
		os.Exit(1)
	}

	if len(os.Args) < 2 {
		log.Fatalf("%v\n", err)
		os.Exit(1)
	}

	cmd := command.Command{
		Name: os.Args[1],
		Args: os.Args[2:],
	}

	err = appCommands.Run(appState, cmd)
	if err != nil {
		log.Fatalf("error: %v\n", err)
		os.Exit(1)
	}
	os.Exit(0)
}
