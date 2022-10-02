package frontend

import "agent/internal/core"

type restFrontEnd struct {
	// c *core.Core
}

func (r restFrontEnd) StartWith(c *core.Core) error {
	c.Agent.GetEvents()
	return nil
}
