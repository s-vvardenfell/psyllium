package core

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
