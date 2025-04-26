package command

import (
	"errors"
	"fmt"

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
	if err := s.Cfg.SetUser(cmd.Args[0]); err != nil {
		return err
	}
	fmt.Println("Username set.")
	return nil
}
