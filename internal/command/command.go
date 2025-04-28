package command

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/k4rldoherty/rss-blog-aggregator/internal/database"
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
