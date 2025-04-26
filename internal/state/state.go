package state

import "github.com/k4rldoherty/rss-blog-aggregator/internal/config"

type State struct {
	Cfg *config.Config
}

func NewState(cfg *config.Config) *State {
	return &State{
		cfg,
	}
}
