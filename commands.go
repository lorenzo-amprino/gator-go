package main

import (
	"fmt"
)

type command struct {
	Name string
	Args []string
}
type commands struct {
	m map[string]func(*state, command) error
}

func (c commands) run(s *state, cmd command) error {
	handler, ok := c.m[cmd.Name]
	if !ok {
		return fmt.Errorf("unknown command: %s", cmd.Name)
	}

	return handler(s, cmd)
}

func (c commands) Register(name string, handler func(*state, command) error) {
	c.m[name] = handler
}

func NewCommands() commands {
	return commands{
		m: make(map[string]func(*state, command) error),
	}
}
