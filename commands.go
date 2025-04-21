package main

import "fmt"

type command struct {
	Name string
	Args []string
}

type commands struct {
	commandMap map[string]func(*state, command) error
}

func (c *commands) register(name string, handler func(*state, command) error) {
	c.commandMap[name] = handler
}

func (c *commands) run(s *state, cmd command) error {
	if handler, exists := c.commandMap[cmd.Name]; exists {
		return handler(s, cmd)
	}
	return fmt.Errorf("command not found: %s", cmd.Name)
}
