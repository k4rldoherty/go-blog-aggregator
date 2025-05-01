package command

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/k4rldoherty/rss-blog-aggregator/internal/database"
	"github.com/k4rldoherty/rss-blog-aggregator/internal/rss"
	"github.com/k4rldoherty/rss-blog-aggregator/internal/state"
)

type commands struct {
	Commands map[string]func(*state.State, Command) error
}

type Command struct {
	Name string
	Args []string
}

func NewCommands() *commands {
	return &commands{
		Commands: make(map[string]func(*state.State, Command) error),
	}
}

func (c *commands) Register(name string, f func(*state.State, Command) error) error {
	_, ok := c.Commands[name]
	if ok {
		return errors.New("command already registered")
	}
	c.Commands[name] = f
	return nil
}

func (c *commands) Run(s *state.State, cmd Command) error {
	command, ok := c.Commands[cmd.Name]
	if !ok {
		return errors.New("command doesn't exist with this name")
	}
	err := command(s, cmd)
	if err != nil {
		return err
	}
	return nil
}

func HandlerLogin(s *state.State, cmd Command) error {
	if len(cmd.Args) != 1 {
		return errors.New("incorrect number of args provided")
	}

	user, err := s.Db.GetUser(context.Background(), cmd.Args[0])
	if err != nil {
		return err
	}
	if err := s.Cfg.SetUser(user.Name); err != nil {
		return err
	}
	fmt.Println("Username set.")
	return nil
}

func HandlerRegister(s *state.State, cmd Command) error {
	if len(cmd.Args) != 1 {
		return errors.New("incorrect number of args provided")
	}
	// Creates a struct of paramters to be used in the query
	params := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.Args[0],
	}
	user, err := s.Db.CreateUser(context.Background(), params)
	if err != nil {
		return err
	}
	fmt.Printf("User created. Name: %v ID: %v CreatedAt & UpdatedAt: %v", user.Name, user.ID, user.CreatedAt)
	if err = s.Cfg.SetUser(user.Name); err != nil {
		return err
	}
	return nil
}

func HandlerReset(s *state.State, cmd Command) error {
	err := s.Db.DeleteAllUsers(context.Background())
	if err != nil {
		return err
	}
	fmt.Println("Database successfully reset.")
	return nil
}

func HandlerUsers(s *state.State, cmd Command) error {
	users, err := s.Db.GetUsers(context.Background())
	if err != nil {
		return err
	}
	for _, v := range users {
		if v.Name == s.Cfg.CurrentUserName {
			fmt.Printf("* %v (current)\n", v.Name)
		} else {
			fmt.Printf("* %v\n", v.Name)
		}
	}
	return nil
}

func HandleAgg(s *state.State, cmd Command) error {
	if len(cmd.Args) != 2 {
		return errors.New("incorrect number of args provided")
	}
	feed, err := rss.FetchFeed(context.Background(), cmd.Args[1])
	if err != nil {
		return err
	}
	fmt.Printf("%+v\n", feed)
	return nil
}

func HandleAddFeed(s *state.State, cmd Command) error {
	if len(cmd.Args) != 2 {
		return errors.New("incorrect number of args provided")
	}
	loggedInUser, err := s.Db.GetUser(context.Background(), s.Cfg.CurrentUserName)
	if err != nil {
		return err
	}
	params := database.CreateFeedParams{
		ID:        uuid.New(),
		UserID:    loggedInUser.ID,
		Name:      cmd.Args[0],
		Url:       cmd.Args[1],
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	feed, err := s.Db.CreateFeed(context.Background(), params)
	if err != nil {
		return err
	}
	fmt.Printf("%+v\n", feed)
	return nil
}

func HandleFeeds(s *state.State, cmd Command) error {
	if len(cmd.Args) != 0 {
		return errors.New("incorrect number of args provided")
	}
	feeds, err := s.Db.GetFeeds(context.Background())
	if err != nil {
		return err
	}
	for _, v := range feeds {
		fmt.Printf("Feed Name: %v\nFeed Url: %v\nUser Name: %v \n\n", v.FeedName, v.Url, v.UserName)
	}
	return nil
}
