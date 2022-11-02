package core // TODO refactor pkg

import "fmt"

type Core struct {
	Agent Agent
}

func NewCore() *Core {
	return &Core{}
}

func (c *Core) WithAgent(a Agent) *Core {
	c.Agent = a
	return c
}

func (c *Core) Start() error {
	fmt.Println("CORE WORKS!")
	return nil
}
