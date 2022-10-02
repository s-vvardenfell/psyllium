package frontend

import "agent/internal/core"

type grpcFrontEnd struct {
	// c *core.Core
}

// TODO?
// NewFrontend()

func (g grpcFrontEnd) StartWith(c *core.Core) error {
	c.Agent.GetEvents()
	return nil
}
